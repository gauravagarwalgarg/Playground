/**
 * LLD Problem: Rate Limiter
 *
 * Implement a rate limiter using the Token Bucket algorithm with a
 * sliding window approach.
 *
 * Token Bucket:
 * - Bucket has max capacity of tokens
 * - Tokens are added at a fixed refill rate
 * - Each request consumes one token
 * - Request denied if no tokens available
 *
 * Also includes a Sliding Window rate limiter for comparison.
 */
#include <iostream>
#include <deque>
#include <cassert>
#include <algorithm>
using namespace std;

// --- Token Bucket Rate Limiter ---
class TokenBucketLimiter {
public:
    TokenBucketLimiter(int maxTokens, double refillRate)
        : maxTokens_(maxTokens), tokens_(maxTokens),
          refillRate_(refillRate), lastRefill_(0.0) {}

    bool allowRequest(double timestamp) {
        refill(timestamp);
        if (tokens_ > 0) {
            tokens_--;
            return true;
        }
        return false;
    }

private:
    int maxTokens_;
    double tokens_;
    double refillRate_;  // tokens per second
    double lastRefill_;

    void refill(double timestamp) {
        double elapsed = timestamp - lastRefill_;
        double newTokens = elapsed * refillRate_;
        tokens_ = min((double)maxTokens_, tokens_ + newTokens);
        lastRefill_ = timestamp;
    }
};

// --- Sliding Window Rate Limiter ---
class SlidingWindowLimiter {
public:
    SlidingWindowLimiter(int maxRequests, double windowSeconds)
        : maxRequests_(maxRequests), windowSeconds_(windowSeconds) {}

    bool allowRequest(double timestamp) {
        // Remove expired timestamps
        while (!window_.empty() && window_.front() <= timestamp - windowSeconds_) {
            window_.pop_front();
        }
        if ((int)window_.size() < maxRequests_) {
            window_.push_back(timestamp);
            return true;
        }
        return false;
    }

private:
    int maxRequests_;
    double windowSeconds_;
    deque<double> window_;
};

int main() {
    // Test Token Bucket: 3 tokens, refill 1 token/sec
    {
        TokenBucketLimiter limiter(3, 1.0);

        assert(limiter.allowRequest(0.0) == true);   // 2 tokens left
        assert(limiter.allowRequest(0.1) == true);   // 1 token left
        assert(limiter.allowRequest(0.2) == true);   // 0 tokens left
        assert(limiter.allowRequest(0.3) == false);  // denied
        assert(limiter.allowRequest(0.5) == false);  // still not enough

        // After 1 second, 1 token refilled
        assert(limiter.allowRequest(1.3) == true);   // ~1 token refilled
        assert(limiter.allowRequest(1.4) == false);  // used it

        cout << "Token Bucket tests passed!" << endl;
    }

    // Test Sliding Window: 3 requests per 1-second window
    {
        SlidingWindowLimiter limiter(3, 1.0);

        assert(limiter.allowRequest(0.0) == true);
        assert(limiter.allowRequest(0.3) == true);
        assert(limiter.allowRequest(0.6) == true);
        assert(limiter.allowRequest(0.9) == false);  // window full

        // After window slides past first request
        assert(limiter.allowRequest(1.1) == true);   // 0.0 expired
        assert(limiter.allowRequest(1.2) == false);  // 0.3, 0.6, 1.1 in window

        cout << "Sliding Window tests passed!" << endl;
    }

    cout << "All tests passed!" << endl;
    return 0;
}
