# System Design: URL Shortener (bit.ly)

## Requirements
- Shorten a long URL to a short URL
- Redirect short URL to original
- Analytics (click count, geo)
- High availability, low latency

## Estimation
- 100M URLs/day created = ~1200/sec write
- 10:1 read:write = 12,000/sec read
- 5 years storage: 100M * 365 * 5 = 182B URLs
- Each URL ~500 bytes = 91 TB

## API
```
POST /api/shorten { "url": "https://..." } -> { "short": "abc123" }
GET  /:shortCode -> 301 Redirect
```

## Key Generation
- Base62 encoding (a-z, A-Z, 0-9) = 62^7 = 3.5 trillion combinations
- Options: Counter + Base62, MD5 hash first 7 chars, Pre-generated key service

## Architecture
```
Client -> API Gateway -> App Server -> DB (write)
Client -> CDN/Cache -> App Server -> Cache (Redis) -> DB (read)
```

## Database
- Key-value store (DynamoDB, Cassandra) -- simple lookup by shortCode
- Schema: shortCode (PK), originalUrl, createdAt, expiresAt, userId

## Cache
- Redis with shortCode -> originalUrl
- LRU eviction, 20% hot URLs serve 80% traffic

## Scaling
- Stateless app servers behind load balancer
- Database sharding by first char of shortCode
- Read replicas for analytics queries
