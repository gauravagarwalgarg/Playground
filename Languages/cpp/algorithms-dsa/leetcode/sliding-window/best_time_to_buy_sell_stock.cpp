/**
 * LeetCode 121: Best Time to Buy and Sell Stock
 * Topic: Sliding Window
 * Difficulty: Easy
 *
 * You are given an array prices where prices[i] is the price of a stock on the ith day.
 * You want to maximize profit by choosing a single day to buy and a different day in the
 * future to sell. Return the maximum profit, or 0 if no profit can be achieved.
 *
 * Approach: Track the minimum price seen so far, and at each step calculate
 * the profit if we sell at the current price. Update max profit accordingly.
 *
 * Time: O(n), Space: O(1)
 */
#include <iostream>
#include <vector>
#include <algorithm>
#include <cassert>
#include <climits>
using namespace std;

int maxProfit(vector<int>& prices) {
    int minPrice = INT_MAX;
    int maxProf = 0;

    for (int price : prices) {
        if (price < minPrice) {
            minPrice = price;
        } else {
            maxProf = max(maxProf, price - minPrice);
        }
    }

    return maxProf;
}

int main() {
    vector<int> v1 = {7, 1, 5, 3, 6, 4};
    assert(maxProfit(v1) == 5);

    vector<int> v2 = {7, 6, 4, 3, 1};
    assert(maxProfit(v2) == 0);

    vector<int> v3 = {1, 2};
    assert(maxProfit(v3) == 1);

    vector<int> v4 = {2, 1, 2, 1, 0, 1, 2};
    assert(maxProfit(v4) == 2);

    vector<int> v5 = {1};
    assert(maxProfit(v5) == 0);

    cout << "All tests passed!" << endl;
    return 0;
}
