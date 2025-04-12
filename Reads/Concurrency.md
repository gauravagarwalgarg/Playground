Here’s a detailed guide covering **C++ concurrency interview questions**, along with **explanations and code examples**.

---

## **1. Basics of Concurrency and Multithreading**
### **Q1: What is a thread in C++?**
**Answer:**  
A thread is the smallest unit of execution within a process. Threads share memory space with other threads in the same process and execute independently.

### **Q2: What is the difference between a process and a thread?**
**Answer:**
| Feature  | Process  | Thread  |
|---|---|---|
| Memory Space  | Separate memory  | Shared memory with other threads in the process |
| Creation Overhead  | High (requires OS involvement)  | Low (lighter weight, part of the process) |
| Communication  | Requires IPC (pipes, message queues, shared memory) | Uses shared memory, mutexes, etc. |
| Isolation  | More isolated (crash doesn’t affect others)  | Less isolated (crash can affect all threads) |

### **Q3: How do you create and manage threads in C++?**
**Code Example:**
```cpp
#include <iostream>
#include <thread>

void hello() {
    std::cout << "Hello from thread!" << std::endl;
}

int main() {
    std::thread t(hello);  // Create a new thread
    t.join();  // Wait for the thread to complete
    return 0;
}
```
**Explanation:**  
- `std::thread` is used to create a new thread.
- `t.join()` ensures the main thread waits for `t` to complete.

---

## **2. Thread Management**
### **Q4: What is `std::thread::detach()` and `std::thread::join()`?**
**Answer:**  
- `join()`: Makes the calling thread wait until the other thread finishes execution.
- `detach()`: Detaches the thread from the calling thread so it runs independently.

**Code Example:**
```cpp
#include <iostream>
#include <thread>
#include <chrono>

void task() {
    std::this_thread::sleep_for(std::chrono::seconds(1));
    std::cout << "Thread executed independently.\n";
}

int main() {
    std::thread t(task);
    t.detach();  // Thread runs independently
    std::cout << "Main function ends.\n";
    std::this_thread::sleep_for(std::chrono::seconds(2));  // Ensure thread has time to complete
    return 0;
}
```

### **Q5: What happens if you don’t join or detach a thread before it goes out of scope?**
**Answer:**  
If a thread object is destroyed without being joined or detached, the program **terminates** with `std::terminate`.

**Code Example (Bad Practice):**
```cpp
void task() {
    std::cout << "Running thread...\n";
}

int main() {
    std::thread t(task);  // Thread is created
    // t goes out of scope without join or detach -> undefined behavior!
    return 0;
}
```

To fix this, always call `t.join()` or `t.detach()`.

---

## **3. Synchronization Mechanisms**
### **Q6: What are mutexes, and how do you use `std::mutex`?**
**Answer:**  
A mutex (short for "mutual exclusion") prevents multiple threads from accessing shared data simultaneously.

**Code Example:**
```cpp
#include <iostream>
#include <thread>
#include <mutex>

std::mutex mtx;  // Mutex declaration

void printMessage(const std::string& msg) {
    std::lock_guard<std::mutex> lock(mtx);  // Automatically locks and unlocks
    std::cout << msg << std::endl;
}

int main() {
    std::thread t1(printMessage, "Hello from Thread 1");
    std::thread t2(printMessage, "Hello from Thread 2");

    t1.join();
    t2.join();
    return 0;
}
```

---

### **Q7: What is a deadlock? How do you prevent it?**
**Answer:**  
A **deadlock** occurs when two or more threads are waiting for each other to release resources, causing a cycle.

**Prevention:**
- Always acquire locks in the same order.
- Use `std::lock()` for deadlock-free locking.

**Deadlock Example (BAD CODE):**
```cpp
#include <iostream>
#include <thread>
#include <mutex>

std::mutex mtx1, mtx2;

void task1() {
    std::lock_guard<std::mutex> lock1(mtx1);
    std::this_thread::sleep_for(std::chrono::milliseconds(100));
    std::lock_guard<std::mutex> lock2(mtx2); // Potential Deadlock
    std::cout << "Task 1 done\n";
}

void task2() {
    std::lock_guard<std::mutex> lock2(mtx2);
    std::this_thread::sleep_for(std::chrono::milliseconds(100));
    std::lock_guard<std::mutex> lock1(mtx1); // Potential Deadlock
    std::cout << "Task 2 done\n";
}

int main() {
    std::thread t1(task1);
    std::thread t2(task2);
    t1.join();
    t2.join();
    return 0;
}
```

**Fix (Using `std::lock()`):**
```cpp
void safe_task1() {
    std::lock(mtx1, mtx2);
    std::lock_guard<std::mutex> lock1(mtx1, std::adopt_lock);
    std::lock_guard<std::mutex> lock2(mtx2, std::adopt_lock);
    std::cout << "Safe Task 1 done\n";
}
```

---

## **4. Thread Communication and Coordination**
### **Q8: What is `std::condition_variable`? Provide an example.**
**Answer:**  
A **condition variable** is used for thread synchronization when one thread needs to wait for another.

**Code Example:**
```cpp
#include <iostream>
#include <thread>
#include <mutex>
#include <condition_variable>

std::mutex mtx;
std::condition_variable cv;
bool ready = false;

void worker() {
    std::unique_lock<std::mutex> lock(mtx);
    cv.wait(lock, [] { return ready; });  // Wait until ready is true
    std::cout << "Worker thread proceeding...\n";
}

int main() {
    std::thread t(worker);

    std::this_thread::sleep_for(std::chrono::seconds(1));
    {
        std::lock_guard<std::mutex> lock(mtx);
        ready = true;
    }
    cv.notify_one();  // Notify the worker

    t.join();
    return 0;
}
```

---

### **Q9: What is `std::future` and `std::promise`?**
**Answer:**  
They allow threads to communicate results asynchronously.

**Code Example:**
```cpp
#include <iostream>
#include <thread>
#include <future>

int computeSquare(int x) {
    return x * x;
}

int main() {
    std::future<int> result = std::async(std::launch::async, computeSquare, 5);
    std::cout << "Square: " << result.get() << std::endl; // Blocks until result is available
    return 0;
}
```

---

### **Q10: What are spinlocks, and when should they be used?**
**Answer:**  
A **spinlock** continuously checks a condition until it's available. It is used when waiting is expected to be short.

**Code Example:**
```cpp
#include <atomic>
#include <thread>
#include <iostream>

std::atomic_flag lock_flag = ATOMIC_FLAG_INIT;

void spinlock_task() {
    while (lock_flag.test_and_set(std::memory_order_acquire));
    std::cout << "Thread got lock\n";
    lock_flag.clear(std::memory_order_release);
}

int main() {
    std::thread t1(spinlock_task), t2(spinlock_task);
    t1.join();
    t2.join();
    return 0;
}
```

---

This covers **C++ concurrency interview questions** with **answers and example code**. Let me know if you need further details!
