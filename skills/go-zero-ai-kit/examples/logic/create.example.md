# 示例：Create（新增）

```go
// Code scaffolded by goctl. Safe to edit.
// goctl 1.9.2

package example

import (
    "context"
    "strconv"

    "example.com/project/cmd/api/internal/svc"
    "example.com/project/cmd/api/internal/types"
    "example.com/project/model"
    "example.com/project/pkg/constant"
    "example.com/project/pkg/dberr"
    "example.com/project/pkg/errors"

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
    // ========== 步骤1：幂等检查 ==========
    exist, err := l.svcCtx.ExampleModel.FindOneByUniqKey(l.ctx, req.UniqKey)
    switch {
    case err == nil:
        return &types.CreateResp{Id: strconv.FormatInt(exist.Id, 10)}, nil
    case errors.Is(err, model.ErrNotFound):
        // 不存在，继续创建
    case err != nil:
        l.Errorw("查询记录失败",
            logx.Field("uniq_key", req.UniqKey),
            logx.Field("error", err),
        )
        return nil, errors.ErrInternalError.Wrap(err)
    }

    // ========== 步骤2：生成主键并写入 ==========
    id := l.svcCtx.Snowflake.Generate()
    row := &model.Example{
        Id:     id,
        Name:   req.Name,
        UniqKey: req.UniqKey,
        Status: constant.StatusInit,
    }

    _, err = l.svcCtx.ExampleModel.Insert(l.ctx, row)
    if err != nil {
        if dberr.IsDuplicateEntry(err) {
            return nil, errors.ErrDuplicateKey.Wrap(err)
        }
        l.Errorw("写入失败",
            logx.Field("uniq_key", req.UniqKey),
            logx.Field("error", err),
        )
        return nil, errors.ErrInternalError.Wrap(err)
    }

    // ========== 步骤3：返回 ==========
    return &types.CreateResp{Id: strconv.FormatInt(id, 10)}, nil
}
```
