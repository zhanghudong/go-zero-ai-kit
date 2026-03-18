# Logic 补全流程（适配 Codex）

本流程用于补全 goctl 生成的 logic 层代码，面向公司现有 go-zero 项目。

适用场景：
- 用户已完成 `.api` 变更并执行生成命令
- 需要补全或重构 `internal/logic` 业务实现
- 需要按团队 go-zero 风格实现一个已有 scaffold 的接口

前置条件：
- 已完成 `.api` 修改
- 已执行 `make`、`goctl api go` 或项目等价生成命令
- 相关 `handler/logic/types` 已按最新 `.api` 生成

如果尚未满足前置条件，先遵循 `api-dev-workflow.md` 的阶段一规则，只修改 `.api`，不要提前补 `handler/logic/types`。

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

## 禁止事项
- 在未生成最新 scaffold 前直接补 `logic`。
- 修改 goctl 生成的函数签名或目录结构。
- 在 `handler` 中承载业务逻辑。
- 通过手改 `types.go` 修正请求/响应结构。

## 提交规范
- 使用中文 commit message
- 遵循 Conventional Commits
- 建议格式：`<type>(<scope>): <subject>`

详细规范见：`commit-message.md`
