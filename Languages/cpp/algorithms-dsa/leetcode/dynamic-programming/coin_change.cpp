/*
 * LeetCode 322 - Coin Change
 * Topic: Dynamic Programming
 * Difficulty: Medium
 *
 * Bottom-up DP: dp[i] = min coins needed to make amount i.
 * Time: O(amount * n) where n = number of coins
 * Space: O(amount)
 */
#include <iostream>
#include <vector>
#include <cassert>
using namespace std;

int coinChange(vector<int>& coins, int amount) {
    vector<int> dp(amount + 1, amount + 1);
    dp[0] = 0;
    for (int i = 1; i <= amount; i++)
        for (int c : coins)
            if (c <= i) dp[i] = min(dp[i], dp[i - c] + 1);
    return dp[amount] > amount ? -1 : dp[amount];
}

int main() {
    vector<int> c1 = {1,5,10};
    assert(coinChange(c1, 11) == 2);

    vector<int> c2 = {2};
    assert(coinChange(c2, 3) == -1);

    vector<int> c3 = {1};
    assert(coinChange(c3, 0) == 0);

    vector<int> c4 = {1,2,5};
    assert(coinChange(c4, 11) == 3);

    cout << "All tests passed!" << endl;
    return 0;
}
