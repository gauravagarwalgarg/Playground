# Python - Interview Preparation

## How to Run

### Individual files
```bash
python <file>.py
```

### Run all tests with pytest
```bash
pip install pytest
python -m pytest python/ -v
```

### Run a specific file
```bash
python python/algorithms-dsa/leetcode/arrays-hashing/two_sum.py
```

## Structure

```
python/
  algorithms-dsa/
    leetcode/
      arrays-hashing/    # Two Sum, Contains Duplicate, Valid Anagram
      trees/             # Invert Binary Tree
      graphs/            # Number of Islands
      dynamic-programming/ # Climbing Stairs
    patterns/
      two-pointers/      # Container With Most Water
      sliding-window/    # Best Time to Buy/Sell Stock
  low-level-design-lld/
    design-patterns/     # Singleton, Observer
    problems/            # Parking Lot
  computer-core/
    operating-systems/   # Producer-Consumer
    networking/          # TCP Echo Server
```

## Conventions

- Each file is self-contained with `if __name__ == "__main__"` block
- All tests use `assert` statements
- Functions are typed with Python type hints
