# C - Interview Preparation

## How to Build

### Build all
```bash
cd c
make
```

### Build and run tests
```bash
make test
```

### Clean
```bash
make clean
```

### Build individual file
```bash
gcc -Wall -Wextra -std=c11 -o bin/two_sum algorithms-dsa/leetcode/arrays-hashing/two_sum.c
```

## Structure

```
c/
  Makefile                          # Auto-discovers all .c files
  algorithms-dsa/
    leetcode/
      arrays-hashing/               # two_sum, contains_duplicate
  computer-core/
    operating-systems/              # producer_consumer (pthreads)
    networking/                     # tcp_echo_server (sockets)
  application-core/
    embedded-systems/               # gpio_abstraction
```

## Conventions

- Each file is self-contained with `main()` and `assert()` tests
- Compiled with `-Wall -Wextra -std=c11`
- Linked with `-lpthread` for threading programs
- Output binaries go to `bin/` (gitignored)
