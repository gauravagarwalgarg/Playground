# API Design

## REST Principles
- **Resource-based URLs**: `/users/123/orders` not `/getUserOrders?id=123`
- **HTTP methods**: GET (read), POST (create), PUT (full update), PATCH (partial), DELETE
- **Stateless**: Each request contains all info needed; no server-side session
- **HATEOAS**: Responses include links to related actions (rarely implemented in practice)
- **Idempotency**: GET, PUT, DELETE are idempotent; POST is not

## Versioning Strategies

| Strategy | Example | Pros/Cons |
|----------|---------|-----------|
| URL path | `/v1/users` | Simple, explicit; pollutes URL |
| Header | `Accept: application/vnd.api.v1+json` | Clean URLs; harder to test |
| Query param | `/users?version=1` | Easy; optional param confusion |

Best practice: URL path versioning (`/v1/`, `/v2/`) most common in industry.

## Pagination

### Offset-Based
```
GET /posts?offset=20&limit=10
```
- Simple but breaks when data changes (inserts shift pages)
- Performance degrades at large offsets (DB scans N rows to skip)

### Cursor-Based (Recommended)
```
GET /posts?cursor=eyJpZCI6MTAwfQ&limit=10
Response: { "data": [...], "next_cursor": "eyJpZCI6MTEwfQ" }
```
- Stable pagination unaffected by concurrent inserts/deletes
- Cursor = opaque encoded value (usually last seen ID or timestamp)

## Rate Limiting
- Protect APIs from abuse and ensure fair usage
- Common limits: 100 req/min (free tier), 1000 req/min (paid)
- Response headers: `X-RateLimit-Limit`, `X-RateLimit-Remaining`, `X-RateLimit-Reset`
- Return `429 Too Many Requests` when exceeded
- Algorithms: Token bucket, sliding window (see rate_limiter_system.md)

## Idempotency
- Allow safe retries without side effects
- Client sends `Idempotency-Key` header with unique request ID
- Server stores result of first execution, returns cached result on retry
- Critical for: Payment APIs, order creation, any non-idempotent mutation
- Implementation: Store key → response mapping in Redis with TTL

## GraphQL vs REST vs gRPC

```
Need flexible queries from frontend?
├── Yes → GraphQL
└── No →
    Internal service-to-service?
    ├── Yes → gRPC (performance, type safety, streaming)
    └── No →
        Public API with simple CRUD?
        ├── Yes → REST
        └── Need real-time bidirectional?
            └── Yes → gRPC streaming or WebSocket
```

| Feature | REST | GraphQL | gRPC |
|---------|------|---------|------|
| Transport | HTTP/1.1 | HTTP/1.1 | HTTP/2 |
| Format | JSON | JSON | Protobuf (binary) |
| Schema | OpenAPI (optional) | Strongly typed | .proto files |
| Overfetching | Common | Solved | N/A |
| Caching | HTTP caching (easy) | Complex | Not built-in |
| Use case | Public APIs, CRUD | Mobile/frontend | Microservices |

## Authentication & Authorization

### OAuth 2.0 Flows
- **Authorization Code**: Web apps (most secure, uses server-side exchange)
- **Client Credentials**: Service-to-service (no user involved)
- **PKCE**: Mobile/SPA (prevents code interception)

### JWT (JSON Web Token)
- Self-contained token: `header.payload.signature`
- Stateless verification no DB lookup needed
- Include: `sub` (user ID), `exp` (expiry), `roles`
- Short-lived access tokens (15 min) + long-lived refresh tokens
- Never store sensitive data in JWT payload (it's base64, not encrypted)

### API Key
- Simple, for server-to-server or public APIs
- Include in header: `X-API-Key: abc123`
- Not suitable for user-level auth (no identity context)

## Key Interview Points
- Always design APIs before implementation contract-first approach
- Use standard HTTP status codes: 200, 201, 400, 401, 403, 404, 429, 500
- Backward compatibility: Add fields, don't remove; deprecate with timeline
- Bulk operations: `POST /users/batch` for multiple creates
- Async operations: Return `202 Accepted` + polling endpoint or webhook
