# 示例：Async Task（异步任务）

```go
// Code scaffolded by goctl. Safe to edit.
// goctl 1.9.2

package example

import (
    "context"
    "time"

    "example.com/project/cmd/api/internal/svc"
    "example.com/project/cmd/api/internal/types"

    "codeup.aliyun.com/zlxt/zl-core/ctxx"
    "github.com/zeromicro/go-zero/core/logx"
    "github.com/zeromicro/go-zero/core/threading"
)

const asyncTaskTimeout = 30 * time.Second

type AsyncTaskLogic struct {
    logx.Logger
    ctx    context.Context
    svcCtx *svc.ServiceContext
}

func NewAsyncTaskLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AsyncTaskLogic {
    return &AsyncTaskLogic{
        Logger: logx.WithContext(ctx),
        ctx:    ctx,
        svcCtx: svcCtx,
    }
}

func (l *AsyncTaskLogic) Handle(req *types.AsyncTaskReq) (*types.AsyncTaskResp, error) {
    // ========== 步骤1：主流程写入并返回 ==========
    bizID := l.svcCtx.Snowflake.Generate()

    // TODO: Insert order...

    resp := &types.AsyncTaskResp{BizId: bizID}

    // ========== 步骤2：触发异步任务（不阻塞主流程） ==========
    reqCopy := *req
    taskCtx := ctxx.CreateNewContextWithSpanContext(l.ctx)

    threading.GoSafeCtx(taskCtx, func() {
        asyncCtx, cancel := ctxx.CreateNewContextWithSpanContextTimeout(taskCtx, asyncTaskTimeout)
        defer cancel()

        threadL := NewAsyncTaskLogic(asyncCtx, l.svcCtx)
        if err := threadL.tryAsyncTask(bizID, &reqCopy); err != nil {
            threadL.Errorw("异步通知失败",
                logx.Field("biz_id", bizID),
                logx.Field("err", err),
            )
        }
    })

    return resp, nil
}

func (l *AsyncTaskLogic) tryAsyncTask(bizID int64, req *types.AsyncTaskReq) error {
    // TODO: 外部通知 / 补偿任务
    _ = bizID
    _ = req
    return nil
}
```
