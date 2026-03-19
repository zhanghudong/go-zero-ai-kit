# 示例：第三方 API client 结构

## 适用场景
- 在 `pkg/provider/<vendor>/` 下新建第三方 API 客户端
- 需要统一 `client.go` 的职责边界

## 对应规范
- `references/gozero/third-party-api-workflow.md`
- `references/gozero/project-layout.md`

## 示例要点
- `client.go` 只放客户端、配置、接口定义、公共 headers。
- 对外接口方法名使用我方业务语义。
- 不在 `client.go` 堆每个接口的实现细节。

```go
package vendor

import (
    "context"

    "codeup.aliyun.com/zlxt/zl-core/httpclient"
)

const (
    DefaultBaseURL = "https://sandbox.example.com"
    DefaultLanguage = "zh_CN"
)

type ServiceClient interface {
    ListProducts(ctx context.Context, req *ListProductsReq) (*ListProductsResp, error)
    GetProductDetail(ctx context.Context, productID string) (*ProductDetailResp, error)
    ValidateOrder(ctx context.Context, req *CreateOrderReq) (*ValidateOrderResp, error)
}

type Client struct {
    apiKey     string
    language   string
    httpClient *httpclient.Client
}

type Config struct {
    APIKey     string            `json:",optional"`
    BaseURL    string            `json:",optional"`
    Language   string            `json:",optional"`
    HttpClient httpclient.Config `json:",optional"`
}

func NewClient(cfg Config) *Client {
    baseURL := cfg.BaseURL
    if baseURL == "" {
        baseURL = DefaultBaseURL
    }

    language := cfg.Language
    if language == "" {
        language = DefaultLanguage
    }

    httpCli := httpclient.NewClient(cfg.HttpClient,
        httpclient.WithURLResolver(httpclient.NewFixedURLResolver(baseURL)),
    )

    return &Client{
        apiKey:     cfg.APIKey,
        language:   language,
        httpClient: httpCli,
    }
}

func (c *Client) defaultHeaders() map[string]string {
    return map[string]string{
        "Accept-Language": c.language,
        "Content-Type":    "application/json",
        "X-API-Key":       c.apiKey,
    }
}
```
