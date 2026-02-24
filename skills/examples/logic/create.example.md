# 示例：Create（新增）

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

type CreateLogic struct {
    logx.Logger
    ctx    context.Context
    svcCtx *svc.ServiceContext
}

// 新增示例
func NewCreateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateLogic {
    return &CreateLogic{
        Logger: logx.WithContext(ctx),
        ctx:    ctx,
        svcCtx: svcCtx,
    }
}

func (l *CreateLogic) Create(req *types.CreateReq) (*types.CreateResp, error) {
    // ========== 步骤1：写入主表 ==========
    // TODO: Insert

    // ========== 步骤2：返回结果 ==========
    return &types.CreateResp{Ok: true}, nil
}
```
