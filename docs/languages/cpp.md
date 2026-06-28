# C++ Language Guide

Complete C++ interview preparation covering DSA, system design, and computer science fundamentals.

## Directory Structure

All C++ code lives under `Languages/cpp/`. Build with CMake: `cmake -B build && cmake --build build`

## Categories

| Category | Topics | Path |
|----------|--------|------|
| **Algorithms & DSA** | 19 coding patterns (see below) + LeetCode | `Languages/cpp/algorithms-dsa/` |
| **Low-Level Design** | SOLID principles, Design Patterns (GoF), OOP problems | `Languages/cpp/low-level-design-lld/` |
| **System Design HLD** | Fundamentals (CAP, scaling, caching), Problems | `Languages/cpp/system-design-hld/` |
| **Computer Core** | Operating Systems, Networking, Databases | `Languages/cpp/computer-core/` |
| **Application Core** | Language features, STL, memory, concurrency | `Languages/cpp/application-core/` |
| **Frameworks & Libraries** | Boost, testing frameworks, build tools | `Languages/cpp/frameworks-libraries/` |

## Algorithms & DSA Patterns

| Pattern | Type | Path |
|---------|------|------|
| Two Pointers | folder | `algorithms-dsa/patterns/two-pointers/` |
| Sliding Window | folder | `algorithms-dsa/patterns/sliding-window/` |
| Binary Search | folder | `algorithms-dsa/patterns/binary-search/` |
| BFS / DFS | folder | `algorithms-dsa/patterns/bfs-dfs/` |
| Backtracking | folder | `algorithms-dsa/patterns/backtracking/` |
| Dynamic Programming | folder | `algorithms-dsa/patterns/dynamic-programming/` |
| Greedy | folder | `algorithms-dsa/patterns/greedy/` |
| Graphs | file | `algorithms-dsa/patterns/graphs.cpp` |
| Heaps | file | `algorithms-dsa/patterns/heaps.cpp` |
| Intervals | file | `algorithms-dsa/patterns/intervals.cpp` |
| Prefix Sums | file | `algorithms-dsa/patterns/prefix_sums.cpp` |
| Stacks | file | `algorithms-dsa/patterns/stacks.cpp` |
| Tries | file | `algorithms-dsa/patterns/tries.cpp` |
| Bit Manipulation | file | `algorithms-dsa/patterns/bit_manipulation.cpp` |

## Low-Level Design

| Sub-folder | Content |
|-----------|---------|
| `design-patterns/` | GoF patterns: Singleton, Factory, Observer, Strategy, etc. |
| `problems/` | LLD interview problems: parking lot, elevator, chess |
| `solid/` | SOLID principle implementations |
| `DesignPatterns.md` | Design patterns reference guide |
| `SOLID.md` / `SOLIDWithCode.md` | SOLID principles with C++ examples |

## Computer Core

| Sub-folder | Topics |
|-----------|--------|
| `operating-systems/` | Process mgmt, threading, memory, scheduling |
| `networking/` | TCP/UDP, sockets, HTTP, DNS |
| `databases/` | SQL concepts, indexing, transactions |
