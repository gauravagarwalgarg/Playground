/**
 * LeetCode 55: Jump Game
 * Topic: Greedy
 * Difficulty: Medium
 *
 * Given an array where each element is the max jump length from that position,
 * determine if you can reach the last index. Track farthest reachable position.
 * Time: O(n), Space: O(1)
 */
#include <iostream>
#include <vector>
#include <cassert>
using namespace std;

bool canJump(vector<int>& nums) {
    int farthest = 0;
    for (int i = 0; i < (int)nums.size(); i++) {
        if (i > farthest) return false;
        farthest = max(farthest, i + nums[i]);
    }
    return true;
}

int main() {
    vector<int> v1 = {2, 3, 1, 1, 4};
    assert(canJump(v1) == true);

    vector<int> v2 = {3, 2, 1, 0, 4};
    assert(canJump(v2) == false);

    vector<int> v3 = {0};
    assert(canJump(v3) == true);

    vector<int> v4 = {2, 0, 0};
    assert(canJump(v4) == true);

    cout << "All tests passed!" << endl;
    return 0;
}
