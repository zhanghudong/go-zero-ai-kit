# gozero-ai-kit

统一的 go-zero AI 规范与代码生成工具链（skills + context + mcp 三合一），用于公司内部所有 go-zero 项目的统一标准与自动化生成。

**目标**
- 统一研发规范（skills）
- 提供 AI 规则上下文（context）
- 提供 MCP Server 工具（mcp）

**Go 版本**
- Go >= 1.21

**安装**
1. 进入仓库根目录。
2. 可选：安装 goctl（go-zero 官方工具）。

**启动 MCP**
```bash
go run ./mcp/cmd/gozero-ai-mcp
```

**安装 Skills（Codex）**
```bash
# 在 Codex 中安装本仓库 Skill
codex skill install https://github.com/your-org/go-zero-ai-kit
```

**在项目中使用 Skills（示例）**
在目标项目根目录创建 `CODEX.md`：
```md
# Codex Rules
- Use skill: go-zero-ai-kit
```

**工具：init_project**
- 作用：初始化 go-zero API 项目骨架
- 关键参数：
  - `project_name`
  - `module_path`
  - `service_name`
  - `output_dir`
  - `template_root`（默认 `~/.goctl`）
  - `goctl_path`（默认 `goctl`）
  - `force`（默认 false）
  - `dry_run`（默认 false）
  
示例（MCP tools/call）：
```json
{"jsonrpc":"2.0","id":1,"method":"tools/call","params":{"name":"init_project","arguments":{"project_name":"demo","module_path":"example.com/demo","service_name":"api","output_dir":"./demo","force":false,"dry_run":false}}}
```

**工具：gen_api_skeleton**
- 作用：生成 .api + handler/logic/types
- 关键参数：
  - `api_name`
  - `base_path`
  - `output_dir`
  - `template_root`（默认 `~/.goctl`）
  - `goctl_path`（默认 `goctl`）
  - `force`（默认 false）
  - `dry_run`（默认 false）
  
示例（MCP tools/call）：
```json
{"jsonrpc":"2.0","id":2,"method":"tools/call","params":{"name":"gen_api_skeleton","arguments":{"api_name":"order","base_path":"/api/v1","output_dir":"./demo","force":false,"dry_run":false}}}
```

**模板路径配置**
- 默认使用 `~/.goctl`（团队统一模板）
- 可通过 `template_root` 参数覆盖
- 仓库 `mcp/internal/templates` 仅作为 fallback

**goctl 依赖说明**
- 优先使用 `goctl` CLI 生成代码
- 当 `goctl` 不可用时，才使用 fallback 模板

**References**
```text
https://github.com/zeromicro/zero-skills
https://github.com/zeromicro/ai-context
https://github.com/zeromicro/mcp-zero
```
