# System Design - High Level Design (HLD)

## Fundamentals

### Key Components
- **Load Balancer**: Distributes traffic (Round Robin, Least Connections, IP Hash)
- **CDN**: Cache static assets at edge (CloudFront, Cloudflare)
- **Database**: SQL (ACID, joins) vs NoSQL (scale, flexibility)
- **Cache**: Redis/Memcached for hot data (Cache-aside, Write-through, Write-back)
- **Message Queue**: Async processing (Kafka, RabbitMQ, SQS)
- **API Gateway**: Rate limiting, auth, routing

### Scaling Patterns
- **Horizontal**: Add more machines (stateless services)
- **Vertical**: Bigger machine (databases, initially)
- **Sharding**: Partition data across nodes (by user_id, geo, etc.)
- **Replication**: Read replicas for read-heavy workloads
- **CQRS**: Separate read and write models

### CAP Theorem
- **CP**: Strong consistency, partition tolerant (ZooKeeper, etcd)
- **AP**: Available, eventually consistent (Cassandra, DynamoDB)
- **CA**: Not possible in distributed systems (single-node only)

### Estimation Cheat Sheet
| Metric | Value |
|--------|-------|
| 1 day | 86,400 seconds |
| 1 million req/day | ~12 req/sec |
| 1 billion req/day | ~12,000 req/sec |
| 1 KB * 1M users | 1 GB |
| 1 MB * 1M users | 1 TB |

---

## Problem Templates

See `problems/` for full designs:
- URL Shortener
- YouTube / Video Streaming
- Twitter / News Feed
- Chat System (WhatsApp)
- Rate Limiter
- Notification System
