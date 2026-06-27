/*
 * LeetCode 198 - House Robber
 * Topic: Dynamic Programming
 * Difficulty: Medium
 *
 * At each house decide to rob (prev2 + current) or skip (prev1). O(1) space DP.
 * Time: O(n)
 * Space: O(1)
 */
#include <iostream>
#include <vector>
#include <cassert>
using namespace std;

int rob(vector<int>& nums) {
    int prev2 = 0, prev1 = 0;
    for (int n : nums) {
        int curr = max(prev1, prev2 + n);
        prev2 = prev1;
        prev1 = curr;
    }
    return prev1;
}

int main() {
    vector<int> t1 = {1,2,3,1};
    assert(rob(t1) == 4);

    vector<int> t2 = {2,7,9,3,1};
    assert(rob(t2) == 12);

    vector<int> t3 = {2,1,1,2};
    assert(rob(t3) == 4);

    vector<int> t4 = {0};
    assert(rob(t4) == 0);

    cout << "All tests passed!" << endl;
    return 0;
}
