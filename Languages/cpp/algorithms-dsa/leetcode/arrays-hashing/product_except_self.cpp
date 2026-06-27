/*
 * LeetCode 238 - Product of Array Except Self
 * Topic: Arrays & Hashing
 * Difficulty: Medium
 *
 * Build prefix products left-to-right, then multiply suffix right-to-left.
 * Time: O(n)
 * Space: O(1) (output array not counted)
 */
#include <iostream>
#include <vector>
#include <cassert>
using namespace std;

vector<int> productExceptSelf(vector<int>& nums) {
    int n = nums.size();
    vector<int> res(n, 1);
    int prefix = 1;
    for (int i = 0; i < n; i++) {
        res[i] = prefix;
        prefix *= nums[i];
    }
    int suffix = 1;
    for (int i = n - 1; i >= 0; i--) {
        res[i] *= suffix;
        suffix *= nums[i];
    }
    return res;
}

int main() {
    vector<int> t1 = {1,2,3,4};
    assert(productExceptSelf(t1) == (vector<int>{24,12,8,6}));

    vector<int> t2 = {-1,1,0,-3,3};
    assert(productExceptSelf(t2) == (vector<int>{0,0,9,0,0}));

    vector<int> t3 = {2,2,2,2};
    assert(productExceptSelf(t3) == (vector<int>{8,8,8,8}));

    cout << "All tests passed!" << endl;
    return 0;
}
