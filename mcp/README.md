# gozero-ai-mcp

Go 实现的 MCP Server，提供 go-zero 代码生成工具。

## 启动
```bash
go run ./mcp/cmd/gozero-ai-mcp
```

## Tools
- `init_project`
- `gen_api_skeleton`

## 模板策略
- 优先使用 `~/.goctl`
- 可通过参数 `template_root` 覆盖
- 内置 fallback 模板位于 `mcp/internal/templates`
