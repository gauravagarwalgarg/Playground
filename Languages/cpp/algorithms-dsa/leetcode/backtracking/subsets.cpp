/**
 * LeetCode 78: Subsets
 * Topic: Backtracking
 * Difficulty: Medium
 *
 * Given an integer array of unique elements, return all possible subsets.
 * Use recursive include/exclude approach.
 * Time: O(2^n), Space: O(n) recursion depth
 */
#include <iostream>
#include <vector>
#include <cassert>
using namespace std;

void backtrack(vector<int>& nums, int idx, vector<int>& curr, vector<vector<int>>& res) {
    if (idx == (int)nums.size()) {
        res.push_back(curr);
        return;
    }
    // Exclude nums[idx]
    backtrack(nums, idx + 1, curr, res);
    // Include nums[idx]
    curr.push_back(nums[idx]);
    backtrack(nums, idx + 1, curr, res);
    curr.pop_back();
}

vector<vector<int>> subsets(vector<int>& nums) {
    vector<vector<int>> res;
    vector<int> curr;
    backtrack(nums, 0, curr, res);
    return res;
}

int main() {
    vector<int> v1 = {1, 2, 3};
    auto r1 = subsets(v1);
    assert(r1.size() == 8);

    vector<int> v2 = {0};
    auto r2 = subsets(v2);
    assert(r2.size() == 2);

    vector<int> v3 = {};
    auto r3 = subsets(v3);
    assert(r3.size() == 1);

    cout << "All tests passed!" << endl;
    return 0;
}
