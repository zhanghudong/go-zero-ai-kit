# 示例：Read（查询）

适用场景：
- 详情查询、单条读取
- 查询后需要做字段映射或时间格式化

对应规范：
- `references/gozero/logic-patterns.md` 模式1
- `references/gozero/logic-style.md`

示例要点：
- 先查询，再判断不存在或非法状态
- 返回阶段统一做响应字段映射
- 时间字段优先使用项目公共时间工具

```go
// Code scaffolded by goctl. Safe to edit.
// goctl 1.9.2

package example

import (
    "context"

    "example.com/project/cmd/api/internal/svc"
    "example.com/project/cmd/api/internal/types"
    "example.com/project/pkg/timeutil"

    "github.com/zeromicro/go-zero/core/logx"
)

type ReadLogic struct {
    logx.Logger
    ctx    context.Context
    svcCtx *svc.ServiceContext
}

// 查询示例
func NewReadLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ReadLogic {
    return &ReadLogic{
        Logger: logx.WithContext(ctx),
        ctx:    ctx,
        svcCtx: svcCtx,
    }
}

func (l *ReadLogic) Read(req *types.ReadReq) (*types.ReadResp, error) {
    // ========== 步骤1：查询 ==========
    // TODO: FindOne

    // ========== 步骤2：返回 ==========
    return &types.ReadResp{
        CreateTime: timeutil.NowDateTime(),
    }, nil
}
```
