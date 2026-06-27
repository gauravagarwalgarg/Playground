# System Design Interview Template

A step-by-step framework for answering system design questions in 45-60 minutes.

## 1. Requirements Clarification (5 min)

### Functional Requirements
- What are the core features? (List 3-5 max)
- Who are the users? (End users, internal services, admins)
- What are the input/output formats?

### Non-Functional Requirements
- Scale: DAU, QPS (reads vs writes), data volume
- Latency: p50, p99 expectations
- Availability: 99.9%? 99.99%?
- Consistency: Strong vs eventual? Where does it matter?
- Durability: Can we lose data? RPO/RTO?

## 2. Back-of-Envelope Estimation (3 min)

```
Users:           ___M DAU
Reads:           ___K QPS
Writes:          ___K QPS
Storage/day:     ___GB
Storage/5yr:     ___TB
Bandwidth:       ___MB/s
```

Key formula: QPS = DAU × requests_per_user / 86400

## 3. API Design (5 min)

Define the top 3-5 API endpoints:
```
POST   /api/v1/resource
GET    /api/v1/resource/{id}
GET    /api/v1/resource?filter=X&page=N
DELETE /api/v1/resource/{id}
```

Consider: Auth, pagination, rate limiting, versioning.

## 4. Data Model (5 min)

- Identify main entities and relationships
- Choose storage: SQL vs NoSQL vs both
- Define key schemas and indexes
- Consider partitioning/sharding strategy

## 5. High-Level Design (10 min)

Draw the architecture:
- Client → Load Balancer → API Servers → Cache → DB
- Message queues for async processing
- CDN for static content
- Search index if needed

## 6. Deep Dive (15 min)

Pick 2-3 components to detail:
- How does the cache layer work? (Write-through, write-back, TTL)
- How do we shard the database? (Hash-based, range-based)
- How do we handle failures? (Circuit breaker, retry, DLQ)
- How does the notification system work?

## 7. Bottlenecks & Tradeoffs (5 min)

- Single points of failure → Replication, failover
- Hot partitions → Consistent hashing, load balancing
- Data consistency → Eventual consistency, CRDTs, saga pattern
- Cost vs performance tradeoffs

## Tips

- **Drive the conversation** don't wait for prompts
- **State assumptions** explicitly and ask if they're reasonable
- **Quantify** everything (QPS, storage, latency)
- **Trade off** there's no perfect design, show you understand costs
- **Start simple** and layer complexity as needed
