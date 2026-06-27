# Messaging Queues

## Kafka vs RabbitMQ vs SQS

| Feature | Kafka | RabbitMQ | SQS |
|---------|-------|----------|-----|
| Model | Distributed log | Message broker | Managed queue |
| Ordering | Per-partition | Per-queue | Best-effort (FIFO available) |
| Throughput | Millions msg/sec | ~50K msg/sec | Auto-scales |
| Retention | Configurable (days/weeks) | Until consumed | 14 days max |
| Replay | Yes (offset-based) | No | No |
| Use case | Event streaming, logs | Task queues, RPC | Simple decoupling, serverless |
| Ops burden | High (ZooKeeper/KRaft) | Medium | Zero (managed) |

## Messaging Patterns

### Pub/Sub (Publish-Subscribe)
- Producer publishes to a topic; all subscribers receive the message
- Use: Event broadcasting, notifications, real-time updates
- Examples: Kafka topics, SNS, Redis Pub/Sub

### Point-to-Point (Queue)
- Message consumed by exactly one consumer from the queue
- Use: Task distribution, work queues, job processing
- Examples: SQS, RabbitMQ queues, Celery

## Delivery Semantics

### At-Most-Once
- Fire and forget no retries
- Message may be lost but never duplicated
- Use: Metrics, logs where occasional loss is acceptable

### At-Least-Once
- Retry until acknowledged may produce duplicates
- Consumer must be idempotent (handle duplicate processing)
- Most common in production systems (Kafka default, SQS)

### Exactly-Once
- Hardest to achieve requires idempotent producers + transactional consumers
- Kafka supports via idempotent producer + transactional API
- Often approximated with at-least-once + deduplication

## Back-Pressure
- When consumers can't keep up with producers
- Strategies:
  - **Drop**: Discard excess messages (acceptable for metrics)
  - **Buffer**: Increase queue size (delays problem, doesn't solve it)
  - **Rate limit producers**: Reject/throttle at ingestion
  - **Scale consumers**: Auto-scale consumer group (Kafka consumer groups, SQS + Lambda)
  - **Circuit breaker**: Stop accepting when downstream is overwhelmed

## Dead Letter Queues (DLQ)
- Messages that fail processing after N retries are moved to a DLQ
- Prevents poison messages from blocking the queue
- DLQ messages can be inspected, fixed, and replayed
- Configure: Max receive count (SQS), max retries (Kafka)
- Alert on DLQ depth indicates systemic failures

## Key Interview Points
- Kafka partitions = unit of parallelism more partitions = more consumers
- Consumer groups in Kafka: Each partition consumed by one consumer in the group
- Message ordering: Kafka guarantees per-partition; SQS FIFO guarantees per-group
- Idempotency key: Attach unique ID to messages for deduplication
- Backlog monitoring: Track consumer lag (Kafka) or ApproximateNumberOfMessages (SQS)
