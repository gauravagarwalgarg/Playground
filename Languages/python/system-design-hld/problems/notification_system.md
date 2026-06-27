# Notification System

## Requirements
- Multi-channel: Push (iOS/Android), SMS, Email
- Millions of notifications/day
- Priority-based delivery (urgent vs marketing)
- User preferences (opt-in/opt-out per channel)
- Template engine for consistent messaging
- Retry failed deliveries with exponential backoff

## High-Level Architecture

```
Trigger Source → Notification Service → Priority Queue → Workers → Delivery Services
                       ↓                                              ↓
                 Template Engine                              APNs / FCM / Twilio / SES
                       ↓
                 User Preferences DB
```

## Delivery Channels

### Push Notifications
- **iOS**: Apple Push Notification Service (APNs) requires device token
- **Android**: Firebase Cloud Messaging (FCM) requires registration token
- Device token stored per user; handle token refresh/invalidation
- Payload limit: APNs = 4 KB, FCM = 4 KB
- Silent push: Wake app without showing notification (data sync)

### SMS
- Providers: Twilio, AWS SNS, MessageBird
- Expensive per message reserve for critical alerts (OTP, security)
- Delivery receipt tracking (sent/delivered/failed)
- Country-specific regulations (opt-in requirements, quiet hours)

### Email
- Providers: AWS SES, SendGrid, Mailgun
- High volume, low cost
- Track: Open rate, click rate, bounce rate, spam complaints
- SPF/DKIM/DMARC configuration for deliverability

## Priority Queuing
```
High Priority Queue   → OTP, security alerts, transaction confirmations
Medium Priority Queue → Social interactions, updates
Low Priority Queue    → Marketing, weekly digests, recommendations
```
- Separate queues per priority level
- High-priority workers always have capacity reserved
- Low-priority can be throttled during peak load

## Template Engine
- Templates stored in DB with variable placeholders
- Example: `"Hello {{user_name}}, your order #{{order_id}} is shipped!"`
- Supports: Localization (i18n), A/B testing variants
- Rendering at send time not at queue time (fresh data)

## User Preferences
```json
{
  "user_id": "123",
  "channels": {
    "push": { "enabled": true, "quiet_hours": "22:00-08:00" },
    "email": { "enabled": true, "frequency": "daily_digest" },
    "sms": { "enabled": false }
  },
  "categories": {
    "marketing": false,
    "security": true,
    "social": true
  }
}
```
- Check preferences before queuing (early filter saves resources)
- Respect quiet hours delay non-urgent notifications

## Retry with Exponential Backoff
```
Attempt 1: Immediate
Attempt 2: Wait 1s
Attempt 3: Wait 2s
Attempt 4: Wait 4s
Attempt 5: Wait 8s (cap at 5 retries)
```
- Add jitter to prevent thundering herd on retries
- Different retry policies per channel (APNs rate limits vs email)
- After max retries → Dead Letter Queue for manual inspection
- Track delivery status: `pending → sent → delivered → failed`

## Deduplication
- Prevent sending same notification twice (exactly-once semantics)
- Idempotency key: hash(user_id + event_type + event_id)
- Store recent notification IDs in Redis with short TTL

## Key Interview Points
- Rate limiting per user: Max N notifications/hour to prevent spam
- Analytics pipeline: Kafka → ClickHouse for delivery metrics
- Fan-out: Group event → millions of notifications; use async workers
- Unsubscribe handling: One-click unsubscribe link (email), in-app toggle
- Failure isolation: One channel failure shouldn't block others
- Scale: Partition workers by channel; push/email/SMS scale independently
