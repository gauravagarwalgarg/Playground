# Interview Preparation Playground

Structured, compilable, interview-ready code across 7 languages and 19+ coding patterns.

## Module Map

| Module | Description | Languages | Path |
|--------|-------------|-----------|------|
| **Algorithms & DSA** | 19 coding patterns + LeetCode problems | C, C++, Go, Java, JS, Python, Rust | `Languages/*/algorithms-dsa/` |
| **Low-Level Design** | SOLID, Design Patterns, OOP Problems | C++, Go, Python, Java | `Languages/*/low-level-design-lld/` |
| **System Design HLD** | Fundamentals + System Design Problems | C++, Go, Python | `Languages/*/system-design-hld/` |
| **Computer Core** | OS, Networking, Databases | C++, Go, Python | `Languages/*/computer-core/` |
| **Application Core** | Language-specific app dev topics | All 7 | `Languages/*/application-core/` |
| **Frameworks & Libraries** | Language ecosystem tools | All 7 | `Languages/*/frameworks-libraries/` |
| **Software Engineering** | Architecture, Design, Product, PM | | `SoftwareEngineering/` |

## Language Coverage

| Language | DSA Patterns | LLD | HLD | Computer Core | Build System |
|----------|-------------|-----|-----|---------------|--------------|
| C++ | 7 folders + 7 files | design-patterns, problems, solid | fundamentals, problems | OS, networking, databases | CMake |
| Python | 7 folders + 13 files | design-patterns, problems, solid | fundamentals, problems | OS, networking, databases | |
| Go | 9 folders + 15 files | design-patterns, problems, solid | fundamentals, problems | OS, networking, databases | go.mod |
| Java | ✓ | ✓ | ✓ | ✓ | |
| JavaScript | ✓ | ✓ | ✓ | ✓ | package.json |
| C | ✓ | ✓ | ✓ | ✓ | Makefile |
| Rust | ✓ | ✓ | ✓ | ✓ | Cargo.toml |

## Quick Start

```bash
# Clone and explore
git clone https://github.com/GauravAgarwalGarg/Playground.git
cd Playground/Languages/cpp/algorithms-dsa/patterns/

# Build C++ examples
cd Languages/cpp && cmake -B build && cmake --build build

# Run Go patterns
cd Languages/go && go run algorithms-dsa/patterns/two_pointers.go
```

## Documentation

- [**Coding Patterns**](coding-patterns.md) All 19 patterns with problem counts and language availability
- [**Pattern Index**](pattern-index.md) Quick reference for pattern selection
- [**Study Roadmap**](study-roadmap.md) Suggested learning order and timeline
- **Language Guides:** [C++](languages/cpp.md) | [Python](languages/python.md) | [Go](languages/go.md) | [Java](languages/java.md) | [JavaScript](languages/javascript.md)
