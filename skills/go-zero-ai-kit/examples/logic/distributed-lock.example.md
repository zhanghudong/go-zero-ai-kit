# 示例：Distributed Lock（分布式锁）

```go
// Code scaffolded by goctl. Safe to edit.
// goctl 1.9.2

package example

import (
    "context"
    "fmt"

    "example.com/project/cmd/api/internal/svc"

    "github.com/zeromicro/go-zero/core/logx"
    "github.com/zeromicro/go-zero/core/stores/redis"
)

const bizLockExpireSeconds = 30

type LockGuardLogic struct {
    logx.Logger
    ctx    context.Context
    svcCtx *svc.ServiceContext
}

func NewLockGuardLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LockGuardLogic {
    return &LockGuardLogic{
        Logger: logx.WithContext(ctx),
        ctx:    ctx,
        svcCtx: svcCtx,
    }
}

func (l *LockGuardLogic) withBizLock(bizID int64, fn func(ctx context.Context) error) error {
    // ========== 步骤1：构建锁 ==========
    lockKey := fmt.Sprintf("lock:biz:%d", bizID)
    lock := redis.NewRedisLock(l.svcCtx.Redis, lockKey)
    lock.SetExpire(bizLockExpireSeconds)

    // ========== 步骤2：获取锁 ==========
    acquired, err := lock.AcquireCtx(l.ctx)
    if err != nil {
        l.Errorw("获取分布式锁失败",
            logx.Field("biz_id", bizID),
            logx.Field("lock_key", lockKey),
            logx.Field("err", err),
        )
        return err
    }
    if !acquired {
        l.Infow("锁竞争未获取，跳过处理",
            logx.Field("biz_id", bizID),
            logx.Field("lock_key", lockKey),
        )
        return nil
    }

    // ========== 步骤3：释放锁 ==========
    defer func() {
        if released, releaseErr := lock.ReleaseCtx(l.ctx); releaseErr != nil {
            l.Errorw("释放分布式锁失败",
                logx.Field("biz_id", bizID),
                logx.Field("lock_key", lockKey),
                logx.Field("err", releaseErr),
            )
        } else if !released {
            l.Errorw("分布式锁已过期或被其他实例释放",
                logx.Field("biz_id", bizID),
                logx.Field("lock_key", lockKey),
            )
        }
    }()

    return fn(l.ctx)
}

func (l *LockGuardLogic) Handle(bizID int64) error {
    return l.withBizLock(bizID, func(ctx context.Context) error {
        // TODO: 临界区逻辑；推荐使用 CAS 状态迁移
        _ = ctx
        return nil
    })
}
```
