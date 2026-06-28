/**
 * LeetCode 338: Counting Bits
 * Topic: Bit Manipulation
 * Difficulty: Easy
 *
 * Given integer n, return an array where ans[i] is the number of 1's in binary of i.
 * DP relation: dp[i] = dp[i >> 1] + (i & 1)
 * Time: O(n), Space: O(n)
 */
#include <iostream>
#include <vector>
#include <cassert>
using namespace std;

vector<int> countBits(int n) {
    vector<int> dp(n + 1, 0);
    for (int i = 1; i <= n; i++) {
        dp[i] = dp[i >> 1] + (i & 1);
    }
    return dp;
}

int main() {
    assert(countBits(2) == vector<int>({0, 1, 1}));
    assert(countBits(5) == vector<int>({0, 1, 1, 2, 1, 2}));
    assert(countBits(0) == vector<int>({0}));

    cout << "All tests passed!" << endl;
    return 0;
}
