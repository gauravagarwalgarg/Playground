# System Design Interview Checklist

> Adapted from awesome-system-design-resources. A structured framework for tackling any system design interview.

---

## Step 1: Requirements Clarification (3-5 min)

### Functional Requirements
- What are the core features?
- Who are the users? How many?
- What actions can users perform?
- What data do they create/consume?

### Non-Functional Requirements
- **Scale**: DAU, QPS, storage estimates
- **Latency**: What's acceptable? (p50, p99)
- **Availability**: 99.9%? 99.99%?
- **Consistency**: Strong or eventual?
- **Durability**: Can we lose data? For how long?

### Back-of-the-Envelope Estimation
```
Users: 100M DAU
Reads: 10 reads/user/day = 1B reads/day ≈ 12K QPS
Writes: 1 write/user/day = 100M writes/day ≈ 1.2K QPS
Storage: 100M users × 1KB/user = 100GB (metadata only)
```

---

## Step 2: High-Level Design (10 min)

Draw the 30,000-foot view:
```
Client → Load Balancer → API Gateway → Services → Database
                                           ↕
                                       Cache Layer
                                           ↕
                                     Message Queue → Workers
```

### Core Components to Consider
| Component | When to Use |
|-----------|-------------|
| Load Balancer | Multiple app servers, >1K QPS |
| API Gateway | Multiple services, auth, rate limiting |
| CDN | Static content, global users |
| Cache (Redis) | Read-heavy, hot data, sessions |
| Message Queue | Async processing, decoupling |
| Search (Elasticsearch) | Full-text search, analytics |
| Object Storage (S3) | Files, images, videos |

---

## Step 3: Deep Dive (15-20 min)

### Database Design
- **Schema design**: Tables, relationships, indices
- **SQL vs NoSQL**: Structured data → SQL; High scale, flexible schema → NoSQL
- **Sharding strategy**: By user_id, region, time
- **Replication**: Leader-follower for reads, multi-leader for geo-distribution

### API Design
```
POST /api/v1/messages
  Body: { "to": "user123", "content": "hello", "type": "text" }
  Response: { "id": "msg456", "status": "sent", "timestamp": "..." }

GET /api/v1/messages?conversation_id=conv789&limit=50&cursor=msg400
```

### Scaling Patterns
- **Horizontal scaling**: Stateless services behind LB
- **Caching**: Read-through, write-behind, cache-aside
- **Database scaling**: Read replicas → sharding → separate read/write DBs
- **Async processing**: Queue writes, process in background

### Data Flow
- Request path (happy path + failure scenarios)
- Write path vs Read path (often different in high-scale systems)
- How data propagates through the system

---

## Step 4: Bottlenecks & Trade-offs (5 min)

### Common Bottlenecks
| Bottleneck | Solution |
|-----------|----------|
| Single DB | Sharding, read replicas, caching |
| Hot partition | Consistent hashing + virtual nodes |
| Thundering herd | Request coalescing, staggered TTL |
| Large fan-out | Push vs Pull, hybrid approach |
| Global latency | Multi-region, CDN, edge computing |

### Trade-offs to Discuss
- Consistency vs Availability (CAP)
- Latency vs Throughput
- Cost vs Performance
- Complexity vs Reliability
- Push vs Pull architecture
- Sync vs Async processing

---

## Building Blocks Quick Reference

### Rate Limiting
- **Token Bucket**: Allows bursts, smooth average rate
- **Leaky Bucket**: Fixed output rate, smooths traffic
- **Fixed Window**: Simple, boundary burst problem
- **Sliding Window**: Most accurate, memory intensive

### Caching Strategies
- **Cache-Aside**: App manages cache (most common)
- **Read-Through**: Cache loads on miss
- **Write-Through**: Write to cache + DB simultaneously
- **Write-Behind**: Write to cache, async DB write (risk of loss)
- **Eviction**: LRU (most common), LFU, TTL-based

### Consistent Hashing
- Minimal remapping when nodes added/removed (~K/N keys move)
- Virtual nodes for even distribution (150+ per physical node)
- Used in: DynamoDB, Cassandra, CDNs, load balancers

### Message Queues
- **At-most-once**: Fire and forget (lowest latency)
- **At-least-once**: Retry until ACK (most common, requires idempotency)
- **Exactly-once**: Hardest (transactions + dedup)
- **Ordering**: Per-partition ordering (Kafka), no global ordering in most systems

---

## Common System Design Problems by Category

### Storage & Retrieval
- URL Shortener, Pastebin, Key-Value Store, Distributed File System

### Messaging & Real-time
- Chat System (WhatsApp), Notification Service, Pub/Sub

### Social & Feed
- Twitter/X, Instagram, Facebook News Feed, Reddit

### Media
- YouTube, Netflix, Spotify, TikTok

### E-commerce & Payments
- Amazon, Uber, Food Delivery, Payment System, Digital Wallet

### Search & Discovery
- Google Search, Autocomplete, Yelp (Location-based)

### Infrastructure
- Rate Limiter, Load Balancer, Distributed Cache, Job Scheduler, Web Crawler

---

## References
- [awesome-system-design-resources](../../../awesome-system-design-resources/README.md)
- [System Design Fundamentals (Go implementations)](../../Languages/go/system-design-hld/fundamentals/)
- Designing Data-Intensive Applications (DDIA) Martin Kleppmann
