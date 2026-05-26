"""
Singleton Design Pattern
Ensures a class has only one instance and provides a global point of access.

Thread-safe implementation using a metaclass.
"""

import threading


class SingletonMeta(type):
    """Thread-safe Singleton metaclass."""
    _instances = {}
    _lock = threading.Lock()

    def __call__(cls, *args, **kwargs):
        with cls._lock:
            if cls not in cls._instances:
                instance = super().__call__(*args, **kwargs)
                cls._instances[cls] = instance
        return cls._instances[cls]


class Database(metaclass=SingletonMeta):
    def __init__(self):
        self.connection = "PostgreSQL@localhost:5432"

    def query(self, sql: str) -> str:
        return f"Executing: {sql}"


if __name__ == "__main__":
    db1 = Database()
    db2 = Database()

    assert db1 is db2, "Both should be the same instance"
    assert id(db1) == id(db2)
    assert db1.connection == "PostgreSQL@localhost:5432"
    assert db1.query("SELECT 1") == "Executing: SELECT 1"

    # Thread safety test
    results = []

    def create_instance():
        results.append(id(Database()))

    threads = [threading.Thread(target=create_instance) for _ in range(10)]
    for t in threads:
        t.start()
    for t in threads:
        t.join()

    assert len(set(results)) == 1, "All threads should get the same instance"

    print("All tests passed!")
