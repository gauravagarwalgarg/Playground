/**
 * Pattern: Sliding Window
 * Problem: Best Time to Buy and Sell Stock (LeetCode 121)
 *
 * Find max profit from one buy and one sell.
 * Time: O(n), Space: O(1)
 */
#include <iostream>
#include <vector>
#include <algorithm>
#include <cassert>
using namespace std;

int maxProfit(vector<int>& prices) {
    int minPrice = prices[0];
    int maxProfit = 0;
    for (int i = 1; i < (int)prices.size(); i++) {
        maxProfit = max(maxProfit, prices[i] - minPrice);
        minPrice = min(minPrice, prices[i]);
    }
    return maxProfit;
}

int main() {
    vector<int> p1 = {7, 1, 5, 3, 6, 4};
    assert(maxProfit(p1) == 5);

    vector<int> p2 = {7, 6, 4, 3, 1};
    assert(maxProfit(p2) == 0);

    cout << "All tests passed!" << endl;
    return 0;
}
