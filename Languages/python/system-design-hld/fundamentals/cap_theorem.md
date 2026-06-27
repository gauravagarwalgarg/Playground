# CAP Theorem

## Overview
In a distributed system, you can only guarantee **two out of three**:
- **C**onsistency: Every read receives the most recent write
- **A**vailability: Every request receives a response (may not be latest)
- **P**artition Tolerance: System continues operating despite network partitions

Since network partitions are inevitable in distributed systems, the real choice is **CP vs AP**.

## CP Systems (Consistency + Partition Tolerance)
- During a partition, system rejects requests rather than serve stale data
- Prioritizes correctness over availability

| System | Why CP |
|--------|--------|
| ZooKeeper | Leader-based consensus (ZAB protocol), rejects writes without quorum |
| MongoDB | Single primary for writes, returns error if primary unreachable |
| HBase | Strong consistency via single RegionServer per region |
| etcd | Raft consensus, linearizable reads |

## AP Systems (Availability + Partition Tolerance)
- During a partition, system serves requests but may return stale data
- Prioritizes uptime and responsiveness

| System | Why AP |
|--------|--------|
| Cassandra | Tunable consistency, defaults to eventual; always accepts writes |
| DynamoDB | Multi-region, always writable, conflict resolution via LWW |
| CouchDB | Multi-master replication, eventual consistency |
| DNS | Returns cached records even if authoritative server is unreachable |

## PACELC Extension
> If **P**artition → choose **A** or **C**; **E**lse (normal operation) → choose **L**atency or **C**onsistency

| System | PAC | ELC | Meaning |
|--------|-----|-----|---------|
| Cassandra | PA | EL | Highly available, low latency, eventual consistency |
| MongoDB | PC | EC | Consistent always, higher latency for consensus |
| DynamoDB | PA | EL | Available + fast, eventual consistency in normal ops |
| Spanner | PC | EC | Globally consistent via TrueTime, higher latency |

## Eventual Consistency
- After all writes stop, replicas will converge to the same state
- Convergence time depends on replication lag (typically milliseconds)
- Conflict resolution strategies:
  - **Last-Write-Wins (LWW)**: Timestamp-based, simple but lossy
  - **Vector Clocks**: Track causality, detect conflicts for manual resolution
  - **CRDTs**: Conflict-free data types that merge automatically (counters, sets)

## Key Interview Points
- CAP is about behavior **during partitions** in normal operation, you can have all three
- Tunable consistency (Cassandra): `ONE`, `QUORUM`, `ALL` trade latency for consistency
- Strong consistency ≠ ACID transactions (CAP is about distributed reads/writes)
- Real systems are nuanced: MongoDB is CP for writes but can read from secondaries (AP-like)
- Design for the failure mode: What does your app do during a partition?
