# 示例：Distributed Lock（分布式锁）

```go
// Code scaffolded by goctl. Safe to edit.
// goctl 1.9.2

package example

import (
    "context"
    "fmt"
    "time"

    "example.com/project/cmd/api/internal/svc"

    "codeup.aliyun.com/zlxt/zl-core/redislock"
    "github.com/zeromicro/go-zero/core/logx"
)

const (
    bizLockKeyFormat  = "lock:<biz>:%d"        // 按业务唯一键与命名空间调整
    bizLockTTL        = 30 * time.Second       // 按临界区最长耗时评估
    bizLockRetryTimes = 3                      // 按幂等要求与竞争概率调整
    bizLockRetryWait  = 150 * time.Millisecond // 按上游超时与重试节奏调整
)

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
    lockKey := fmt.Sprintf(bizLockKeyFormat, bizID)
    locker := redislock.NewLocker(
        l.svcCtx.Redis,
        lockKey,
        redislock.WithTTL(bizLockTTL),
        redislock.WithRetryTimes(bizLockRetryTimes),
        redislock.WithRetryInterval(bizLockRetryWait),
    )

    // ========== 步骤2：获取锁 ==========
    err := locker.Acquire(l.ctx)
    if err != nil {
        l.Errorw("获取分布式锁失败",
            logx.Field("biz_id", bizID),
            logx.Field("lock_key", lockKey),
            logx.Field("err", err),
        )
        return err
    }

    // ========== 步骤3：释放锁 ==========
    defer func() {
        if releaseErr := locker.Release(l.ctx); releaseErr != nil {
            l.Errorw("释放分布式锁失败",
                logx.Field("biz_id", bizID),
                logx.Field("lock_key", lockKey),
                logx.Field("err", releaseErr),
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
