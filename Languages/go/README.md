# Go - Interview Preparation

## How to Run

### Individual files
```bash
cd go
go run algorithms-dsa/patterns/heaps.go
go run low-level-design-lld/design-patterns/strategy.go
go run system-design-hld/fundamentals/consistent_hashing.go
```

### Run all in a category
```bash
find . -path "./algorithms-dsa/patterns/*.go" -exec go run {} \;
```

## Structure

```
go/
├── algorithms-dsa/
│   ├── patterns/
│   │   ├── two_pointers.go           # Opposite Ends, Fast/Slow, Same Direction
│   │   ├── sliding_window.go         # Fixed Window, Variable Window
│   │   ├── binary_search.go          # Classic, Rotated Array, Search Range
│   │   ├── dynamic_programming.go    # Knapsack, LCS, LIS
│   │   ├── backtracking.go           # Permutations, N-Queens, Subsets
│   │   ├── graph_traversal.go        # BFS, DFS basics
│   │   ├── graphs.go                 # BFS, DFS, Topological Sort, Dijkstra, Union-Find, Cycle Detection
│   │   ├── stacks.go                 # Monotonic Stack, Next Greater, Histogram, RPN
│   │   ├── heaps.go                  # Top-K, Merge-K-Sorted, Running Median
│   │   ├── intervals.go              # Merge, Insert, Meeting Rooms, Intersection
│   │   ├── prefix_sums.go            # Range Sum, Subarray Sum K, Product Except Self, 2D
│   │   ├── tries.go                  # Insert/Search, Autocomplete, Wildcard, LCP
│   │   └── bit_manipulation.go       # Single Number, Power of 2, Subsets, Hamming
│   └── leetcode/
│       └── arrays_hashing/
│
├── low-level-design-lld/
│   ├── design-patterns/
│   │   ├── singleton.go              # Thread-safe Singleton (sync.Once)
│   │   ├── factory.go                # Factory Method
│   │   ├── builder.go                # Step-by-step object construction
│   │   ├── observer.go               # Event subscription/notification
│   │   ├── strategy.go               # Interchangeable algorithms
│   │   ├── decorator.go              # Layered behavior addition
│   │   ├── adapter.go                # Incompatible interface bridging
│   │   ├── proxy.go                  # Access control, caching proxy
│   │   ├── facade.go                 # Simplified subsystem interface
│   │   ├── state.go                  # Vending machine state transitions
│   │   ├── composite.go              # Tree structures (filesystem)
│   │   ├── bridge.go                 # Abstraction ↔ Implementation decoupling
│   │   ├── iterator.go               # Sequential access (BST in-order)
│   │   ├── template_method.go        # Algorithm skeleton with varying steps
│   │   ├── chain_of_responsibility.go # Request pipeline (middleware)
│   │   ├── mediator.go               # Chat room communication
│   │   └── flyweight.go              # Shared object pool (text rendering)
│   ├── solid/
│   │   └── solid_principles.go       # SRP, OCP, LSP, ISP, DIP with code
│   └── problems/
│       ├── lru_cache.go              # LRU with doubly-linked list + hashmap
│       ├── rate_limiter.go           # Token Bucket (thread-safe)
│       ├── parking_lot.go            # Multi-level, multi-vehicle type
│       ├── elevator_system.go        # Multi-elevator scheduling
│       ├── pub_sub_system.go         # Topic-based publish/subscribe
│       ├── logging_framework.go      # Sinks, levels, singleton factory
│       ├── task_scheduler.go         # Priority queue + worker pool
│       └── online_shopping.go        # Cart, inventory, payment, orders
│
├── system-design-hld/
│   ├── fundamentals/
│   │   ├── consistent_hashing.go     # Hash ring, virtual nodes, replication
│   │   ├── load_balancer.go          # Round Robin, Weighted, Least Conn, IP Hash
│   │   ├── rate_limiter_algorithms.go # Token Bucket, Leaky Bucket, Fixed/Sliding Window
│   │   ├── bloom_filter.go           # Probabilistic membership testing
│   │   └── circuit_breaker.go        # CLOSED → OPEN → HALF_OPEN transitions
│   └── problems/
│       ├── url_shortener.go          # Hash-based shortening, TTL, analytics
│       └── distributed_cache.go      # Sharded LRU + TTL + concurrent access
│
├── computer-core/
│   ├── operating-systems/
│   │   └── concurrency_patterns.go   # Fan-Out/Fan-In, Worker Pool, Pipeline, Bounded Parallelism
│   ├── networking/
│   └── databases/
│
├── application-core/
│   ├── webapp/
│   └── embedded-systems/
│
├── frameworks-libraries/
│   ├── gin/
│   └── grpc/
│
└── go.mod
```

## Conventions

- Each file has `package main` and a `main()` function
- Tests use `fmt.Println` with assertions via `panic()`
- Go naming: snake_case for files, camelCase for functions
- Every file is self-contained and runnable: `go run <file>.go`
- Design patterns include real-world examples, not toy code

## Sources Adapted

Content curated and adapted from:
- [awesome-low-level-design](https://github.com/ashishps1/awesome-low-level-design) Design patterns, LLD problems
- [awesome-system-design-resources](https://github.com/ashishps1/awesome-system-design-resources) HLD fundamentals
- [coding-interview-patterns](https://github.com/ByteByteGoHq/coding-interview-patterns) DSA patterns

All code rewritten to match our single-file, self-contained, assertion-tested style.
