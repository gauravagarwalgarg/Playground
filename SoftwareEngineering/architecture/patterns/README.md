# Architectural Patterns

## Patterns Overview

### 1. Microservices Architecture
- **What**: Decompose application into small, independently deployable services
- **When**: Large teams, independent scaling needs, polyglot tech stacks
- **Trade-offs**: Operational complexity, distributed data management, network latency
- **Key concepts**: Service discovery, API gateway, circuit breaker, saga pattern

### 2. Event-Driven Architecture (EDA)
- **What**: Components communicate through events (async messages)
- **When**: Decoupled systems, real-time processing, audit trails
- **Trade-offs**: Eventual consistency, debugging complexity, message ordering
- **Key concepts**: Event sourcing, CQRS, message brokers (Kafka, RabbitMQ)

### 3. CQRS (Command Query Responsibility Segregation)
- **What**: Separate read and write models
- **When**: Read-heavy workloads, complex domains, different scaling needs for reads vs writes
- **Trade-offs**: Increased complexity, eventual consistency between models
- **Key concepts**: Command handlers, query handlers, event store, projections

### 4. Hexagonal Architecture (Ports & Adapters)
- **What**: Core business logic is isolated from external concerns via ports (interfaces) and adapters (implementations)
- **When**: Testability is critical, multiple integration points, long-lived systems
- **Trade-offs**: More boilerplate, over-engineering for simple CRUD
- **Key concepts**: Ports (interfaces), adapters (implementations), domain core

### 5. Layered Architecture
- **What**: Organize code into horizontal layers (Presentation → Business → Data)
- **When**: Simple applications, small teams, well-understood domains
- **Trade-offs**: Tight coupling between layers, monolithic tendencies
- **Key concepts**: Separation of concerns, dependency direction

### 6. Clean Architecture
- **What**: Concentric circles with dependencies pointing inward (Entities → Use Cases → Interface Adapters → Frameworks)
- **When**: Complex business logic, long-term maintainability
- **Trade-offs**: Verbose, many abstractions, steep learning curve
- **Key concepts**: Dependency inversion, use cases, entities, interface adapters

---

## Interview Talking Points

- No architecture is universally "best" new_textit depends on team size, domain complexity, and scaling needs
- Start simple (monolith), extract services when boundaries become clear
- Event-driven and CQRS often go together but are independent concepts
- Hexagonal and Clean Architecture share the same core idea: isolate business logic from infrastructure
