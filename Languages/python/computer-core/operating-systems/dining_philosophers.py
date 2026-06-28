"""Dining Philosophers Problem deadlock-free solution using resource ordering.

Five philosophers sit at a table. Each needs two forks to eat.
Deadlock-free solution: always pick up the lower-numbered fork first.
This breaks the circular wait condition.
"""

import threading
import time
import random

NUM_PHILOSOPHERS = 5
NUM_MEALS = 3

def philosopher(pid: int, forks: list[threading.Lock]) -> None:
    """Philosopher alternates between thinking and eating.
    
    Uses resource ordering: always acquire lower-numbered fork first
    to prevent deadlock (breaks circular wait).
    """
    left = pid
    right = (pid + 1) % NUM_PHILOSOPHERS

    # Resource ordering: always pick up lower-numbered fork first
    first = min(left, right)
    second = max(left, right)

    for meal in range(NUM_MEALS):
        # Think
        print(f"  Philosopher {pid}: thinking...")
        time.sleep(random.uniform(0.1, 0.3))

        # Pick up forks in order
        forks[first].acquire()
        forks[second].acquire()

        # Eat
        print(f"🍝 Philosopher {pid}: eating meal {meal + 1}/{NUM_MEALS}")
        time.sleep(random.uniform(0.1, 0.2))

        # Put down forks
        forks[second].release()
        forks[first].release()

    print(f"✓ Philosopher {pid}: done eating all meals")

if __name__ == "__main__":
    print(f"Dining Philosophers (n={NUM_PHILOSOPHERS}, meals={NUM_MEALS})")
    print("Solution: Resource ordering (always pick lower-numbered fork first)\n")

    forks = [threading.Lock() for _ in range(NUM_PHILOSOPHERS)]
    philosophers = [
        threading.Thread(target=philosopher, args=(i, forks))
        for i in range(NUM_PHILOSOPHERS)
    ]

    for p in philosophers:
        p.start()
    for p in philosophers:
        p.join()

    print("\nAll philosophers finished without deadlock!")
