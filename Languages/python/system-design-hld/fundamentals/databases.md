# Databases

## SQL vs NoSQL Decision Tree

```
Need ACID transactions?
├── Yes → SQL (PostgreSQL, MySQL)
└── No →
    Need flexible schema?
    ├── Yes → Document DB (MongoDB, DynamoDB)
    └── No →
        Need high write throughput at scale?
        ├── Yes → Wide-column (Cassandra, HBase)
        └── No →
            Need relationships/graph traversal?
            ├── Yes → Graph DB (Neo4j, Neptune)
            └── No → Key-Value (Redis, DynamoDB)
```

## ACID vs BASE

| Property | ACID (SQL) | BASE (NoSQL) |
|----------|-----------|--------------|
| Atomicity | All or nothing | Best effort |
| Consistency | Strong consistency | Eventual consistency |
| Isolation | Transactions isolated | Soft state |
| Durability | Committed = persisted | Eventually durable |

- ACID: Banking, inventory, booking systems
- BASE: Social feeds, analytics, recommendation engines

## Sharding Strategies

### Range-Based Sharding
- Partition by key range (e.g., users A-M → shard 1, N-Z → shard 2)
- Pro: Range queries efficient within a shard
- Con: Hotspots if distribution is uneven (e.g., popular letter ranges)

### Hash-Based Sharding
- Hash(key) mod N → shard assignment
- Pro: Even distribution
- Con: Range queries span all shards; resharding is expensive
- Use consistent hashing to minimize data movement on scale-out

### Directory-Based Sharding
- Lookup service maps each key to its shard
- Pro: Flexible, can move data without changing hash function
- Con: Lookup service is a single point of failure and bottleneck

### Sharding Considerations
- Cross-shard joins are expensive → denormalize or co-locate related data
- Shard key selection is critical choose high-cardinality, even-distribution key
- Auto-sharding: MongoDB, CockroachDB handle it; manual in MySQL/PostgreSQL

## Replication

### Master-Slave (Primary-Replica)
- One master handles writes; replicas serve reads
- Async replication → eventual consistency (read-after-write issues)
- Failover: Promote replica to master (manual or automatic)
- Use: Read-heavy workloads (90% reads)

### Multi-Master
- Multiple nodes accept writes
- Conflict resolution needed: Last-write-wins, vector clocks, CRDTs
- Use: Multi-region writes, high availability requirements
- Examples: CockroachDB, Cassandra, DynamoDB Global Tables

### Replication Lag Mitigation
- Read-your-own-writes: Route reads to master for N seconds after write
- Monotonic reads: Pin user to a specific replica
- Consistent prefix reads: Preserve causal ordering

## Index Types

### B-Tree Index
- Balanced tree, O(log n) reads and writes
- Default in PostgreSQL, MySQL
- Good for: Range queries, equality, sorting
- Storage: On-disk, page-oriented

### LSM Tree (Log-Structured Merge Tree)
- Writes go to in-memory memtable → flush to sorted SSTables
- Compaction merges SSTables periodically
- O(1) writes (append-only), reads may check multiple levels
- Good for: Write-heavy workloads
- Used by: Cassandra, RocksDB, LevelDB, HBase

### When to Use What

| Workload | Index Type | Database |
|----------|-----------|----------|
| Read-heavy, OLTP | B-Tree | PostgreSQL, MySQL |
| Write-heavy, time-series | LSM Tree | Cassandra, InfluxDB |
| Full-text search | Inverted Index | Elasticsearch |
| Geospatial | R-Tree | PostGIS, MongoDB |

## Key Interview Points
- Vertical scaling (bigger machine) vs Horizontal scaling (more machines/shards)
- Connection pooling: PgBouncer, HikariCP avoid connection exhaustion
- Read replicas reduce read load but don't help write throughput
- Database per service (microservices) vs shared database (monolith)
- NewSQL: Combines SQL guarantees with NoSQL scalability (CockroachDB, Spanner)
