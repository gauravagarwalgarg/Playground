"""
Producer-Consumer Problem
Classic concurrency problem using threading and a bounded buffer.

Demonstrates: mutex, condition variables, bounded queue.
"""

import threading
import queue
import time


def producer_consumer_demo():
    """Run a producer-consumer simulation with assertions."""
    buffer = queue.Queue(maxsize=5)
    produced = []
    consumed = []
    num_items = 10

    def producer():
        for i in range(num_items):
            item = f"item-{i}"
            buffer.put(item)  # blocks if full
            produced.append(item)

    def consumer():
        for _ in range(num_items):
            item = buffer.get()  # blocks if empty
            consumed.append(item)
            buffer.task_done()

    prod_thread = threading.Thread(target=producer)
    cons_thread = threading.Thread(target=consumer)

    prod_thread.start()
    cons_thread.start()

    prod_thread.join()
    cons_thread.join()

    return produced, consumed


if __name__ == "__main__":
    produced, consumed = producer_consumer_demo()

    assert len(produced) == 10
    assert len(consumed) == 10
    assert set(produced) == set(consumed), "All produced items must be consumed"

    # Test with multiple producers and consumers
    buffer = queue.Queue(maxsize=3)
    results = []
    lock = threading.Lock()

    def multi_producer(pid, count):
        for i in range(count):
            buffer.put(f"p{pid}-{i}")

    def multi_consumer(count):
        for _ in range(count):
            item = buffer.get()
            with lock:
                results.append(item)
            buffer.task_done()

    threads = []
    # 3 producers, each producing 5 items
    for pid in range(3):
        threads.append(threading.Thread(target=multi_producer, args=(pid, 5)))
    # 3 consumers, each consuming 5 items
    for _ in range(3):
        threads.append(threading.Thread(target=multi_consumer, args=(5,)))

    for t in threads:
        t.start()
    for t in threads:
        t.join()

    assert len(results) == 15, f"Expected 15 items, got {len(results)}"

    print("All tests passed!")
