# 示例：Update（更新）

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

type UpdateLogic struct {
    logx.Logger
    ctx    context.Context
    svcCtx *svc.ServiceContext
}

// 更新示例
func NewUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateLogic {
    return &UpdateLogic{
        Logger: logx.WithContext(ctx),
        ctx:    ctx,
        svcCtx: svcCtx,
    }
}

func (l *UpdateLogic) Update(req *types.UpdateReq) (*types.UpdateResp, error) {
    // ========== 步骤1：查询当前状态 ==========
    // TODO: FindOne

    // ========== 步骤2：按字段构建更新内容 ==========
    // TODO: if req.Name != "" { update.Name = req.Name }

    // ========== 步骤3：执行更新 ==========
    // TODO: Update

    // ========== 步骤4：返回 ==========
    return &types.UpdateResp{Ok: true}, nil
}
```
