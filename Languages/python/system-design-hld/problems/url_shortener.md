# URL Shortener (TinyURL)

## Requirements

### Functional
- Given a long URL, generate a short unique URL
- Redirect short URL to original long URL
- Optional: Custom aliases, expiration, analytics

### Non-Functional
- Read-heavy (100:1 read-to-write ratio)
- Low latency redirects (<100ms)
- High availability (URLs should never be "down")
- 100M URLs created/month, 10B redirects/month

## Back-of-Envelope
```
Writes: 100M/month ≈ 40 writes/sec
Reads: 10B/month ≈ 4,000 reads/sec
Storage (5 years): 100M × 60 months × 500 bytes ≈ 3 TB
Cache (20% hot URLs): 10B × 20% / 30 days × 500 bytes ≈ 30 GB
```

## API Design

```
POST /api/v1/shorten
  Body: { "long_url": "https://...", "custom_alias": "optional", "expiry": "optional" }
  Response: { "short_url": "https://short.ly/abc123" }

GET /{short_code}
  Response: 301 Redirect (cacheable) or 302 (analytics-friendly)
```

## Encoding Strategy

### Base62 Encoding
- Characters: `[a-zA-Z0-9]` → 62 possible chars per position
- 7 characters → 62^7 ≈ 3.5 trillion unique URLs
- Counter-based: Global auto-increment ID → base62 encode
- Hash-based: MD5/SHA256(long_url) → take first 7 chars of base62

### ID Generation Options
| Approach | Pros | Cons |
|----------|------|------|
| Auto-increment DB | Simple, unique | Single point of failure, predictable |
| UUID | No coordination | Too long (128-bit), wasteful |
| Snowflake ID | Distributed, sorted | Complex setup |
| Pre-generated key service | Fast, no collision | Extra infra to maintain |

### Collision Handling
- Check if generated short code exists in DB
- If collision: append random char or rehash
- Pre-generated key service avoids collisions entirely

## High-Level Architecture

```
Client → Load Balancer → Web Servers → Cache (Redis) → Database
                                              ↓
                                     Key Generation Service
```

## Storage Design

### Database Schema
```sql
urls:
  id BIGINT PRIMARY KEY
  short_code VARCHAR(7) UNIQUE INDEX
  long_url TEXT
  created_at TIMESTAMP
  expires_at TIMESTAMP
  user_id BIGINT
```

- NoSQL (DynamoDB/Cassandra) works well: Simple key-value access pattern
- Partition key: `short_code` for even distribution

## Caching Layer
- Cache hot URLs in Redis (key: short_code, value: long_url)
- LRU eviction, TTL-based expiry
- Cache-aside pattern: Check cache → miss → query DB → populate cache
- 80/20 rule: Cache top 20% URLs to serve 80% of traffic

## Analytics
- Async event logging: short_code, timestamp, IP, user-agent, referrer
- Write to Kafka → consume into analytics DB (ClickHouse, BigQuery)
- Dashboard: Click counts, geographic distribution, time-series

## Rate Limiting
- Per-user/IP rate limits on URL creation (e.g., 100/hour)
- Token bucket at API gateway level
- Prevents abuse and spam link generation

## Key Interview Discussion Points
- 301 vs 302: 301 is permanent (browser caches), 302 allows tracking each redirect
- Read-heavy → cache-first architecture, eventually consistent reads are fine
- Key generation service can pre-generate and store in a pool for zero-collision writes
- Multi-region: DNS-based routing to nearest datacenter, global DB replication
- Cleanup: TTL-based expiration, periodic job to purge expired URLs
