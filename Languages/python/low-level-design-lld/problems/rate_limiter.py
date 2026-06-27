"""
Low-Level Design: Rate Limiter

Implement rate limiting using two strategies:
1. Token Bucket - tokens refill at fixed rate, each request consumes one
2. Sliding Window - track timestamps, reject if window is full

Both provide allowRequest(timestamp) -> bool interface.

Time: O(1) amortized for token bucket, O(1) amortized for sliding window
Space: O(1) for token bucket, O(max_requests) for sliding window
"""

from collections import deque


class TokenBucketLimiter:
    """Rate limiter using token bucket algorithm."""

    def __init__(self, max_tokens: int, refill_rate: float):
        """
        Args:
            max_tokens: Maximum tokens (burst capacity)
            refill_rate: Tokens added per second
        """
        self.max_tokens = max_tokens
        self.tokens = float(max_tokens)
        self.refill_rate = refill_rate
        self.last_refill = 0.0

    def allow_request(self, timestamp: float) -> bool:
        self._refill(timestamp)
        if self.tokens >= 1.0:
            self.tokens -= 1.0
            return True
        return False

    def _refill(self, timestamp: float) -> None:
        elapsed = timestamp - self.last_refill
        self.tokens = min(self.max_tokens, self.tokens + elapsed * self.refill_rate)
        self.last_refill = timestamp


class SlidingWindowLimiter:
    """Rate limiter using sliding window log."""

    def __init__(self, max_requests: int, window_seconds: float):
        self.max_requests = max_requests
        self.window_seconds = window_seconds
        self.timestamps: deque[float] = deque()

    def allow_request(self, timestamp: float) -> bool:
        # Evict expired entries
        while self.timestamps and self.timestamps[0] <= timestamp - self.window_seconds:
            self.timestamps.popleft()

        if len(self.timestamps) < self.max_requests:
            self.timestamps.append(timestamp)
            return True
        return False


if __name__ == "__main__":
    # Token Bucket: 3 tokens, refill 1/sec
    tb = TokenBucketLimiter(max_tokens=3, refill_rate=1.0)
    assert tb.allow_request(0.0) is True
    assert tb.allow_request(0.1) is True
    assert tb.allow_request(0.2) is True
    assert tb.allow_request(0.3) is False  # exhausted
    assert tb.allow_request(1.3) is True   # 1 token refilled
    assert tb.allow_request(1.4) is False  # used it
    print("Token Bucket tests passed!")

    # Sliding Window: 3 requests per 1 second
    sw = SlidingWindowLimiter(max_requests=3, window_seconds=1.0)
    assert sw.allow_request(0.0) is True
    assert sw.allow_request(0.3) is True
    assert sw.allow_request(0.6) is True
    assert sw.allow_request(0.9) is False  # window full
    assert sw.allow_request(1.1) is True   # 0.0 expired
    assert sw.allow_request(1.2) is False  # 0.3, 0.6, 1.1 still in window
    print("Sliding Window tests passed!")

    print("All tests passed!")
