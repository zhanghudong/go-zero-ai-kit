# 错误码规范（基于项目统一 errors 包）

以项目统一的 `errors` 包为标准，不自定义分段区间。

## 基础约定
- 成功：`SuccessCode = 0`，`message = "success"`
- 失败：返回 `errors.BusinessError` 或普通 error
- 统一响应结构由项目统一 `response` 包输出

## 统一错误码（示例）
- `CodeInternalError` 内部错误
- `CodeInvalidParam` 参数错误
- `CodeUnauthorized` 未授权
- `CodeForbidden` 禁止访问
- `CodeNotFound` 资源不存在
- `CodeTimeout` 超时
- `CodeServiceBusy` 服务繁忙
- `CodeNetworkError` 网络错误
- `CodeDuplicateKey` 重复键
- `CodeDataNotFound` 数据不存在
- `CodePermissionDenied` 权限不足

## 使用方式
- 逻辑层返回 `errors.ErrInvalidParam` 等预定义错误（可 `Wrap/Wrapf`）
- handler 层统一用 `response.WriteJSON(ctx, w, data, err)` 输出
- **不要**在 handler 里手工拼装错误码
