# JavaScript - Interview Preparation

## How to Run

### Run tests (Node.js 18+)
```bash
cd javascript
node --test
```

### Run individual file
```bash
node --test algorithms-dsa/leetcode/arrays-hashing/two_sum.mjs
```

### Verbose output
```bash
npm test
# or
node --test --test-reporter=spec
```

## Structure

```
javascript/
  algorithms-dsa/
    leetcode/
      arrays-hashing/    # two_sum, contains_duplicate
    patterns/
      sliding-window/    # best_time_buy_sell
  low-level-design-lld/
    design-patterns/     # singleton, observer
```

## Conventions

- ES Modules (`.mjs` extension, `"type": "module"` in package.json)
- Tests use `node:test` built-in test runner (Node 18+)
- Assertions use `node:assert/strict`
- Functions are exported for potential reuse
