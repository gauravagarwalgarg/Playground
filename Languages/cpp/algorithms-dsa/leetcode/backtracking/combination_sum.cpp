/**
 * LeetCode 39: Combination Sum
 * Topic: Backtracking
 * Difficulty: Medium
 *
 * Given candidates and a target, find all unique combinations that sum to target.
 * Each number may be used unlimited times. Backtrack with remaining target.
 * Time: O(2^t) where t = target/min(candidates), Space: O(t)
 */
#include <iostream>
#include <vector>
#include <cassert>
using namespace std;

void backtrack(vector<int>& cands, int idx, int remain, vector<int>& curr, vector<vector<int>>& res) {
    if (remain == 0) {
        res.push_back(curr);
        return;
    }
    for (int i = idx; i < (int)cands.size(); i++) {
        if (cands[i] > remain) break;
        curr.push_back(cands[i]);
        backtrack(cands, i, remain - cands[i], curr, res);
        curr.pop_back();
    }
}

vector<vector<int>> combinationSum(vector<int>& candidates, int target) {
    sort(candidates.begin(), candidates.end());
    vector<vector<int>> res;
    vector<int> curr;
    backtrack(candidates, 0, target, curr, res);
    return res;
}

int main() {
    vector<int> c1 = {2, 3, 6, 7};
    auto r1 = combinationSum(c1, 7);
    assert(r1.size() == 2); // [2,2,3] and [7]

    vector<int> c2 = {2, 3, 5};
    auto r2 = combinationSum(c2, 8);
    assert(r2.size() == 3); // [2,2,2,2], [2,3,3], [3,5]

    vector<int> c3 = {2};
    auto r3 = combinationSum(c3, 1);
    assert(r3.size() == 0);

    cout << "All tests passed!" << endl;
    return 0;
}
