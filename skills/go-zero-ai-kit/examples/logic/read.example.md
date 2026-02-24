# 示例：Read（查询）

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
