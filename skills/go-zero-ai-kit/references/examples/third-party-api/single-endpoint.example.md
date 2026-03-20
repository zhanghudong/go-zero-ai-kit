# 示例：第三方 API 单接口实现

## 适用场景
- 在 `pkg/provider/<vendor>/` 下新增一个接口文件
- 需要统一日志、请求、解析写法
- 需要参考 `klook` 中“一个接口一个文件”的结构

## 对应规范
- `references/gozero/third-party-api-workflow.md`
- `references/gozero/logging-tracing.md`

## 示例要点
- 一个接口一个文件。
- 请求链路直接展开，避免收益很小的额外封装。
- 针对 GET + query 参数接口，可以在同文件放一个小型 query builder helper。
- 日志里优先打印我方真正排障需要的字段。

```go
package klook

import (
	"context"
	"fmt"
	"net/url"
	"strconv"
	"strings"

	"github.com/zeromicro/go-zero/core/logx"
)

// ListProducts 获取商品列表。
// 文档说明该接口可按城市、国家、类目等条件分页查询可售卖的商品基础信息。
func (c *Client) ListProducts(ctx context.Context, req *ListProductsReq) (*ListProductsResp, error) {
	logger := logx.WithContext(ctx)

	logger.Infow("开始请求客路商品列表接口",
		logx.Field("limit", req.Limit),
		logx.Field("page", req.Page),
	)

	var resp ListProductsResp
	path := fmt.Sprintf("/v3/activities?%s", buildListProductsQuery(req).Encode())
	apiResp, err := c.httpClient.Get(path).
		WithContext(ctx).
		WithHeaders(c.defaultHeaders()).
		Do()
	if err != nil {
		logger.Errorw("请求客路商品列表接口失败", logx.Field("error", err))
		return nil, err
	}

	if err = apiResp.ParseJSON(&resp); err != nil {
		logger.Errorw("解析客路商品列表响应失败", logx.Field("error", err))
		return nil, err
	}

	return &resp, nil
}

func buildListProductsQuery(req *ListProductsReq) url.Values {
	values := url.Values{}
	values.Set("limit", strconv.FormatInt(req.Limit, 10))
	values.Set("page", strconv.FormatInt(req.Page, 10))

	if len(req.CityIDs) > 0 {
		values.Set("city_ids", joinInt64List(req.CityIDs))
	}
	if len(req.CountryIDs) > 0 {
		values.Set("country_ids", joinInt64List(req.CountryIDs))
	}
	if len(req.CategoryIDs) > 0 {
		values.Set("category_ids", joinInt64List(req.CategoryIDs))
	}

	return values
}

func joinInt64List(ids []int64) string {
	parts := make([]string, 0, len(ids))
	for _, id := range ids {
		parts = append(parts, strconv.FormatInt(id, 10))
	}

	return strings.Join(parts, ",")
}
```
