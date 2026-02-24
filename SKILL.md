# go-zero-ai-kit Skill

## Name
`go-zero-ai-kit`

## Summary
团队 go-zero 规范与逻辑补全指南，包含 API/项目规范、logic 风格、错误码与流程检查清单。

## Entry Points
- `skills/gozero/project-layout.md`
- `skills/gozero/api-style.md`
- `skills/gozero/logic-style.md`
- `skills/gozero/logic-workflow.md`
- `skills/gozero/logic-checklist.md`
- `skills/gozero/logic-patterns.md`
- `skills/gozero/error-code.md`

## Usage
- 生成/补充 logic 实现时，严格遵循 `logic-style`、`logic-workflow`、`logic-checklist` 与 `logic-patterns`。
- 错误码与统一响应使用 `skills/gozero/error-code.md` 约定（项目统一 errors + response 规范）。

## Constraints
- 模板来源以 `~/.goctl` 为唯一权威；不得从业务代码反推模板。
- 不做 `req == nil` 判断（goctl 生成调用不会传入 nil）。
- 逻辑实现必须对齐本仓库定义的风格。

## Examples
- `skills/examples/logic/receipt_create.example.md`
- `skills/examples/logic/refund_confirm.example.md`
- `skills/examples/logic/didatravel_get_room_list.style.md`
