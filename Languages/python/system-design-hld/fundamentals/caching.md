# Caching

## Cache Write Strategies

### Write-Through
- Write to cache AND database simultaneously
- Pro: Cache always consistent with DB
- Con: Higher write latency (two writes per operation)
- Use: Systems where read consistency is critical

### Write-Back (Write-Behind)
- Write to cache only; async flush to DB in batches
- Pro: Low write latency, reduced DB load
- Con: Data loss risk if cache node crashes before flush
- Use: Write-heavy workloads (analytics, logging)

### Write-Around
- Write directly to DB, bypass cache
- Cache populated only on read (miss → fetch from DB → cache)
- Pro: Cache not polluted by write-once data
- Con: First read after write is always a cache miss
- Use: Infrequently read data

## Cache Read Strategies

### Cache-Aside (Lazy Loading)
- App checks cache → miss → reads DB → writes to cache
- App manages cache explicitly
- Pro: Only requested data is cached
- Con: Cache miss penalty, stale data possible

### Read-Through
- Cache itself fetches from DB on miss (transparent to app)
- Pro: Simpler application code
- Con: Cold start latency, cache library must support it

## Eviction Policies

| Policy | Evicts | Best For |
|--------|--------|----------|
| LRU | Least Recently Used | General purpose, temporal locality |
| LFU | Least Frequently Used | Stable hot datasets |
| FIFO | First In First Out | Simple, time-based expiry |
| TTL | Time-based expiry | Session data, tokens |

## CDN Caching
- Cache static assets at edge locations (images, JS, CSS, videos)
- Reduces origin server load and latency for global users
- Cache-Control headers: `max-age`, `s-maxage`, `no-cache`, `no-store`
- Invalidation: Purge API, versioned URLs (`style.v2.css`)
- Pull CDN (lazy) vs Push CDN (pre-populate)

## Redis vs Memcached

| Feature | Redis | Memcached |
|---------|-------|-----------|
| Data structures | Strings, Lists, Sets, Hashes, Sorted Sets | Strings only |
| Persistence | RDB snapshots + AOF | None |
| Replication | Master-replica | None (client-side) |
| Pub/Sub | Yes | No |
| Memory efficiency | Higher overhead per key | Better for simple k/v |
| Use case | Feature-rich caching, sessions, leaderboards | Simple high-throughput caching |

## Thundering Herd Problem
- Many concurrent requests for the same expired/missing key hit DB simultaneously
- Solutions:
  - **Locking**: First request acquires lock, others wait for cache fill
  - **Request coalescing**: Collapse duplicate in-flight requests
  - **Stale-while-revalidate**: Serve stale data while refreshing async
  - **Early expiration jitter**: Randomize TTL to avoid simultaneous expiry

## Key Interview Points
- Cache hit ratio = hits / (hits + misses) aim for >95%
- Cache warming: Pre-populate on deploy to avoid cold start
- Distributed cache: Partition keys across nodes (consistent hashing)
- Cache invalidation is the hardest problem prefer TTL + eventual consistency
