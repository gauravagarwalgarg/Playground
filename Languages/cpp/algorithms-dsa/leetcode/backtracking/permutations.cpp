/**
 * LeetCode 46: Permutations
 * Topic: Backtracking
 * Difficulty: Medium
 *
 * Given an array of distinct integers, return all possible permutations.
 * Swap-based backtracking approach.
 * Time: O(n!), Space: O(n) recursion depth
 */
#include <iostream>
#include <vector>
#include <cassert>
using namespace std;

void backtrack(vector<int>& nums, int start, vector<vector<int>>& res) {
    if (start == (int)nums.size()) {
        res.push_back(nums);
        return;
    }
    for (int i = start; i < (int)nums.size(); i++) {
        swap(nums[start], nums[i]);
        backtrack(nums, start + 1, res);
        swap(nums[start], nums[i]);
    }
}

vector<vector<int>> permute(vector<int>& nums) {
    vector<vector<int>> res;
    backtrack(nums, 0, res);
    return res;
}

int main() {
    vector<int> v1 = {1, 2, 3};
    auto r1 = permute(v1);
    assert(r1.size() == 6);

    vector<int> v2 = {0, 1};
    auto r2 = permute(v2);
    assert(r2.size() == 2);

    vector<int> v3 = {1};
    auto r3 = permute(v3);
    assert(r3.size() == 1);

    cout << "All tests passed!" << endl;
    return 0;
}
