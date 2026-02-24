# API 设计规范

- 版本前缀统一：`/api/v1`
- 资源名使用名词，动作使用动词
- GET 查询、POST 创建、PUT 更新、DELETE 删除
- Request/Response 结构体单独定义在 `internal/types`
- 接口注释要清晰、可用于文档生成
