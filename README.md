# Interview Preparation Playground

[![CI](https://github.com/gauravagarwalgarg/playground/actions/workflows/ci.yml/badge.svg)](https://github.com/gauravagarwalgarg/playground/actions/workflows/ci.yml) [![Docs](https://img.shields.io/badge/docs-live-blue?logo=github)](https://gauravagarwalgarg.github.io/playground/) ![Multi-Language](https://img.shields.io/badge/C++-Go-Python-Java-blue) [![License](https://img.shields.io/github/license/gauravagarwalgarg/playground)](https://github.com/gauravagarwalgarg/playground/blob/develop/LICENSE)

> 📖 **Documentation**: [https://gauravagarwalgarg.github.io/playground/](https://gauravagarwalgarg.github.io/playground/)
>
> 📦 **Repository**: [GitHub](https://github.com/gauravagarwalgarg/playground)


> Structured, compilable, interview-ready code across multiple languages. DSA, System Design, LLD, OS, Networking, Databases, Design Patterns, and more.

---

## Repository Structure

```
Playground/
├── Languages/                 # All programming languages
│   ├── cpp/                   # C++20, CMake
│   ├── c/                     # C11, Makefile
│   ├── python/                # Python 3, pytest
│   ├── go/                    # Go modules
│   ├── java/                  # Java, javac
│   ├── javascript/            # ES Modules, node:test
│   └── rust/                  # Cargo
│
├── SoftwareEngineering/       # Architecture, Design, Process
├── Linux/                     # Kernel, scripts, cloud ops
├── Notes/                     # Conferences, tech talks, courses
└── Reads/                     # Book notes and reading list
```

---

## Languages

Each language follows the same internal structure:

```
<language>/
├── algorithms-dsa/
│   ├── patterns/              # Two Pointers, Sliding Window, BFS/DFS, DP, etc.
│   └── leetcode/              # By topic: arrays, trees, graphs, dp, heap, etc.
├── system-design-hld/
│   ├── fundamentals/          # Load balancers, caching, CAP, estimation
│   └── problems/              # URL Shortener, YouTube, Twitter, etc.
├── low-level-design-lld/
│   ├── design-patterns/       # GoF: Singleton, Observer, Factory, Strategy...
│   ├── solid/                 # SOLID principles with code
│   └── problems/              # Parking Lot, Elevator, LRU Cache, Order Book
├── computer-core/
│   ├── operating-systems/     # Threads, mutexes, scheduling, signals
│   ├── networking/            # TCP/UDP, sockets, HTTP
│   └── databases/             # SQL, indexing, ACID, sharding
├── application-core/
│   ├── webapp/                # REST APIs, web servers
│   └── embedded-systems/      # GPIO, RTOS, I2C/SPI
└── frameworks-libraries/      # Language-specific frameworks
```

| Language | Build | Run | Frameworks |
|----------|-------|-----|-----------|
| **C++** | `cmake .. && make` | `./bin/<name>` | - |
| **C** | `make` | `./bin/<name>` | - |
| **Python** | - | `python <file>.py` | Django, Flask, FastAPI, PyTorch |
| **Go** | - | `go run <file>.go` | Gin, gRPC |
| **Java** | `javac` | `java -ea <Class>` | Spring, Gradle |
| **JavaScript** | - | `node --test` | React, Angular, Express, Next.js |
| **Rust** | `cargo build` | `cargo test` | Actix, Tokio |

---

## Linux

```
Linux/
├── kernel/                    # Subsystems, modules, device drivers
├── userspace/                 # systemd, cgroups, namespaces
├── scripts/
│   ├── cloud-ops/             # Docker cleanup, K8s status, AWS EC2 list
│   ├── networking/            # Port scan, connectivity checks
│   ├── monitoring/            # System health, disk/CPU/memory
│   └── containers/            # Docker compose helpers
└── IRQ.md                     # Interrupt handling notes
```

---

## SoftwareEngineering

```
SoftwareEngineering/
├── architecture/
│   ├── patterns/              # Microservices, Event-Driven, CQRS, Hexagonal
│   └── distributed-systems/   # CAP, Consensus, Replication, Sharding
├── design/
│   ├── product-design/        # User stories, wireframing, MVP
│   └── api-design/            # REST, GraphQL, gRPC, versioning
├── product-engineering/       # Agile, Scrum, CI/CD, DevOps
└── project-management/        # Estimation, planning, retrospectives
```

---

## Notes & Reads

- **Notes/** -- Conference notes, tech talks, courses, deep dives
- **Reads/** -- Curated reading list: DDIA, Clean Code, SICP, OSTEP, and more

---

## Quick Start

```bash
# C++
cd Languages/cpp && mkdir build && cd build && cmake .. && make && ./bin/algorithms-dsa_leetcode_arrays-hashing_two_sum

# Python
python Languages/python/algorithms-dsa/leetcode/arrays-hashing/two_sum.py

# Go
go run Languages/go/algorithms-dsa/leetcode/arrays_hashing/two_sum.go

# C
cd Languages/c && make && ./bin/algorithms-dsa/leetcode/arrays-hashing/two_sum

# JavaScript
node Languages/javascript/algorithms-dsa/leetcode/arrays-hashing/two_sum.mjs

# Rust
cd Languages/rust && cargo test
```

---

## Adding a New Problem

1. Pick the language: `Languages/<lang>/`
2. Pick the category: `algorithms-dsa/leetcode/<topic>/`
3. Write the solution with `main()` + assertions
4. Build and run -- no config edits needed (CMake/Make auto-discover)

Same algorithm, same structure, every language.
