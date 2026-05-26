# Java - Interview Preparation

## How to Run

### Compile and run individual files
```bash
cd java/algorithms-dsa/leetcode/arrays-hashing
javac TwoSum.java
java -ea TwoSum
```

> Note: Use `-ea` flag to enable assertions.

### Compile all
```bash
find java/ -name "*.java" -exec javac {} \;
```

## Structure

```
java/
  algorithms-dsa/
    leetcode/
      arrays-hashing/    # TwoSum, ContainsDuplicate
      trees/             # InvertBinaryTree
  low-level-design-lld/
    design-patterns/     # Singleton
    problems/            # ParkingLot
```

## Conventions

- Each file has `public static void main(String[] args)` with test assertions
- Run with `-ea` to enable Java assertions
- Standard Java naming: PascalCase for classes, camelCase for methods
