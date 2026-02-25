# Logic 补全流程（适配 Codex）

本流程用于补全 goctl 生成的 logic 层代码，面向公司现有 go-zero 项目。

## 输入
- 详设文档（优先）
- .api
- .sql
- 已生成的 handler/logic/types/model

## 流程步骤
1. 读详设，确认核心业务流程与状态迁移。
2. 读 .api，核对请求/响应与字段含义。
3. 读 model，确认表结构与可用方法。
4. 在 logic 内实现：参数校验 -> 核心流程 -> DB 操作 -> 返回。
5. 添加分步注释和关键日志字段。
6. 保持函数签名与 goctl scaffold 结构不变。

## 输出
- 完整 logic 实现
- 保持 gofmt 与项目现有风格一致

## 提交规范
- 使用中文 commit message
- 遵循 Conventional Commits
- 建议格式：`<type>(<scope>): <subject>`

详细规范见：`commit-message.md`
