"""
Proxy Pattern

Provides a surrogate or placeholder for another object to control access.
CachingProxy wraps a slow DatabaseQuery with lazy loading and caching.
"""

from abc import ABC, abstractmethod
import time


class DataSource(ABC):
    @abstractmethod
    def query(self, sql: str) -> list[dict]:
        pass


class SlowDatabaseQuery(DataSource):
    """Simulates a slow database with artificial delay."""

    def __init__(self):
        self._data = {
            "SELECT * FROM users": [
                {"id": 1, "name": "Alice"},
                {"id": 2, "name": "Bob"},
            ],
            "SELECT * FROM orders": [
                {"id": 1, "total": 99.99},
            ],
        }

    def query(self, sql: str) -> list[dict]:
        time.sleep(0.1)  # simulate latency
        return self._data.get(sql, [])


class CachingProxy(DataSource):
    """Caches query results to avoid repeated slow lookups."""

    def __init__(self, source: DataSource | None = None):
        self._source = source  # lazy: may be set later
        self._cache: dict[str, list[dict]] = {}
        self._hits = 0

    @property
    def source(self) -> DataSource:
        if self._source is None:
            self._source = SlowDatabaseQuery()  # lazy loading
        return self._source

    def query(self, sql: str) -> list[dict]:
        if sql in self._cache:
            self._hits += 1
            return self._cache[sql]
        result = self.source.query(sql)
        self._cache[sql] = result
        return result

    def invalidate(self, sql: str | None = None) -> None:
        if sql:
            self._cache.pop(sql, None)
        else:
            self._cache.clear()


if __name__ == "__main__":
    proxy = CachingProxy()  # lazy: no DB created yet
    assert proxy._source is None

    # First query: triggers lazy loading + actual query
    start = time.time()
    result = proxy.query("SELECT * FROM users")
    first_duration = time.time() - start
    assert len(result) == 2
    assert proxy._source is not None

    # Second query: served from cache (fast)
    start = time.time()
    cached = proxy.query("SELECT * FROM users")
    cache_duration = time.time() - start
    assert cached == result
    assert cache_duration < first_duration
    assert proxy._hits == 1

    # Invalidate and re-query
    proxy.invalidate("SELECT * FROM users")
    proxy.query("SELECT * FROM users")
    assert proxy._hits == 1  # cache miss after invalidation

    print("All tests passed!")
