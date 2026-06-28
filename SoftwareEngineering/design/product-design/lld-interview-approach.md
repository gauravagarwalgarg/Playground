# Low-Level Design Interview Approach

> Adapted from awesome-low-level-design. A structured framework for LLD/OOD interviews.

---

## Step-by-Step Framework (45 min interview)

### 1. Clarify Requirements (5 min)
- What are the core entities?
- What actions/operations are needed?
- What constraints exist? (concurrency, scale, real-time?)
- What's out of scope?

### 2. Identify Core Classes & Relationships (5 min)
- List the main classes (nouns in requirements)
- Identify relationships: has-a (composition), is-a (inheritance), uses (dependency)
- Note which classes will be interfaces vs concrete

### 3. Define Class Interfaces (10 min)
- Public methods and their signatures
- What each method does (not how)
- Enums for fixed-value types (status, type, direction)

### 4. Implement Core Logic (20 min)
- Start with the most important class
- Implement key methods
- Handle edge cases
- Add thread-safety where needed

### 5. Discuss Extensibility (5 min)
- What design patterns did you use and why?
- How would you extend for new requirements?
- What would you change for 10x scale?

---

## Design Patterns Cheat Sheet (When to Use)

| Pattern | Use When |
|---------|----------|
| **Singleton** | Global config, connection pools, logger |
| **Factory** | Object creation varies by type, hide instantiation logic |
| **Builder** | Complex objects with many optional parameters |
| **Strategy** | Multiple algorithms interchangeable at runtime |
| **Observer** | Notify multiple objects of state changes |
| **State** | Object behavior changes based on internal state |
| **Decorator** | Add behavior dynamically without subclassing |
| **Adapter** | Make incompatible interfaces work together |
| **Facade** | Simplify a complex subsystem behind one interface |
| **Command** | Encapsulate requests as objects (undo/redo, queue) |
| **Chain of Responsibility** | Pipeline/middleware processing |
| **Template Method** | Same algorithm structure, varying steps |
| **Composite** | Tree structures (files/folders, UI components) |
| **Proxy** | Control access, add caching/logging layer |

---

## Top LLD Interview Problems (by Difficulty)

### Easy (30 min)
- Parking Lot
- Vending Machine
- Logging Framework
- Traffic Signal System
- Task Management System

### Medium (45 min)
- LRU Cache
- Pub/Sub System
- Elevator System
- Online Shopping Cart
- Rate Limiter (Token Bucket)
- Hotel Booking System
- Library Management System

### Hard (60 min)
- Chess Game
- Splitwise (expense sharing)
- Ride-Sharing Service (Uber)
- Movie Ticket Booking (concurrency)
- Online Stock Brokerage
- Food Delivery Service

---

## SOLID Principles Quick Reference

| Principle | Meaning | Violation Signal |
|-----------|---------|-----------------|
| **S**ingle Responsibility | One class, one reason to change | Class doing too many things |
| **O**pen/Closed | Open for extension, closed for modification | Adding feature requires editing existing code |
| **L**iskov Substitution | Subtypes work wherever parent type is expected | Override breaks parent's contract |
| **I**nterface Segregation | Many small interfaces > one big interface | Implementing methods you don't need |
| **D**ependency Inversion | Depend on abstractions, not concretions | Hard-coded `new ConcreteClass()` everywhere |

---

## Concurrency Considerations

When the problem involves shared resources:
1. **Identify shared state** what's accessed by multiple threads?
2. **Choose synchronization** mutex, read-write lock, channels, CAS
3. **Minimize critical section** lock only what's necessary
4. **Consider lock-free alternatives** atomic operations, immutable data

Common concurrency-relevant LLD problems:
- Rate Limiter (token bucket with atomic ops)
- Pub/Sub (concurrent readers/writers)
- Parking Lot (spot reservation race conditions)
- Task Scheduler (worker pool + job queue)

---

## Code Organization Tips

```
problem/
├── models.go          # Data structures (Vehicle, ParkingSpot, etc.)
├── interfaces.go      # Contracts (PaymentProcessor, Notifier)
├── service.go         # Core business logic (ParkingLotService)
├── strategies.go      # Strategy implementations (pricing, allocation)
└── main.go            # Demo / test harness
```

For interview (single file):
- Start with models/enums at the top
- Then interfaces
- Then implementations
- Then main() with demo

---

## References
- [awesome-low-level-design](../../../awesome-low-level-design/README.md)
- [Go LLD Implementations](../../Languages/go/low-level-design-lld/)
- [Python LLD Implementations](../../Languages/python/low-level-design-lld/)
- Clean Code Robert C. Martin
- Head First Design Patterns Freeman & Robson
