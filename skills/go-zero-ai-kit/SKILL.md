---
name: go-zero-ai-kit
description: Team go-zero standards, logic style, and workflow guidance.
metadata:
  short-description: go-zero standards and logic style
---

# go-zero-ai-kit Skill

## Summary
团队 go-zero 规范与逻辑补全指南，包含 API/项目规范、logic 风格、错误码与流程检查清单。

## Entry Points
- ../gozero/project-layout.md
- ../gozero/api-style.md
- ../gozero/logic-style.md
- ../gozero/logic-workflow.md
- ../gozero/logic-checklist.md
- ../gozero/logic-patterns.md
- ../gozero/error-code.md

## Examples
- ../examples/logic/create.example.md
- ../examples/logic/read.example.md
- ../examples/logic/update.example.md
- ../examples/logic/delete.example.md
- ../examples/logic/list.example.md

## Constraints
- 模板来源以 `~/.goctl` 为唯一权威；不得从业务代码反推模板。
- 不做 `req == nil` 判断。
- 逻辑实现必须对齐本仓库定义的风格。
