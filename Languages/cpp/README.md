# C++ Interview Preparation

## Build

```bash
mkdir build && cd build
cmake ..
make -j$(nproc)
```

All executables land in `build/bin/`. Each `.cpp` file with a `main()` auto-compiles -- no CMakeLists edits needed.

## Structure

```
cpp/
├── algorithms-dsa/
│   ├── patterns/
│   │   ├── two-pointers/
│   │   ├── sliding-window/
│   │   ├── binary-search/
│   │   ├── bfs-dfs/
│   │   ├── backtracking/
│   │   ├── dynamic-programming/
│   │   └── greedy/
│   └── leetcode/
│       ├── arrays-hashing/
│       ├── two-pointers/
│       ├── sliding-window/
│       ├── stack/
│       ├── binary-search/
│       ├── linked-list/
│       ├── trees/
│       ├── graphs/
│       ├── dynamic-programming/
│       └── heap/
├── system-design-hld/
│   ├── fundamentals/
│   └── problems/
├── low-level-design-lld/
│   ├── design-patterns/
│   ├── solid/
│   └── problems/
├── computer-core/
│   ├── operating-systems/
│   ├── networking/
│   └── databases/
├── application-core/
│   ├── webapp/
│   └── embedded-systems/
└── CMakeLists.txt
```

## Adding a Problem

Create a `.cpp` file anywhere in the tree with a `main()` function and test cases:

```cpp
#include <iostream>
#include <vector>
#include <cassert>
using namespace std;

// Solution here...

int main() {
    // Test cases
    assert(solution({2,7,11,15}, 9) == vector<int>({0,1}));
    cout << "All tests passed!" << endl;
    return 0;
}
```

Rebuild: `cd build && make`
