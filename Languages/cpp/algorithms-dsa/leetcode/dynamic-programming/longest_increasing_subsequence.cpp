/*
 * LeetCode 300 - Longest Increasing Subsequence
 * Topic: Dynamic Programming
 * Difficulty: Medium
 *
 * Patience sorting: maintain a tails array, use binary search to place elements.
 * Time: O(n log n)
 * Space: O(n)
 */
#include <iostream>
#include <vector>
#include <algorithm>
#include <cassert>
using namespace std;

int lengthOfLIS(vector<int>& nums) {
    vector<int> tails;
    for (int n : nums) {
        auto it = lower_bound(tails.begin(), tails.end(), n);
        if (it == tails.end()) tails.push_back(n);
        else *it = n;
    }
    return tails.size();
}

int main() {
    vector<int> t1 = {10,9,2,5,3,7,101,18};
    assert(lengthOfLIS(t1) == 4);

    vector<int> t2 = {0,1,0,3,2,3};
    assert(lengthOfLIS(t2) == 4);

    vector<int> t3 = {7,7,7,7};
    assert(lengthOfLIS(t3) == 1);

    vector<int> t4 = {1,3,6,7,9,4,10,5,6};
    assert(lengthOfLIS(t4) == 6);

    cout << "All tests passed!" << endl;
    return 0;
}
