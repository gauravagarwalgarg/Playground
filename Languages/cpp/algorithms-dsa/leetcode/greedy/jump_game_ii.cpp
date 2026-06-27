/**
 * LeetCode 45: Jump Game II
 * Topic: Greedy
 * Difficulty: Medium
 *
 * Return the minimum number of jumps to reach the last index.
 * BFS-style greedy: expand the farthest reachable window each level.
 * Time: O(n), Space: O(1)
 */
#include <iostream>
#include <vector>
#include <cassert>
using namespace std;

int jump(vector<int>& nums) {
    int jumps = 0, curEnd = 0, farthest = 0;
    for (int i = 0; i < (int)nums.size() - 1; i++) {
        farthest = max(farthest, i + nums[i]);
        if (i == curEnd) {
            jumps++;
            curEnd = farthest;
        }
    }
    return jumps;
}

int main() {
    vector<int> v1 = {2, 3, 1, 1, 4};
    assert(jump(v1) == 2);

    vector<int> v2 = {2, 3, 0, 1, 4};
    assert(jump(v2) == 2);

    vector<int> v3 = {1, 2, 3};
    assert(jump(v3) == 2);

    vector<int> v4 = {0};
    assert(jump(v4) == 0);

    cout << "All tests passed!" << endl;
    return 0;
}
