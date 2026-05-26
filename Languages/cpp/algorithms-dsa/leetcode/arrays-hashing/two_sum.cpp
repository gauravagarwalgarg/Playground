/**
 * LeetCode 1: Two Sum
 * Topic: Arrays & Hashing
 * Difficulty: Easy
 *
 * Given an array of integers, return indices of two numbers that add up to target.
 * Time: O(n), Space: O(n)
 */
#include <iostream>
#include <vector>
#include <unordered_map>
#include <cassert>
using namespace std;

vector<int> twoSum(vector<int>& nums, int target) {
    unordered_map<int, int> seen;
    for (int i = 0; i < (int)nums.size(); i++) {
        int complement = target - nums[i];
        if (seen.count(complement)) {
            return {seen[complement], i};
        }
        seen[nums[i]] = i;
    }
    return {};
}

int main() {
    vector<int> v1 = {2, 7, 11, 15};
    auto r1 = twoSum(v1, 9);
    assert(r1[0] == 0 && r1[1] == 1);

    vector<int> v2 = {3, 2, 4};
    auto r2 = twoSum(v2, 6);
    assert(r2[0] == 1 && r2[1] == 2);

    cout << "All tests passed!" << endl;
    return 0;
}
