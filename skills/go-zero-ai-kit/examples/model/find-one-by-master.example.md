# 示例：Model 主库查询（读写分离强一致）

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
