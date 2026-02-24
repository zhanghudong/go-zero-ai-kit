# 示例：List（分页查询）

```go
// Code scaffolded by goctl. Safe to edit.
// goctl 1.9.2

package example

import (
    "context"

    "example.com/project/cmd/api/internal/svc"
    "example.com/project/cmd/api/internal/types"
    "example.com/project/pkg/errors"

    "github.com/Masterminds/squirrel"
    "github.com/zeromicro/go-zero/core/logx"
)

type ListLogic struct {
    logx.Logger
    ctx    context.Context
    svcCtx *svc.ServiceContext
}

// 列表示例
func NewListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ListLogic {
    return &ListLogic{
        Logger: logx.WithContext(ctx),
        ctx:    ctx,
        svcCtx: svcCtx,
    }
}

func (l *ListLogic) List(req *types.ListReq) (*types.ListResp, error) {
    page := req.Page
    pageSize := req.PageSize
    // 注意：page 默认值由 .api 提供；pageSize 仅做上限裁剪

    rowBuilder := l.svcCtx.ExampleModel.RowBuilder()
    countBuilder := l.svcCtx.ExampleModel.CountBuilder("*")

    if req.TenantId > 0 {
        rowBuilder = rowBuilder.Where(squirrel.Eq{"`tenant_id`": req.TenantId})
        countBuilder = countBuilder.Where(squirrel.Eq{"`tenant_id`": req.TenantId})
    }

    total, err := l.svcCtx.ExampleModel.FindCount(l.ctx, countBuilder)
    if err != nil {
        l.Errorw("查询列表统计失败",
            logx.Field("tenant_id", req.TenantId),
            logx.Field("error", err),
        )
        // 按项目规范选择合适的错误类型（示例：内部错误）
        return nil, errors.ErrInternalError.Wrap(err)
    }

    if total == 0 {
        return &types.ListResp{Total: 0, Page: page, PageSize: pageSize, List: make([]types.ListItem, 0)}, nil
    }

    rows, err := l.svcCtx.ExampleModel.FindPageListByPage(l.ctx, rowBuilder, page, pageSize, "id DESC")
    if err != nil {
        l.Errorw("查询列表失败",
            logx.Field("tenant_id", req.TenantId),
            logx.Field("error", err),
        )
        // 按项目规范选择合适的错误类型
        return nil, errors.ErrInternalError.Wrap(err)
    }

    // TODO: map rows to resp.List
    return &types.ListResp{Total: total, Page: page, PageSize: pageSize, List: make([]types.ListItem, 0, len(rows))}, nil
}
```
