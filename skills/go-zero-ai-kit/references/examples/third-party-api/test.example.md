# 示例：第三方 API 轻量测试

## 适用场景
- 为第三方 API 封装补最小必要测试
- 需要统一环境变量、示例数据、输出方式
- 需要区分单元测试和轻量集成测试

## 对应规范
- `references/gozero/third-party-api-workflow.md`

## 示例要点
- 先补不依赖外网的 `httptest` 单元测试，校验 path、header、query、body。
- 再补 1 个轻量集成测试，从环境变量读取密钥，不把凭证写死。
- 集成测试可以抽 `requireTestClient` 之类的 helper，避免每个文件重复初始化。

```go
package klook

import (
	"context"
	"io"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"zl-procurement/pkg/util"

	"github.com/stretchr/testify/require"
)

func TestClient_ListProducts_RequestHeadersAndQuery(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		require.Equal(t, http.MethodGet, r.Method)
		require.Equal(t, "/v3/activities", r.URL.Path)
		require.Equal(t, "20", r.URL.Query().Get("limit"))
		require.Equal(t, "3", r.URL.Query().Get("page"))
		require.Equal(t, "1,2", r.URL.Query().Get("city_ids"))
		require.Equal(t, "test-key", r.Header.Get("X-API-Key"))
		require.Equal(t, "en_US", r.Header.Get("Accept-Language"))

		_, err := io.WriteString(w, `{"success":true,"activity":{"total":1,"page":3,"limit":20,"has_next":false,"activity_list":[{"activity_id":101,"title":"demo"}]}}`)
		require.NoError(t, err)
	}))
	defer server.Close()

	client := NewClient(Config{
		APIKey:   "test-key",
		BaseURL:  server.URL,
		Language: "en_US",
	})

	resp, err := client.ListProducts(context.Background(), &ListProductsReq{
		Limit:   20,
		Page:    3,
		CityIDs: []int64{1, 2},
	})
	require.NoError(t, err)
	require.NotNil(t, resp)
	require.NotNil(t, resp.Activity)
	require.Len(t, resp.Activity.ActivityList, 1)
}

func TestClient_GetProductDetail(t *testing.T) {
	apiKey := os.Getenv("KLOOK_API_KEY")
	if apiKey == "" {
		t.Skip("skip integration test: KLOOK_API_KEY is empty")
	}

	activityID := os.Getenv("KLOOK_ACTIVITY_ID")
	if activityID == "" {
		activityID = "19"
	}

	client := NewClient(Config{
		APIKey:   apiKey,
		BaseURL:  DefaultBaseURL,
		Language: DefaultLanguage,
	})

	resp, err := client.GetProductDetail(context.Background(), activityID)
	require.NoError(t, err)
	t.Log(util.StructToJSON(resp))
}
```
