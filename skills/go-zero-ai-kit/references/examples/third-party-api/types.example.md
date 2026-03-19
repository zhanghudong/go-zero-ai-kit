# 示例：第三方 API 类型命名

## 适用场景
- 设计第三方 API 请求/响应结构
- 需要平衡“我方语义”和“协议映射”

## 对应规范
- `references/gozero/third-party-api-workflow.md`

## 示例要点
- 顶层请求/响应类型优先用我方业务语义。
- JSON tag 必须和第三方协议保持一致。
- 更深层协议对象若直接映射更清晰，可以保留第三方术语。

```go
package vendor

type ListProductsReq struct {
    Limit       int64
    Page        int64
    CityIDs     []int64
    CountryIDs  []int64
    CategoryIDs []int64
}

type ListProductsResp struct {
    BaseResponse
    Product *ProductListObject `json:"product,omitempty"`
}

type ProductListObject struct {
    Total       int64             `json:"total"`
    Page        int64             `json:"page"`
    Limit       int64             `json:"limit"`
    HasNext     bool              `json:"has_next"`
    ProductList []ProductListItem `json:"product_list,omitempty"`
}

type ProductListItem struct {
    ProductID  int64  `json:"product_id"`
    Title      string `json:"title"`
    SubTitle   string `json:"sub_title"`
    Currency   string `json:"currency"`
    Price      string `json:"price"`
}

type ActivityPackageObject struct {
    PackageID   int64  `json:"package_id"`
    PackageName string `json:"package_name"`
}
```
