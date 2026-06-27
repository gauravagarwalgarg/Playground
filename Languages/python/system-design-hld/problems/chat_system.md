# Chat System (WhatsApp/Messenger)

## Requirements

### Functional
- 1:1 messaging and group chat
- Online/offline presence indicator
- Read receipts (sent/delivered/read)
- Push notifications for offline users
- Message history and search

### Non-Functional
- Real-time delivery (<100ms same-region)
- High availability (99.99%)
- Message ordering guaranteed per conversation
- Support 500M+ daily active users

## Real-Time Communication WebSocket

### Why WebSocket?
- Full-duplex persistent connection
- Server can push messages instantly (no polling)
- Low overhead after handshake (vs HTTP long-polling)

### Connection Management
```
Client → WebSocket Gateway (stateful) → Chat Service → Storage
```
- Gateway maintains active connections in memory
- Connection registry: Redis mapping `user_id → gateway_server_id`
- Heartbeat every 30s to detect dead connections
- Fallback: Long-polling for environments blocking WebSocket

### Scaling WebSocket Servers
- Sticky sessions via user_id hash to gateway
- Horizontal scaling: Add gateway servers, each handles ~100K connections
- Service mesh routes messages to correct gateway using connection registry

## Message Flow

### 1:1 Chat
```
1. User A sends message via WebSocket to their gateway
2. Chat service stores message in DB
3. Lookup User B's gateway in connection registry
4. If online: Forward via User B's WebSocket gateway
5. If offline: Queue for push notification
```

### Group Chat (Fan-out)
- Small groups (<100 members): Fan-out on write
  - Message written once, delivered to each member's inbox
- Large groups (>100): Fan-out on read
  - Message stored in group channel, members pull on open
- Hybrid: Fan-out on write for active members, store for inactive

## Message Storage

### Data Model (Cassandra Wide Column)
```
Partition Key: conversation_id
Clustering Key: message_id (time-ordered UUID)

| conversation_id | message_id | sender_id | content | timestamp | status |
```
- Cassandra: Optimized for write-heavy, ordered time-series data
- Messages are immutable append-only pattern
- Partition per conversation ensures data locality
- Retention: Keep recent messages hot; archive old messages to cold storage

### Why Not SQL?
- Write volume too high for single-master relational DB
- Access pattern is sequential reads within a conversation (perfect for Cassandra)
- No complex joins needed

## Online Presence (Heartbeat)

### Design
- Client sends heartbeat every 30 seconds
- Server updates `last_seen` timestamp in Redis
- Status: Online (heartbeat within 30s), Away (30s-5min), Offline (>5min)

### Fan-out Presence Updates
- Subscribe to friends' presence via pub/sub channel
- Only notify friends who are currently online (avoid unnecessary work)
- Group presence: Only fetch on demand (when user opens group info)

## Read Receipts
```
Message states: Sent ✓ → Delivered ✓✓ → Read (blue ✓✓)
```
- **Sent**: Server acknowledged receipt (ACK back to sender)
- **Delivered**: Recipient's device received the message
- **Read**: Recipient opened the conversation

Implementation:
- Each state change = lightweight status update message
- Batch receipt updates (don't send one update per message)
- Store last-read message_id per user per conversation

## Push Notifications
- Triggered when recipient is offline (not connected to WebSocket)
- Use APNs (iOS) / FCM (Android) for device push
- Aggregate: "3 new messages from Alice" instead of 3 separate pushes
- Respect user notification preferences and quiet hours

## Key Interview Discussion Points

### Message Ordering
- Per-conversation ordering via monotonically increasing message_id
- Server-assigned timestamps (not client) to handle clock skew
- Vector clocks for conflict detection in multi-device scenarios

### End-to-End Encryption
- Signal Protocol (used by WhatsApp): Sender encrypts, only recipient decrypts
- Server stores ciphertext only cannot read messages
- Key exchange: Diffie-Hellman on device pairing

### Multi-Device Sync
- Each device maintains its own WebSocket connection
- Message delivered to all active devices of a user
- Sync cursor: Last synced message_id per device

### Scalability Numbers
```
500M DAU × 40 messages/day = 20B messages/day
Storage: 20B × 100 bytes = 2 TB/day
QPS: 20B / 86400 ≈ 230K messages/sec
WebSocket connections: 500M concurrent peak ≈ 5000 gateway servers
```
