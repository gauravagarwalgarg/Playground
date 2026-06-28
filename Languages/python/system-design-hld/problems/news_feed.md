# News Feed System (Facebook/Twitter)

## Requirements
- Users see posts from people they follow, ranked by relevance
- Support text, images, videos
- Real-time updates for new posts
- Handle celebrity accounts (millions of followers)
- Pagination for infinite scroll

## Fan-Out Strategies

### Fan-Out on Write (Push Model)
```
User posts → Write to poster's timeline + ALL followers' feed caches
```
- Pre-compute feeds: When user opens app, feed is ready instantly
- Pros: Fast reads (O(1) just fetch pre-built feed)
- Cons: Expensive writes for celebrities (1 post → millions of writes)
- Good for: Users with moderate follower count (<10K)

### Fan-Out on Read (Pull Model)
```
User opens feed → Fetch latest posts from all followed users → Merge & rank
```
- Compute feed at read time
- Pros: No write amplification, always fresh
- Cons: Slow reads (must query N sources + merge + rank)
- Good for: Celebrity accounts, inactive users

### Hybrid Approach (Industry Standard)
```
Regular users → Fan-out on write (push to followers' feeds)
Celebrities (>100K followers) → Fan-out on read (pull at read time)
```
- Follower's feed = pre-computed feed + merged celebrity posts at read time
- Twitter/Instagram use this hybrid model

## Feed Ranking

### Signals for Ranking
- Recency (newer posts score higher)
- Engagement prediction (like/comment/share probability)
- Relationship strength (interaction frequency with poster)
- Content type preference (user prefers videos over text)
- Diversity (don't show 5 posts from same person consecutively)

### Ranking Pipeline
```
Candidate Generation → Pre-ranking (lightweight model) → Full Ranking (ML model) → Re-ranking (diversity, ads)
```

## The Celebrity Problem
- Celebrities: Millions of followers, frequent posts
- Fan-out on write is impractical (1 post = 100M feed writes)
- Solution:
  1. Don't fan-out celebrity posts on write
  2. At read time: Merge user's pre-built feed + latest celebrity posts
  3. Cache celebrity posts separately (they're requested millions of times)
  4. Celebrity post cache: Small set of recent posts, high hit ratio

## Caching Hot Feeds

### Multi-Level Cache
```
L1: In-memory cache on web server (most recent feed, per-user)
L2: Redis cluster (pre-computed feeds per user)
L3: Database (permanent storage)
```

### Cache Strategy
- Cache the first 2-3 pages of each active user's feed
- Pre-warm feeds for users logging in (predict from patterns)
- Invalidation: New post from followed user → append to follower caches
- TTL: Short (5 min) to balance freshness vs compute cost

## Pagination Cursor-Based

### Why Not Offset?
- New posts shift the offset → user sees duplicate posts
- Large offsets = expensive DB queries

### Cursor-Based Pagination
```
GET /feed?cursor=eyJ0IjoxNjk4MDAwMDAwfQ&limit=20

Response: {
  "posts": [...],
  "next_cursor": "eyJ0IjoxNjk3OTk5MDAwfQ"
}
```
- Cursor encodes last seen timestamp or post_id
- Stable: New posts don't affect pagination of older content
- Efficient: Index scan from cursor position

## Storage Design

### Posts Table
```
post_id (Snowflake ID) | user_id | content | media_urls | created_at
```

### Feed Table (Pre-computed)
```
user_id (partition key) | post_id (sort key) | poster_id | created_at
```
- Cassandra/DynamoDB: Partition by user_id, sorted by time
- Keep only last N posts per user feed (e.g., 1000)

## Key Interview Points
- Trade-off: Write amplification (push) vs read latency (pull)
- Social graph service: Maintains follower/following relationships
- Media handling: Upload to S3/CDN, store URLs in post metadata
- Real-time updates: WebSocket/SSE for "new posts available" indicator
- Ads injection: Insert sponsored posts at rank positions (every Nth post)
- Unfollow/block: Remove from pre-computed feed, filter at read time
