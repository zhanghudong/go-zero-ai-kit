# Review 规则（输入 review 时执行）

当用户输入 `review` 时，执行以下流程并输出中文 commit message：

## 流程
1. 仅检查未提交的修改（`git status`、`git diff`）。
2. 审核每个变更文件，重点关注：缺陷、逻辑问题、风险点、风格一致性。
3. 输出结果分为：必须修复、建议修复、可选优化。
4. 生成中文 commit message（Conventional Commits 格式）。

## 约束
- 只评审未提交代码，不评审已提交内容。
- 不修改代码，只给出审查意见。
- 忽略自动生成文件（如 `*_gen.go`、`pb.go`、goctl 生成的 `types.go`）。
- 遇到敏感文件路径（`config/**`, `.env*`, `secrets/**`, `k8s/**`, `etc/**`）只提示跳过，不做内容分析。

## Commit Message 生成规则
- 新增功能/首次实现：使用 `feat`
- 修复已存在逻辑缺陷：使用 `fix`

## Commit Message 格式
使用 `commit-message.md` 规范。
