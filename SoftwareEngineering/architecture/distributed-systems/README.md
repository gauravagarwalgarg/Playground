# Distributed Systems

## Core Concepts

### CAP Theorem
- **Consistency**: Every read receives the most recent write
- **Availability**: Every request receives a response (not necessarily the latest)
- **Partition Tolerance**: System continues operating despite network partitions

> In practice, P is non-negotiable. Choose between CP (strong consistency) or AP (high availability).

| System | Type | Example |
|--------|------|---------|
| CP | Strong consistency | ZooKeeper, etcd, HBase |
| AP | High availability | Cassandra, DynamoDB, CouchDB |

### Consensus Algorithms
- **Raft**: Leader-based, understandable, used in etcd, CockroachDB
- **Paxos**: Theoretical foundation, complex to implement
- **ZAB**: ZooKeeper Atomic Broadcast, similar to Raft
- **Key idea**: Majority quorum (2f+1 nodes tolerate f failures)

### Replication Strategies
| Strategy | Consistency | Latency | Use Case |
|----------|-------------|---------|----------|
| Synchronous | Strong | High | Financial systems |
| Asynchronous | Eventual | Low | Social media, analytics |
| Semi-synchronous | Tunable | Medium | General purpose |

- **Leader-Follower**: One writer, multiple readers
- **Multi-Leader**: Multiple writers, conflict resolution needed
- **Leaderless**: Quorum reads/writes (Dynamo-style)

### Sharding (Partitioning)
- **Hash-based**: Consistent hashing, uniform distribution
- **Range-based**: Ordered data, range queries efficient
- **Directory-based**: Lookup table, flexible but single point of failure

**Challenges**: Cross-shard queries, rebalancing, hotspots

---

## Interview Talking Points

- CAP is about trade-offs during network partitions, not normal operation
- Consensus is expensive new_textminimize what needs consensus
- Replication strategy depends on consistency requirements and acceptable latency
- Sharding introduces complexity new_textavoid until necessary, then shard by access pattern
