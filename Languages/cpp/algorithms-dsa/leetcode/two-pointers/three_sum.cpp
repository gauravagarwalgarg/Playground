/*
 * LeetCode 15 - 3Sum
 * Topic: Two Pointers
 * Difficulty: Medium
 *
 * Sort array, fix one element, use two pointers to find pairs. Skip duplicates.
 * Time: O(n^2)
 * Space: O(1) (excluding output)
 */
#include <iostream>
#include <vector>
#include <algorithm>
#include <cassert>
using namespace std;

vector<vector<int>> threeSum(vector<int>& nums) {
    sort(nums.begin(), nums.end());
    vector<vector<int>> res;
    for (int i = 0; i < (int)nums.size() - 2; i++) {
        if (i > 0 && nums[i] == nums[i-1]) continue;
        int lo = i + 1, hi = nums.size() - 1;
        while (lo < hi) {
            int sum = nums[i] + nums[lo] + nums[hi];
            if (sum < 0) lo++;
            else if (sum > 0) hi--;
            else {
                res.push_back({nums[i], nums[lo], nums[hi]});
                while (lo < hi && nums[lo] == nums[lo+1]) lo++;
                while (lo < hi && nums[hi] == nums[hi-1]) hi--;
                lo++; hi--;
            }
        }
    }
    return res;
}

int main() {
    vector<int> t1 = {-1,0,1,2,-1,-4};
    auto r1 = threeSum(t1);
    assert(r1.size() == 2);

    vector<int> t2 = {0,1,1};
    auto r2 = threeSum(t2);
    assert(r2.empty());

    vector<int> t3 = {0,0,0};
    auto r3 = threeSum(t3);
    assert(r3.size() == 1 && r3[0] == (vector<int>{0,0,0}));

    cout << "All tests passed!" << endl;
    return 0;
}
