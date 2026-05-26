# API Design

## Paradigms

### REST (Representational State Transfer)
- **Resources**: Nouns, not verbs (`/users/123`, not `/getUser`)
- **HTTP Methods**: GET (read), POST (create), PUT (replace), PATCH (update), DELETE
- **Status Codes**: 2xx success, 4xx client error, 5xx server error
- **HATEOAS**: Hypermedia links for discoverability
- **Best for**: CRUD operations, public APIs, browser clients

### GraphQL
- **Single endpoint**: Client specifies exactly what data it needs
- **Schema-first**: Strongly typed schema defines API contract
- **Resolvers**: Functions that fetch data for each field
- **Best for**: Complex data relationships, mobile clients (bandwidth), rapid frontend iteration
- **Trade-offs**: Caching complexity, N+1 queries, learning curve

### gRPC
- **Protocol Buffers**: Binary serialization, schema-defined
- **HTTP/2**: Multiplexing, streaming, header compression
- **Code generation**: Client/server stubs from .proto files
- **Best for**: Internal microservices, low-latency, streaming
- **Trade-offs**: Not browser-friendly (needs proxy), binary format harder to debug

### Versioning Strategies
| Strategy | Example | Pros | Cons |
|----------|---------|------|------|
| URL path | `/v1/users` | Simple, explicit | URL pollution |
| Header | `Accept: application/vnd.api.v1+json` | Clean URLs | Hidden |
| Query param | `/users?version=1` | Easy to test | Caching issues |

---

## Design Principles

1. **Consistency**: Same patterns across all endpoints
2. **Pagination**: Cursor-based (scalable) or offset-based (simple)
3. **Filtering**: Query parameters for common filters
4. **Error handling**: Structured error responses with codes and messages
5. **Rate limiting**: Protect services, communicate limits via headers
6. **Idempotency**: Safe retries (idempotency keys for POST)

---

## Interview Talking Points

- Choose paradigm based on use case: REST for public, gRPC for internal, GraphQL for complex frontends
- Backward compatibility is critical new_textnever break existing clients
- API is a contract new_textdesign it like a product with clear documentation
