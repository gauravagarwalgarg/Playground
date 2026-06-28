# Rust - Interview Preparation

## How to Run

### Build and run
```bash
cd rust
cargo run
```

### Run tests
```bash
cd rust
cargo test
```

### Run tests with output
```bash
cargo test -- --nocapture
```

## Structure

```
rust/
  Cargo.toml
  src/
    main.rs          # All algorithms + #[test] annotations
```

## Algorithms Implemented

- Two Sum (HashMap)
- Contains Duplicate (HashSet)
- Best Time to Buy/Sell Stock
- Climbing Stairs (DP)
- Container With Most Water (Two Pointers)

## Conventions

- Functions use `&[T]` slices for input arrays
- Tests use `#[cfg(test)]` module with `assert_eq!`
- Each function is documented with `///` doc comments
