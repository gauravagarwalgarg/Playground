# Back-of-Envelope Estimation

## Powers of 2 Quick Reference

| Power | Value | Approx |
|-------|-------|--------|
| 2^10 | 1,024 | ~1 Thousand (1 KB) |
| 2^20 | 1,048,576 | ~1 Million (1 MB) |
| 2^30 | 1,073,741,824 | ~1 Billion (1 GB) |
| 2^40 | | ~1 Trillion (1 TB) |

## Common Scale Numbers

| Metric | Value |
|--------|-------|
| Seconds in a day | 86,400 (~100K) |
| Seconds in a month | ~2.5M |
| Seconds in a year | ~30M |
| 1M users, 10% DAU | 100K daily active |
| 100M requests/day | ~1,200 QPS |
| 1B requests/day | ~12,000 QPS |
| Peak traffic | 2-3x average QPS |

## Latency Numbers Every Programmer Should Know

| Operation | Latency |
|-----------|---------|
| L1 cache reference | 0.5 ns |
| L2 cache reference | 7 ns |
| Main memory (RAM) | 100 ns |
| SSD random read | 150 μs |
| HDD seek | 10 ms |
| Same datacenter round trip | 0.5 ms |
| Cross-continent round trip | 150 ms |
| Read 1 MB from memory | 250 μs |
| Read 1 MB from SSD | 1 ms |
| Read 1 MB from HDD | 20 ms |
| Send 1 MB over 1 Gbps network | 10 ms |

## Estimation Templates

### Storage Estimation
```
Users: 100M
Average data per user: 1 KB (profile) + 500 KB (media)
Total storage = 100M × 501 KB ≈ 50 TB
Growth: 10% YoY → plan for 5 years = 50 TB × 1.5 ≈ 75 TB
Replication factor 3 → 225 TB raw storage
```

### Bandwidth Estimation
```
100M requests/day × 50 KB avg response = 5 TB/day outbound
5 TB / 86,400 sec ≈ 60 MB/s sustained
Peak (3x) = 180 MB/s ≈ 1.5 Gbps
```

### QPS Estimation
```
DAU: 10M users
Each user makes 20 requests/day
Total = 200M requests/day
QPS = 200M / 86,400 ≈ 2,300 QPS
Peak = 2,300 × 3 ≈ 7,000 QPS
```

### Database Size Estimation
```
Each record: 500 bytes
New records/day: 1M
Daily growth: 1M × 500 B = 500 MB/day
Yearly: 500 MB × 365 ≈ 180 GB/year
With indexes (2x): 360 GB/year
5-year plan: 1.8 TB
```

## Quick Conversion Rules
- 1 day ≈ 100K seconds (actual: 86,400)
- 1 million requests/day ≈ 12 QPS
- 1 billion requests/day ≈ 12,000 QPS
- 1 GB/day ≈ 12 KB/sec
- Multiply average by 3 for peak estimation
- Char = 1 byte, Int = 4 bytes, Long/Timestamp = 8 bytes
- UUID = 16 bytes, URL = ~100 bytes, Tweet = ~300 bytes

## Interview Tips
- State assumptions clearly before calculating
- Round aggressively precision doesn't matter, order of magnitude does
- Always consider: storage, bandwidth, QPS, memory (cache sizing)
- Work from DAU → actions per user → total requests → QPS
- Factor in replication, redundancy, and growth (3-5 year horizon)
