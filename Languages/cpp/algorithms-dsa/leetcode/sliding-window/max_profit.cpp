/*
 * LeetCode 121 - Best Time to Buy and Sell Stock
 * Topic: Sliding Window
 * Difficulty: Easy
 *
 * Track minimum price seen so far and compute max profit at each step.
 * Time: O(n)
 * Space: O(1)
 */
#include <iostream>
#include <vector>
#include <climits>
#include <cassert>
using namespace std;

int maxProfit(vector<int>& prices) {
    int minPrice = INT_MAX, profit = 0;
    for (int p : prices) {
        minPrice = min(minPrice, p);
        profit = max(profit, p - minPrice);
    }
    return profit;
}

int main() {
    vector<int> t1 = {7,1,5,3,6,4};
    assert(maxProfit(t1) == 5);

    vector<int> t2 = {7,6,4,3,1};
    assert(maxProfit(t2) == 0);

    vector<int> t3 = {2,4,1};
    assert(maxProfit(t3) == 2);

    vector<int> t4 = {1,2};
    assert(maxProfit(t4) == 1);

    cout << "All tests passed!" << endl;
    return 0;
}
