/**
 * Rate Limiting Algorithms
 *
 * Rate limiters control the rate of requests a client can make.
 * Two common algorithms implemented here:
 *
 * 1. Token Bucket: Allows bursts up to bucket capacity, then rate-limits.
 *    Tokens refill at a steady rate. Thread-safe via synchronized.
 *
 * 2. Fixed Window Counter: Divides time into fixed windows (e.g., 1 second).
 *    Counts requests per window. Resets at window boundary.
 *
 * Run: javac RateLimiter.java && java -ea RateLimiter
 */
public class RateLimiter {

    // ──────────────────────────────────────────────────────────────
    // Interface
    // ──────────────────────────────────────────────────────────────

    interface Limiter {
        /**
         * @return true if the request is allowed, false if rate-limited.
         */
        boolean allow();
    }

    // ──────────────────────────────────────────────────────────────
    // Token Bucket
    // ──────────────────────────────────────────────────────────────

    /**
     * Token Bucket Algorithm
     *
     * - Bucket holds up to `capacity` tokens.
     * - Tokens are added at `refillRate` tokens per second.
     * - Each allow() call consumes one token.
     * - Allows bursts up to capacity, then limits to refill rate.
     * - Thread-safe: all state access is synchronized.
     */
    static class TokenBucket implements Limiter {
        private final int capacity;
        private final double refillRate; // tokens per second
        private double tokens;
        private long lastRefillTimestamp;

        /**
         * @param capacity   Maximum tokens in the bucket
         * @param refillRate Tokens added per second
         */
        public TokenBucket(int capacity, double refillRate) {
            this.capacity = capacity;
            this.refillRate = refillRate;
            this.tokens = capacity; // Start full
            this.lastRefillTimestamp = System.nanoTime();
        }

        @Override
        public synchronized boolean allow() {
            refill();
            if (tokens >= 1.0) {
                tokens -= 1.0;
                return true;
            }
            return false;
        }

        private void refill() {
            long now = System.nanoTime();
            double elapsed = (now - lastRefillTimestamp) / 1_000_000_000.0;
            tokens = Math.min(capacity, tokens + elapsed * refillRate);
            lastRefillTimestamp = now;
        }

        // For testing: allows injecting time
        synchronized void setTokens(double t) {
            this.tokens = Math.min(capacity, t);
        }

        synchronized double getTokens() {
            refill();
            return tokens;
        }
    }

    // ──────────────────────────────────────────────────────────────
    // Fixed Window Counter
    // ──────────────────────────────────────────────────────────────

    /**
     * Fixed Window Counter Algorithm
     *
     * - Time is divided into fixed-size windows (e.g., 1 second).
     * - Each window allows up to `maxRequests` requests.
     * - Counter resets when a new window starts.
     * - Simple but can allow 2x burst at window boundaries.
     */
    static class FixedWindowCounter implements Limiter {
        private final int maxRequests;
        private final long windowSizeMs;
        private long windowStart;
        private int counter;

        /**
         * @param maxRequests  Maximum requests allowed per window
         * @param windowSizeMs Window duration in milliseconds
         */
        public FixedWindowCounter(int maxRequests, long windowSizeMs) {
            this.maxRequests = maxRequests;
            this.windowSizeMs = windowSizeMs;
            this.windowStart = System.currentTimeMillis();
            this.counter = 0;
        }

        @Override
        public synchronized boolean allow() {
            long now = System.currentTimeMillis();
            // Check if we've moved to a new window
            if (now - windowStart >= windowSizeMs) {
                windowStart = now;
                counter = 0;
            }
            if (counter < maxRequests) {
                counter++;
                return true;
            }
            return false;
        }

        // For testing
        synchronized int getCounter() {
            return counter;
        }
    }

    // ──────────────────────────────────────────────────────────────
    // Tests
    // ──────────────────────────────────────────────────────────────

    public static void main(String[] args) throws InterruptedException {
        testTokenBucketBurst();
        testTokenBucketRefill();
        testTokenBucketThreadSafety();
        testFixedWindowBasic();
        testFixedWindowReset();
        System.out.println("\n✓ All rate limiter tests passed!");
    }

    /**
     * Token bucket allows burst up to capacity, then blocks.
     */
    private static void testTokenBucketBurst() {
        System.out.println("=== Test: Token Bucket - Burst Behavior ===");
        TokenBucket tb = new TokenBucket(5, 1.0); // capacity=5, refill=1/sec

        // Should allow first 5 (burst)
        int allowed = 0;
        for (int i = 0; i < 10; i++) {
            if (tb.allow()) allowed++;
        }

        System.out.println("  Sent 10 requests, allowed: " + allowed);
        assert allowed == 5 : "Expected 5 allowed in burst, got: " + allowed;
        System.out.println("  ✓ Burst of 5 allowed, remaining 5 rejected\n");
    }

    /**
     * After waiting, tokens refill and more requests are allowed.
     */
    private static void testTokenBucketRefill() throws InterruptedException {
        System.out.println("=== Test: Token Bucket - Refill ===");
        TokenBucket tb = new TokenBucket(5, 10.0); // capacity=5, refill=10/sec

        // Drain all tokens
        for (int i = 0; i < 5; i++) {
            assert tb.allow() : "Should allow initial burst";
        }
        assert !tb.allow() : "Should be empty after burst";

        // Wait 300ms → should refill ~3 tokens (10/sec * 0.3s = 3)
        Thread.sleep(300);

        int allowed = 0;
        for (int i = 0; i < 10; i++) {
            if (tb.allow()) allowed++;
        }

        System.out.println("  After 300ms wait (refill rate=10/s): " + allowed + " allowed");
        // Allow some tolerance for timing: expect 2-4 tokens
        assert allowed >= 2 && allowed <= 4 : "Expected 2-4 tokens after 300ms, got: " + allowed;
        System.out.println("  ✓ Tokens refilled correctly\n");
    }

    /**
     * Multiple threads competing for tokens shouldn't exceed capacity.
     */
    private static void testTokenBucketThreadSafety() throws InterruptedException {
        System.out.println("=== Test: Token Bucket - Thread Safety ===");
        TokenBucket tb = new TokenBucket(100, 0); // No refill, just 100 tokens

        int numThreads = 10;
        int requestsPerThread = 20; // Total 200 requests for 100 tokens
        int[] results = new int[numThreads];
        Thread[] threads = new Thread[numThreads];

        for (int t = 0; t < numThreads; t++) {
            final int idx = t;
            threads[t] = new Thread(() -> {
                for (int i = 0; i < requestsPerThread; i++) {
                    if (tb.allow()) {
                        results[idx]++;
                    }
                }
            });
        }

        for (Thread thread : threads) thread.start();
        for (Thread thread : threads) thread.join();

        int totalAllowed = 0;
        for (int r : results) totalAllowed += r;

        System.out.println("  " + numThreads + " threads × " + requestsPerThread + " requests = " +
                (numThreads * requestsPerThread) + " total");
        System.out.println("  Total allowed: " + totalAllowed + " (capacity: 100)");
        assert totalAllowed == 100 : "Thread safety violated! Got: " + totalAllowed;
        System.out.println("  ✓ Exactly 100 requests allowed across all threads\n");
    }

    /**
     * Fixed window allows up to max, then blocks within the same window.
     */
    private static void testFixedWindowBasic() {
        System.out.println("=== Test: Fixed Window - Basic Limiting ===");
        FixedWindowCounter fw = new FixedWindowCounter(5, 1000); // 5 req per 1 second

        int allowed = 0;
        int rejected = 0;
        for (int i = 0; i < 10; i++) {
            if (fw.allow()) allowed++;
            else rejected++;
        }

        System.out.println("  Sent 10 requests in same window: " + allowed + " allowed, " + rejected + " rejected");
        assert allowed == 5 : "Expected 5 allowed, got: " + allowed;
        assert rejected == 5 : "Expected 5 rejected, got: " + rejected;
        System.out.println("  ✓ Fixed window correctly limits to max requests\n");
    }

    /**
     * Counter resets when a new window starts.
     */
    private static void testFixedWindowReset() throws InterruptedException {
        System.out.println("=== Test: Fixed Window - Window Reset ===");
        FixedWindowCounter fw = new FixedWindowCounter(3, 200); // 3 req per 200ms

        // Use up the window
        for (int i = 0; i < 3; i++) {
            assert fw.allow() : "Should allow within limit";
        }
        assert !fw.allow() : "Should reject at limit";
        System.out.println("  Window 1: 3 allowed, then rejected");

        // Wait for new window
        Thread.sleep(250);

        // New window should allow again
        int allowed = 0;
        for (int i = 0; i < 5; i++) {
            if (fw.allow()) allowed++;
        }

        System.out.println("  Window 2 (after 250ms): " + allowed + " allowed");
        assert allowed == 3 : "Expected 3 in new window, got: " + allowed;
        System.out.println("  ✓ Counter resets in new window\n");
    }
}
