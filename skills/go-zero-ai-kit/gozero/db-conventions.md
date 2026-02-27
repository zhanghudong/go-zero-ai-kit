# 数据库规范

- 表名使用小写下划线
- 主键统一使用 snowflake 或业务约定
- 软删除字段：deleted_at
- 索引命名：idx_<table>_<columns>
- 刚写后读且要求强一致时，必须查询主库（避免读写分离延迟导致脏读/未读到）。
- 需要主库查询时，在 model 增加明确的方法（如 `FindOneByMaster`），不要在 logic 层拼接 SQL。
- 主库查询 SQL 统一使用 `/*FORCE_MASTER*/` 注释前缀，命名统一 `ByMaster` 后缀。
