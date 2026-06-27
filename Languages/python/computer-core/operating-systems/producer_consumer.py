"""Classic Producer/Consumer with queue.Queue and threading.Condition.

Demonstrates bounded buffer pattern where producers wait when buffer is full
and consumers wait when buffer is empty.
"""

import queue
import threading
import time
import random

def producer(buffer: queue.Queue, name: str, items: int) -> None:
    """Produce items and place them in the bounded buffer."""
    for i in range(items):
        item = f"{name}-item-{i}"
        buffer.put(item)  # Blocks if buffer is full
        print(f"[{name}] Produced: {item} (queue size: {buffer.qsize()})")
        time.sleep(random.uniform(0.1, 0.3))
    print(f"[{name}] Finished producing")

def consumer(buffer: queue.Queue, name: str, stop_event: threading.Event) -> None:
    """Consume items from the buffer until stop event is set."""
    while not stop_event.is_set() or not buffer.empty():
        try:
            item = buffer.get(timeout=0.5)
            print(f"  [{name}] Consumed: {item}")
            buffer.task_done()
            time.sleep(random.uniform(0.2, 0.5))
        except queue.Empty:
            continue
    print(f"  [{name}] Finished consuming")

if __name__ == "__main__":
    BUFFER_SIZE = 5
    buffer: queue.Queue = queue.Queue(maxsize=BUFFER_SIZE)
    stop_event = threading.Event()

    # Start producers and consumers
    producers = [
        threading.Thread(target=producer, args=(buffer, f"P{i}", 5))
        for i in range(2)
    ]
    consumers = [
        threading.Thread(target=consumer, args=(buffer, f"C{i}", stop_event))
        for i in range(3)
    ]

    print(f"Starting with buffer size={BUFFER_SIZE}")
    print(f"Producers: {len(producers)}, Consumers: {len(consumers)}\n")

    for c in consumers:
        c.start()
    for p in producers:
        p.start()

    # Wait for all producers to finish
    for p in producers:
        p.join()

    # Wait for buffer to drain, then signal consumers to stop
    buffer.join()
    stop_event.set()

    for c in consumers:
        c.join()

    print("\nAll done. Buffer is empty.")
