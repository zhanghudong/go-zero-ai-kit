# Review 规则（输入 review 时执行）

当用户输入 `review` 时，默认输出两部分：代码审查结论 + commit message。

## 流程
1. 仅检查未提交的修改（`git status`、`git diff`）。
2. 审核每个变更文件，重点关注：缺陷、逻辑问题、风险点、风格一致性。
3. 输出结果分为：必须修复、建议修复、可选优化。
4. 基于同一批 `git diff` 生成中文 commit message（遵循 `commit-message.md`）。

## 约束
- 只评审未提交代码，不评审已提交内容。
- 不修改代码，只给出审查意见。
- 忽略自动生成文件（如 `*_gen.go`、`pb.go`、goctl 生成的 `types.go`）。
- 遇到敏感文件路径（`config/**`, `.env*`, `secrets/**`, `k8s/**`, `etc/**`）只提示跳过，不做内容分析。
- 输出顺序固定：先输出 review 结果，再单独输出 commit message；两者内容不得混写。
- 若用户明确要求“只要 review，不要 commit message”，则跳过 commit message 生成。
