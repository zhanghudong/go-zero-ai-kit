# Logging & Tracing

- 统一日志字段：trace_id, span_id, request_id
- 入口处生成 request_id 并注入 ctx
- 逻辑层使用结构化日志
- 关键链路埋点（DB/外部调用）
