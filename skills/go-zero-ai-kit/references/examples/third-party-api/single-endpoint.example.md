# 示例：第三方 API 单接口实现

## 适用场景
- 在 `pkg/provider/<vendor>/` 下新增一个接口文件
- 需要统一日志、请求、解析、业务错误处理写法

## 对应规范
- `references/gozero/third-party-api-workflow.md`
- `references/gozero/logging-tracing.md`

## 示例要点
- 一个接口一个文件。
- 请求链路直接展开，避免收益很小的额外封装。
- 区分网络错误、解析错误、业务错误。

```go
package vendor

import (
    "context"
    "fmt"

    "github.com/zeromicro/go-zero/core/logx"
)

// GetProductDetail 根据商品 ID 查询商品详情。
func (c *Client) GetProductDetail(ctx context.Context, productID string) (*ProductDetailResp, error) {
    logger := logx.WithContext(ctx)

    logger.Infow("开始请求第三方商品详情接口", logx.Field("product_id", productID))

    var resp ProductDetailResp
    path := fmt.Sprintf("/v1/products/%s", productID)
    apiResp, err := c.httpClient.Get(path).
        WithContext(ctx).
        WithHeaders(c.defaultHeaders()).
        Do()
    if err != nil {
        err = fmt.Errorf("请求第三方 API 失败: %w", err)
        logger.Errorw("请求第三方商品详情接口失败", logx.Field("error", err))
        return nil, err
    }

    if err = apiResp.JSON(&resp); err != nil {
        err = fmt.Errorf("解析第三方 API 响应失败: %w", err)
        logger.Errorw("请求第三方商品详情接口失败", logx.Field("error", err))
        return nil, err
    }

    if !resp.Success {
        err = buildVendorError(resp.Error)
        logger.Errorw("第三方商品详情接口返回错误", logx.Field("error", err))
        return nil, err
    }

    return &resp, nil
}
```
