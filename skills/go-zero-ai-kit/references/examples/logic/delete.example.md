# 示例：Delete（删除）

适用场景：
- 删除或软删除资源
- 删除前需要先校验状态、归属或存在性

对应规范：
- `references/gozero/logic-patterns.md` 模式2
- `references/gozero/logic-style.md`

示例要点：
- 先查询当前状态，避免盲删
- 删除操作应区分物理删除与软删除
- 返回前保留明确的成功结果

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
