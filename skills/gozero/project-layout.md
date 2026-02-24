# Go-Zero 项目结构规范

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
