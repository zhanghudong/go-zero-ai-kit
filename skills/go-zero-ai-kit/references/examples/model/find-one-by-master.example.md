# 示例：Model 主库查询（读写分离强一致）

适用场景：
- 刚写后读，且要求强一致
- 主从延迟可能影响读结果

对应规范：
- `references/gozero/db-conventions.md`

示例要点：
- 主库查询统一使用 `/*FORCE_MASTER*/` 注释前缀
- 方法命名统一使用 `ByMaster` 后缀
- `sqlc.ErrNotFound` 需要转换为项目模型层 `ErrNotFound`

```go
package model

import (
    "context"
    "fmt"

    "github.com/zeromicro/go-zero/core/stores/sqlc"
)

// 刚写后读且要求强一致时，使用主库查询
func (m *customExampleModel) FindOneByMaster(ctx context.Context, bizID int64) (*Example, error) {
    var resp Example

    query := fmt.Sprintf("/*FORCE_MASTER*/ select %s from %s where `biz_id` = ? limit 1", exampleRows, m.table)

    err := m.QueryRowNoCacheCtx(ctx, &resp, query, bizID)

    switch err {
    case nil:
        return &resp, nil
    case sqlc.ErrNotFound:
        return nil, ErrNotFound
    default:
        return nil, err
    }
}
```
