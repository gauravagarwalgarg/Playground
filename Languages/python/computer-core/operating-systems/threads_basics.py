"""Threading basics: create threads, join, daemon threads, thread-local data, and mutex/lock.

Demonstrates race conditions and how locks fix them.
"""

import threading
import time

# --- Basic Thread Creation ---

def worker(name: str, delay: float) -> None:
    """Simple worker that simulates work with a sleep."""
    print(f"[{name}] Starting")
    time.sleep(delay)
    print(f"[{name}] Done")

# --- Race Condition Demo ---

counter: int = 0

def increment_unsafe(n: int) -> None:
    """Increment counter without lock causes race condition."""
    global counter
    for _ in range(n):
        temp = counter
        temp += 1
        counter = temp

def increment_safe(n: int, lock: threading.Lock) -> None:
    """Increment counter with lock thread-safe."""
    global counter
    for _ in range(n):
        with lock:
            counter += 1

# --- Thread-Local Data ---

local_data = threading.local()

def thread_local_demo(value: str) -> None:
    """Each thread gets its own copy of local_data."""
    local_data.name = value
    time.sleep(0.1)
    print(f"[Thread] local_data.name = {local_data.name}")

if __name__ == "__main__":
    # Basic threads with join
    print("=== Basic Threads ===")
    threads = [threading.Thread(target=worker, args=(f"T{i}", 0.2)) for i in range(3)]
    for t in threads:
        t.start()
    for t in threads:
        t.join()

    # Daemon thread (won't block program exit)
    print("\n=== Daemon Thread ===")
    daemon = threading.Thread(target=worker, args=("Daemon", 5.0), daemon=True)
    daemon.start()
    print("Main continues without waiting for daemon")

    # Race condition
    print("\n=== Race Condition (unsafe) ===")
    counter = 0
    t1 = threading.Thread(target=increment_unsafe, args=(100000,))
    t2 = threading.Thread(target=increment_unsafe, args=(100000,))
    t1.start(); t2.start()
    t1.join(); t2.join()
    print(f"Expected: 200000, Got: {counter} (likely wrong)")

    # Fixed with lock
    print("\n=== With Lock (safe) ===")
    counter = 0
    lock = threading.Lock()
    t1 = threading.Thread(target=increment_safe, args=(100000, lock))
    t2 = threading.Thread(target=increment_safe, args=(100000, lock))
    t1.start(); t2.start()
    t1.join(); t2.join()
    print(f"Expected: 200000, Got: {counter}")

    # Thread-local data
    print("\n=== Thread-Local Data ===")
    threads = [threading.Thread(target=thread_local_demo, args=(f"value-{i}",)) for i in range(3)]
    for t in threads:
        t.start()
    for t in threads:
        t.join()
