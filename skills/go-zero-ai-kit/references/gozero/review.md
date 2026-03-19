# Go-Zero Review 规则

当用户使用 Codex 内建 review 能力审查 go-zero 相关代码时，应用本规则。
本规则只定义审查策略、输出格式和边界，不定义 review 入口。

适用场景：
- Review against a base branch
- Review uncommitted changes
- Review a commit
- Custom review instructions

## 范围
- Review against a base branch：仅审查当前分支相对目标 base branch 的差异。
- Review uncommitted changes：仅审查当前工作区未提交的修改。
- Review a commit：仅审查指定 commit 的改动。
- Custom review instructions：按用户明确指定的审查范围执行；若用户未指定范围，优先按当前 review 入口的默认范围执行。

## 审查重点
1. 审核每个变更文件，重点关注：缺陷、逻辑问题、风险点、风格一致性。
2. 结合 go-zero 约束检查生成代码边界、`.api` 变更流程、logic 实现风格、错误码、日志与链路追踪约束。
3. 忽略自动生成文件的风格类问题，重点看手写业务代码是否符合团队规范。

## 输出要求
- 必须使用中文，不沿用默认英文 review 风格。
- 默认输出两部分：review 结果 + commit message；两者内容不得混写。
- review 结果固定分为：必须修复、建议修复、可选优化。
- 若某一类无内容，明确写“无”。
- Review against a base branch 的 review 结果应按 PR 风格组织，先给总体结论，再给分项问题。
- Review against a base branch 生成的是一条推荐的 squash commit message。
- Review uncommitted changes 生成的是一条基于当前未提交 diff 的普通 commit message。
- Review a commit 默认只输出 review 结果；只有用户明确要求时才补充 commit message。
- 若用户明确要求“只要 review，不要 commit message”，则跳过 commit message 生成。

## 约束
- 只评审当前 review 范围内的代码，不扩展到无关历史变更。
- 不修改代码，只给出审查意见。
- review 模式下禁止主动执行 `go test`、`make test`、`lint`、`build`、安装依赖或启动服务。
- review 模式下仅允许使用读取型命令获取上下文，例如 `git status`、`git diff`、`git diff --name-only`、`git show`、`rg`、`sed` 及必要的文件内容查看。
- 若审查结论必须依赖运行结果，先明确说明原因并征求用户确认；未经确认不得执行。
- 忽略自动生成文件（如 `*_gen.go`、`pb.go`、goctl 生成的 `types.go`）的风格类审查。
- 遇到敏感文件路径（`config/**`, `.env*`, `secrets/**`, `k8s/**`, `etc/**`）只提示跳过，不做内容分析。

## 禁止事项
- 一边 review 一边顺手改代码。
- 把已提交内容或无关历史变更纳入本次 review。
- 对自动生成文件做风格类审查。
- 在 base branch review 中臆造多条历史 commit message；默认只给一条推荐 squash commit message。
