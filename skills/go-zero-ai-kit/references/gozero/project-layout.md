# Go-Zero 项目结构规范

适用场景：
- 判断代码应该放在哪个目录
- 评估某个新增结构是否符合 go-zero 项目分层
- 确认生成代码与业务代码边界

## 基本结构
- `cmd/<service>/`：服务入口
- `internal/`：业务核心代码
- `internal/config/`：配置结构体与加载
- `internal/handler/`：HTTP 处理器
- `internal/logic/`：业务逻辑
- `internal/svc/`：ServiceContext 依赖注入
- `internal/types/`：API 请求/响应类型
- `etc/`：配置文件

## 关键原则
- 以 goctl 模板为唯一权威来源
- 业务代码不得反推模板
- 所有新增结构必须可通过模板生成

禁止事项：
- 把业务逻辑写入 `handler`。
- 通过手改生成文件改变项目结构。
- 在分层之外随意扩散接口相关类型定义。
