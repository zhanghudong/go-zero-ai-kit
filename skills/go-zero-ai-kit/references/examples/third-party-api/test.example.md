# 示例：第三方 API 轻量测试

## 适用场景
- 为第三方 API 封装补最小必要测试
- 需要统一环境变量、示例数据、输出方式

## 对应规范
- `references/gozero/third-party-api-workflow.md`

## 示例要点
- 优先轻量集成测试，不急着引入复杂 mock。
- 从环境变量读取密钥，不把凭证写死。
- 默认测试数据优先使用文档示例值。

```go
package vendor

import (
    "context"
    "os"
    "testing"

    "zl-procurement/pkg/util"
)

func TestClient_GetProductDetail(t *testing.T) {
    apiKey := os.Getenv("VENDOR_API_KEY")
    if apiKey == "" {
        t.Skip("skip integration test: VENDOR_API_KEY is empty")
    }

    productID := os.Getenv("VENDOR_PRODUCT_ID")
    if productID == "" {
        productID = "19"
    }

    client := NewClient(Config{
        APIKey:   apiKey,
        BaseURL:  DefaultBaseURL,
        Language: DefaultLanguage,
    })

    resp, err := client.GetProductDetail(context.Background(), productID)
    if err != nil {
        t.Fatal(err)
    }
    t.Log(util.StructToJSON(resp))
}
```
