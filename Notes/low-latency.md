## 🧵 1. **Multithreading & Concurrency**

### 🔹 `std::thread`
Create native threads directly.

```cpp
#include <thread>

void task() { /*...*/ }

std::thread t(task);
t.join(); // wait for thread to finish
```

### 🔹 `std::mutex`, `std::lock_guard`, `std::unique_lock`

```cpp
std::mutex mtx;

void safeIncrement(int& value) {
    std::lock_guard<std::mutex> lock(mtx); // auto releases lock
    value++;
}
```

- `lock_guard`: lightweight scoped lock
- `unique_lock`: flexible (can defer, release, etc.)

---

## 🕒 2. **Asynchronous Programming**

### 🔹 `std::async`, `std::future`, `std::promise`

```cpp
auto f = std::async(std::launch::async, []() { return 42; });
int result = f.get(); // blocks until result ready
```

- Launch policy: `std::launch::deferred`, `std::launch::async`
- `std::promise`: manually push value into a `future`

---

## 🏁 3. **Lock-Free & Low-Latency Programming**

### 🔹 `std::atomic`

```cpp
#include <atomic>

std::atomic<int> counter(0);

void fastIncrement() {
    counter.fetch_add(1, std::memory_order_relaxed);
}
```

- Safe for concurrent modification without mutex
- Fine-tuned memory models: `memory_order_relaxed`, `acquire`, `release`, `seq_cst`

### 🔹 `std::atomic_flag`

For ultra-lightweight spinlocks:

```cpp
std::atomic_flag lock = ATOMIC_FLAG_INIT;

void spinLock() {
    while (lock.test_and_set(std::memory_order_acquire));
    // critical section
    lock.clear(std::memory_order_release);
}
```

---

## 🧬 4. **Template Meta-programming**

### 🔹 `constexpr` (compile-time computation)

```cpp
constexpr int square(int x) { return x * x; }
int area = square(5); // evaluated at compile-time
```

### 🔹 `static_assert`

```cpp
static_assert(sizeof(void*) == 8, "64-bit only");
```

### 🔹 Variadic templates

```cpp
template<typename T>
void print(T t) { std::cout << t << "\n"; }

template<typename T, typename... Args>
void print(T t, Args... args) {
    std::cout << t << ", ";
    print(args...);
}
```

---

## 💡 5. **Performance & Latency-Oriented Features**

### 🔹 Move semantics (`&&`, `std::move`)

Avoid deep copies:

```cpp
std::vector<int> getVector() {
    std::vector<int> v = {1, 2, 3};
    return v; // move happens here
}
```

### 🔹 `emplace_back` instead of `push_back`

```cpp
vec.emplace_back(1, 2); // constructs in place
```

### 🔹 `alignas`, `alignof`

Control memory layout useful for SIMD/low-level optimization.

```cpp
alignas(64) struct AlignedData {
    float data[16];
};
```

---
