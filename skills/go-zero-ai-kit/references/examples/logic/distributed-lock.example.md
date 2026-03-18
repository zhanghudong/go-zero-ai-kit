# 示例：Distributed Lock（分布式锁）

适用场景：
- 按业务键串行化处理高并发请求
- 防止重复提交、重复执行或并发状态冲突

对应规范：
- `references/gozero/logic-patterns.md` 模式8
- `references/gozero/logic-checklist.md`

示例要点：
- 锁 key 必须绑定业务主键
- 优先使用项目统一封装的锁组件
- 锁竞争与系统异常要区分处理
- 临界区内部仍需考虑 CAS、幂等和影响行数校验

```go
// Code scaffolded by goctl. Safe to edit.
// goctl 1.9.2

package example

import (
    "context"
    "fmt"
    "time"

    "example.com/project/cmd/api/internal/svc"

    "codeup.aliyun.com/zlxt/zl-core/backoff"
    "codeup.aliyun.com/zlxt/zl-core/redislock"
    "github.com/zeromicro/go-zero/core/logx"
)

const (
    bizLockKeyFormat = "resource:%d"           // 按业务唯一键调整
    bizLockTTL       = 30 * time.Second        // 按临界区最长耗时评估
    bizLockNamespace = "<biz>"                 // 按业务命名空间调整
    bizLockRetryMax  = 3                       // 按幂等要求与竞争概率调整
    bizLockRetryWait = 150 * time.Millisecond  // 按上游超时与重试节奏调整
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
        redislock.WithNamespace(bizLockNamespace),
        redislock.WithRetry(backoff.RetryConfig{
            MaxRetries: bizLockRetryMax,
            Strategy:   backoff.StrategyFixed,
            BaseDelay:  bizLockRetryWait,
        }),
    )

    // ========== 步骤2：获取锁 ==========
    err := locker.DoWithLock(l.ctx, func() error {
        return fn(l.ctx)
    })
    if err != nil {
        if redislock.IsLockBusy(err) {
            l.Infow("锁竞争未获取，跳过处理", logx.Field("biz_id", bizID), logx.Field("lock_key", lockKey))
            return nil
        }

        l.Errorw("执行加锁临界区失败", logx.Field("biz_id", bizID), logx.Field("lock_key", lockKey), logx.Field("err", err))
        return err
    }

    return nil
}
```
