# Rate Limiter System

## Requirements
- Limit requests per user/IP within a time window
- Low latency (in-path for every request)
- Distributed (works across multiple servers)
- Configurable rules (different limits per API endpoint)
- Return 429 Too Many Requests when exceeded

## Algorithms

### Token Bucket
- Bucket holds N tokens; refills at rate R tokens/sec
- Each request consumes 1 token; rejected if bucket empty
- Allows bursts up to bucket size
- Pros: Simple, memory efficient, allows controlled bursts
- Cons: Tuning bucket size vs refill rate
- Used by: AWS API Gateway, Stripe

### Fixed Window Counter
- Divide time into fixed windows (e.g., 1-minute intervals)
- Count requests per window; reject when limit exceeded
- Pros: Simple, low memory
- Cons: Boundary problem 2x traffic possible at window edges
  - Example: 100 requests at 0:59, 100 more at 1:00 = 200 in 2 seconds

### Sliding Window Log
- Store timestamp of each request in a sorted set
- On new request: Remove timestamps older than window, count remaining
- Pros: Precise, no boundary issue
- Cons: High memory (stores every timestamp)
- Implementation: Redis ZSET with `ZRANGEBYSCORE` + `ZCARD`

### Sliding Window Counter
- Hybrid: Combine fixed window counts with weighted overlap
- `count = prev_window_count × overlap_% + current_window_count`
- Pros: Low memory, smooth rate limiting
- Cons: Approximation (not 100% precise)
- Best balance of accuracy and efficiency

## Architecture

```
Client → API Gateway → Rate Limiter → Application
                           ↓
                    Redis (counters/tokens)
                           ↓
                    Rules Engine (config DB)
```

## Distributed Rate Limiting with Redis

### Token Bucket (Redis Lua Script)
```lua
-- Atomic token bucket in Redis
local tokens = redis.call('GET', KEYS[1]) or MAX_TOKENS
local last_refill = redis.call('GET', KEYS[2]) or NOW
-- Calculate tokens to add based on elapsed time
-- Deduct 1 token if available, reject otherwise
```

### Sliding Window (Redis Sorted Set)
```
MULTI
  ZREMRANGEBYSCORE key 0 (now - window_size)
  ZADD key now now
  ZCARD key
EXEC
-- If ZCARD > limit → reject and ZREM the just-added entry
```

### Considerations
- Race conditions: Use Lua scripts or Redis transactions for atomicity
- Clock sync: Use Redis server time, not client time
- Replication lag: Accept slight over-limit in multi-region setups

## Configuration
```yaml
rules:
  - endpoint: "/api/v1/messages"
    limit: 100
    window: 60s
    key: "user_id"
  - endpoint: "/api/v1/login"
    limit: 5
    window: 300s
    key: "ip_address"
```

## Key Interview Points
- Where to place: API gateway (edge) vs application middleware vs both
- Hard vs soft limits: Hard = reject; Soft = allow but log/alert
- Response headers: `X-RateLimit-Limit`, `X-RateLimit-Remaining`, `X-RateLimit-Reset`
- Graceful degradation: Queue excess requests instead of immediate reject
- Client-side: Exponential backoff with jitter on 429 responses
- Multi-tier: Different limits for free/paid/enterprise tiers
