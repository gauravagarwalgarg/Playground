/**
 * Dynamic Programming Templates
 * 
 * Four foundational DP patterns that cover most DP problems.
 * Each template shows the recurrence, base case, and iteration order.
 * 
 * Compile: g++ -std=c++17 -o dp_templates dp_templates.cpp
 */

#include <iostream>
#include <vector>
#include <string>
#include <algorithm>

using namespace std;

// =============================================================================
// TEMPLATE 1: Fibonacci-Style (1D Linear DP)
// Pattern: dp[i] depends on dp[i-1], dp[i-2], ... (constant lookback)
// Example: Climbing Stairs count ways to reach step n (take 1 or 2 steps)
// =============================================================================
int climbStairs(int n) {
    if (n <= 2) return n;
    
    // Space-optimized: only need previous two values
    int prev2 = 1;  // dp[i-2]: ways to reach step 1
    int prev1 = 2;  // dp[i-1]: ways to reach step 2
    
    for (int i = 3; i <= n; i++) {
        int curr = prev1 + prev2;  // dp[i] = dp[i-1] + dp[i-2]
        prev2 = prev1;
        prev1 = curr;
    }
    return prev1;
}

// Full array version (when you need all intermediate values):
// vector<int> dp(n+1);
// dp[0] = 1; dp[1] = 1;
// for (int i = 2; i <= n; i++) dp[i] = dp[i-1] + dp[i-2];

// =============================================================================
// TEMPLATE 2: 0/1 Knapsack
// Pattern: For each item, choose to TAKE or SKIP. Capacity constraint.
// Example: Given weights[] and values[], maximize value within capacity W
// =============================================================================
int knapsack01(const vector<int>& weights, const vector<int>& values, int W) {
    int n = weights.size();
    
    // dp[w] = max value achievable with capacity w
    // Space-optimized from 2D dp[i][w] to 1D dp[w]
    vector<int> dp(W + 1, 0);
    
    for (int i = 0; i < n; i++) {
        // CRITICAL: iterate capacity in REVERSE to avoid using item i twice
        // (Forward iteration = unbounded knapsack)
        for (int w = W; w >= weights[i]; w--) {
            dp[w] = max(dp[w],                          // skip item i
                        dp[w - weights[i]] + values[i]); // take item i
        }
    }
    return dp[W];
}

// 2D version (clearer logic):
// dp[i][w] = max value using items [0..i-1] with capacity w
// dp[i][w] = max(dp[i-1][w], dp[i-1][w-wt[i-1]] + val[i-1])

// =============================================================================
// TEMPLATE 3: Longest Common Subsequence (2D Two-Sequence DP)
// Pattern: dp[i][j] considers first i chars of s1 and first j chars of s2
// Example: Find length of longest subsequence common to both strings
// =============================================================================
int longestCommonSubsequence(const string& s1, const string& s2) {
    int m = s1.size(), n = s2.size();
    
    // dp[i][j] = LCS length of s1[0..i-1] and s2[0..j-1]
    vector<vector<int>> dp(m + 1, vector<int>(n + 1, 0));
    
    for (int i = 1; i <= m; i++) {
        for (int j = 1; j <= n; j++) {
            if (s1[i-1] == s2[j-1]) {
                dp[i][j] = dp[i-1][j-1] + 1;  // chars match: extend LCS
            } else {
                dp[i][j] = max(dp[i-1][j],     // skip char from s1
                               dp[i][j-1]);     // skip char from s2
            }
        }
    }
    return dp[m][n];
}

// Space-optimized to O(min(m,n)) using two rows:
// vector<int> prev(n+1, 0), curr(n+1, 0);
// swap prev and curr each iteration

// =============================================================================
// TEMPLATE 4: Longest Increasing Subsequence (LIS)
// Pattern: dp[i] = length of LIS ending at index i
// Example: Find length of longest strictly increasing subsequence
// =============================================================================

// O(n²) DP approach clearer, good for reconstructing the sequence
int lis_dp(const vector<int>& nums) {
    int n = nums.size();
    if (n == 0) return 0;
    
    // dp[i] = length of LIS ending at index i
    vector<int> dp(n, 1);  // every element is a subsequence of length 1
    
    for (int i = 1; i < n; i++) {
        for (int j = 0; j < i; j++) {
            if (nums[j] < nums[i]) {
                dp[i] = max(dp[i], dp[j] + 1);  // extend LIS ending at j
            }
        }
    }
    return *max_element(dp.begin(), dp.end());
}

// O(n log n) patience sorting approach optimal for just the length
int lis_binary_search(const vector<int>& nums) {
    // tails[i] = smallest tail element for increasing subsequence of length i+1
    vector<int> tails;
    
    for (int x : nums) {
        // Find first tail >= x (for strict increasing; use > for non-decreasing)
        auto it = lower_bound(tails.begin(), tails.end(), x);
        
        if (it == tails.end()) {
            tails.push_back(x);  // x extends the longest subsequence
        } else {
            *it = x;  // x can create a "better" subsequence of this length
        }
    }
    return tails.size();
}

// =============================================================================
int main() {
    // Fibonacci-style: Climbing Stairs
    cout << "Climb 5 stairs: " << climbStairs(5) << " ways" << endl;
    
    // 0/1 Knapsack
    vector<int> weights = {2, 3, 4, 5};
    vector<int> values = {3, 4, 5, 6};
    cout << "Knapsack(capacity=8): " << knapsack01(weights, values, 8) << endl;
    
    // LCS
    cout << "LCS(\"abcde\", \"ace\"): " << longestCommonSubsequence("abcde", "ace") << endl;
    
    // LIS
    vector<int> seq = {10, 9, 2, 5, 3, 7, 101, 18};
    cout << "LIS (O(n^2)): " << lis_dp(seq) << endl;
    cout << "LIS (O(nlogn)): " << lis_binary_search(seq) << endl;
    
    return 0;
}
