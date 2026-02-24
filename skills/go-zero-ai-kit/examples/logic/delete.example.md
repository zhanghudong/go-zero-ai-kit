# 示例：Delete（删除）

```go
// Code scaffolded by goctl. Safe to edit.
// goctl 1.9.2

package example

import (
    "context"

    "example.com/project/cmd/api/internal/svc"
    "example.com/project/cmd/api/internal/types"

    "github.com/zeromicro/go-zero/core/logx"
)

type DeleteLogic struct {
    logx.Logger
    ctx    context.Context
    svcCtx *svc.ServiceContext
}

// 删除示例
func NewDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteLogic {
    return &DeleteLogic{
        Logger: logx.WithContext(ctx),
        ctx:    ctx,
        svcCtx: svcCtx,
    }
}

func (l *DeleteLogic) Delete(req *types.DeleteReq) (*types.DeleteResp, error) {
    // ========== 步骤1：查询当前状态 ==========
    // TODO: FindOne

    // ========== 步骤2：删除/软删 ==========
    // TODO: Delete

    // ========== 步骤3：返回 ==========
    return &types.DeleteResp{Ok: true}, nil
}
```
