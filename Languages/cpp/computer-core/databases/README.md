# Databases - Interview Notes

## SQL vs NoSQL

| Aspect | SQL (PostgreSQL, MySQL) | NoSQL (MongoDB, Cassandra) |
|--------|------------------------|---------------------------|
| Schema | Fixed, predefined | Flexible, schema-less |
| Scaling | Vertical (primarily) | Horizontal (sharding) |
| Joins | Native, efficient | Application-level |
| ACID | Full support | Varies (eventual consistency) |
| Use case | Complex queries, transactions | High write throughput, flexible data |

## ACID Properties
- **Atomicity**: All or nothing (transaction commits fully or rolls back)
- **Consistency**: DB moves from one valid state to another
- **Isolation**: Concurrent transactions don't interfere
- **Durability**: Committed data survives crashes (WAL, fsync)

## Indexing
- **B-Tree**: Default, good for range queries, O(log n) lookup
- **Hash Index**: O(1) exact match, no range support
- **Composite Index**: (col1, col2) -- leftmost prefix rule
- **Covering Index**: Index contains all needed columns (no table lookup)

## Isolation Levels
1. **Read Uncommitted**: Dirty reads possible
2. **Read Committed**: No dirty reads (PostgreSQL default)
3. **Repeatable Read**: No phantom reads (MySQL InnoDB default)
4. **Serializable**: Full isolation (slowest)

## Sharding Strategies
- **Range-based**: user_id 1-1M on shard 1, 1M-2M on shard 2
- **Hash-based**: hash(user_id) % num_shards
- **Geo-based**: US users on US shard, EU on EU shard
- **Consistent hashing**: Minimizes redistribution on shard add/remove

## Common Interview Questions
1. Design a schema for Twitter (users, tweets, follows, likes)
2. How would you handle a slow query? (EXPLAIN, indexing, denormalization)
3. When would you choose NoSQL over SQL?
4. How does a B-Tree index work internally?
5. What is write-ahead logging (WAL)?
