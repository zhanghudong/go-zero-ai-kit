# 示例：第三方 API client 结构

## 适用场景
- 在 `pkg/provider/<vendor>/` 下新建第三方 API 客户端
- 需要统一 `client.go` 的职责边界
- 需要参考 `zl-procurement/pkg/provider/klook` 的拆分方式

## 对应规范
- `references/gozero/third-party-api-workflow.md`
- `references/gozero/project-layout.md`

## 示例要点
- `client.go` 只放客户端、配置、接口定义、公共 headers。
- 对外接口方法名使用我方业务语义，不直接暴露第三方路径名。
- 不在 `client.go` 堆每个接口的实现细节。
- 如果第三方平台存在 webhook/回调配置，但本轮不实现消费逻辑，可以先在 `Config` 里预留配置位。

```go
package klook

import (
	"context"

	"codeup.aliyun.com/zlxt/zl-core/httpclient"
)

const (
	// DefaultBaseURL 客路沙箱 API 地址。
	DefaultBaseURL = "https://sandbox-api.klktech.com"
	// DefaultLanguage 客路接口默认语言，请求头使用 Accept-Language 传递。
	DefaultLanguage = "zh_CN"
)

type ServiceClient interface {
	// ListProducts 按页查询商品基础信息列表，可按城市、国家、类目过滤。
	ListProducts(ctx context.Context, req *ListProductsReq) (*ListProductsResp, error)
	// CheckPriceAndStock 在用户填写联系人信息前先校验价格库存与下单结构。
	CheckPriceAndStock(ctx context.Context, req []CheckPriceAndStockItem) (*CheckPriceAndStockResp, error)
	// ValidateOrder 在正式下单前校验订单参数、日期、SKU 约束与附加信息。
	ValidateOrder(ctx context.Context, req *CreateOrderReq) (*ValidateOrderResp, error)
	// CreateOrder 根据 SKU、数量、联系人和附加信息创建客路订单。
	CreateOrder(ctx context.Context, req *CreateOrderReq) (*CreateOrderResp, error)
	// GetProductDetail 查询单个商品的详情信息，包括套餐、SKU、规格定义和联系人填写规则。
	GetProductDetail(ctx context.Context, activityID string) (*ProductDetailResp, error)
}

type Client struct {
	apiKey     string             // 客路 API Key
	language   string             // 默认语言，如 zh_CN / en_US
	httpClient *httpclient.Client // 底层 HTTP 客户端
}

type Config struct {
	APIKey     string            `json:",optional"` // 客路 API Key
	BaseURL    string            `json:",optional"` // API 基础地址，默认使用客路沙箱地址
	Language   string            `json:",optional"` // 默认语言，如 zh_CN / en_US
	Webhook    WebhookConfig     `json:",optional"` // 回调配置，仅存放我方预留地址
	HttpClient httpclient.Config `json:",optional"` // HTTP 客户端配置
}

type WebhookConfig struct {
	TicketConfirmURL string `json:",optional"` // 票券确认回调地址
	OrderRefundURL   string `json:",optional"` // 订单退款回调地址
	CodeRedeemURL    string `json:",optional"` // 原始凭证核销回调地址
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
