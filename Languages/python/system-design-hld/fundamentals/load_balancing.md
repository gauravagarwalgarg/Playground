# Load Balancing

## Algorithms

### Round Robin
- Requests distributed sequentially across servers
- Simple, works well when servers have equal capacity
- No awareness of server load or health

### Weighted Round Robin
- Assigns weights based on server capacity (e.g., 8-core gets 4x traffic of 2-core)
- Use when servers have heterogeneous hardware
- Weight adjustment can be static or dynamic

### Least Connections
- Routes to server with fewest active connections
- Best for long-lived connections (WebSocket, database connections)
- Handles uneven request durations well

### IP Hash
- Hash of client IP determines server
- Guarantees session affinity (same client → same server)
- Problem: uneven distribution if IP ranges are clustered

### Consistent Hashing
- Servers placed on a hash ring; requests map to nearest server clockwise
- Adding/removing a server only affects adjacent keys (minimal redistribution)
- Used by: CDNs, distributed caches (Memcached), Cassandra
- Virtual nodes solve hotspot issues

## When to Use Each

| Algorithm | Best For |
|-----------|----------|
| Round Robin | Stateless services, equal-capacity servers |
| Weighted RR | Mixed hardware, gradual rollouts |
| Least Connections | Long-lived connections, varied response times |
| IP Hash | Session persistence without cookies |
| Consistent Hashing | Distributed caches, sharded databases |

## L4 vs L7 Load Balancing

### Layer 4 (Transport)
- Operates on TCP/UDP packets
- Fast no payload inspection
- Cannot route based on URL, headers, or cookies
- Examples: AWS NLB, HAProxy (TCP mode)

### Layer 7 (Application)
- Operates on HTTP/HTTPS
- Can route by URL path, headers, cookies, body content
- Supports SSL termination, compression, caching
- Examples: AWS ALB, Nginx, Envoy

### Decision Guide
- Need content-based routing → L7
- Need raw throughput, minimal latency → L4
- Microservices with path-based routing → L7
- TCP passthrough (databases, gRPC) → L4

## Key Interview Points

- Health checks: Passive (monitor failures) vs Active (periodic pings)
- DNS-based LB: Route53 weighted/latency routing coarse-grained, TTL issues
- Global Server Load Balancing (GSLB): Geo-DNS routes users to nearest region
- Sticky sessions: Avoid if possible breaks horizontal scaling
- Scalability: LB itself can be a bottleneck use multiple LBs with DNS round-robin
