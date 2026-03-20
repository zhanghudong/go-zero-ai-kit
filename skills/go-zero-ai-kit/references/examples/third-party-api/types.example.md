# 示例：第三方 API 类型命名

## 适用场景
- 设计第三方 API 请求/响应结构
- 需要平衡“我方语义”和“协议映射”
- 需要参考 `klook/types.go` 的命名方式

## 对应规范
- `references/gozero/third-party-api-workflow.md`

## 示例要点
- 顶层请求/响应类型优先用我方业务语义。
- JSON tag 必须和第三方协议保持一致。
- 更深层协议对象若直接映射更清晰，可以保留第三方术语，例如 `ActivityPackageObject`、`ExtraInfoObject`。
- 如果第三方统一返回 `success + error`，先抽 `BaseResponse` 和 `ErrorObject`。

```go
package klook

type ErrorObject struct {
	Code    string `json:"code"`     // 业务错误码
	Status  int    `json:"status"`   // HTTP 语义状态码
	Message string `json:"message"`  // 错误详情
	Help    string `json:"help"`     // 帮助文档链接
	TraceID string `json:"trace_id"` // 第三方链路追踪 ID
}

type BaseResponse struct {
	Success bool         `json:"success"`
	Error   *ErrorObject `json:"error,omitempty"`
}

type ListProductsReq struct {
	Limit       int64   // 每页数量
	Page        int64   // 页码，从 1 开始
	CityIDs     []int64 // 城市 ID 列表
	CountryIDs  []int64 // 国家 ID 列表
	CategoryIDs []int64 // 类目 ID 列表
}

type ListProductsResp struct {
	BaseResponse
	Activity *ProductListObject `json:"activity,omitempty"`
}

type ProductListObject struct {
	Total        int64             `json:"total"`
	Page         int64             `json:"page"`
	Limit        int64             `json:"limit"`
	HasNext      bool              `json:"has_next"`
	ActivityList []ProductListItem `json:"activity_list,omitempty"`
}

type CreateOrderReq struct {
	AgentOrderID string            `json:"agent_order_id"`
	Timestamp    int64             `json:"timestamp,omitempty"`
	ContactInfo  ContactInfoObject `json:"contact_info"`
	Items        []CreateOrderItem `json:"items"`
}

type CreateOrderItem struct {
	PackageID        int64                 `json:"package_id"`
	StartTime        string                `json:"start_time"`
	BookingExtraInfo []ExtraInfoObject     `json:"booking_extra_info,omitempty"`
	UnitExtraInfo    []UnitExtraInfoObject `json:"unit_extra_info,omitempty"`
	SKUList          []CreateOrderSKU      `json:"sku_list"`
}

type ExtraInfoObject struct {
	Key       string            `json:"key"`
	Content   string            `json:"content,omitempty"`
	Selected  []ExtraInfoObject `json:"selected,omitempty"`
	InputType string            `json:"input_type,omitempty"`
}

type ActivityPackageObject struct {
	PackageID   int64  `json:"package_id"`
	PackageName string `json:"package_name"`
	TimeZone    string `json:"time_zone"`
	TicketType  int64  `json:"ticket_type"`
}
```
