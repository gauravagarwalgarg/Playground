# Go - Interview Preparation

## How to Run

### Individual files
```bash
cd go/algorithms-dsa/leetcode/arrays_hashing
go run two_sum.go
```

### Run all
```bash
cd go
find . -name "*.go" -exec go run {} \;
```

## Structure

```
go/
  algorithms-dsa/
    leetcode/
      arrays_hashing/    # Two Sum, Contains Duplicate
      trees/             # Invert Binary Tree
    patterns/
      two_pointers/      # Container With Most Water
      sliding_window/    # Best Time to Buy/Sell Stock
  low-level-design-lld/
    design_patterns/     # Singleton
  computer-core/
    operating_systems/   # Producer-Consumer (goroutines + channels)
```

## Conventions

- Each file has `package main` and a `main()` function
- Tests use `fmt.Println` with manual assertions via `panic`
- Go naming: snake_case for files, camelCase for functions
