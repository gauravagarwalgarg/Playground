/*
 * LeetCode 153 - Find Minimum in Rotated Sorted Array
 * Topic: Binary Search
 * Difficulty: Medium
 *
 * Binary search for the pivot point where rotation occurs.
 * Time: O(log n)
 * Space: O(1)
 */
#include <iostream>
#include <vector>
#include <cassert>
using namespace std;

int findMin(vector<int>& nums) {
    int lo = 0, hi = nums.size() - 1;
    while (lo < hi) {
        int mid = lo + (hi - lo) / 2;
        if (nums[mid] > nums[hi]) lo = mid + 1;
        else hi = mid;
    }
    return nums[lo];
}

int main() {
    vector<int> t1 = {3,4,5,1,2};
    assert(findMin(t1) == 1);

    vector<int> t2 = {4,5,6,7,0,1,2};
    assert(findMin(t2) == 0);

    vector<int> t3 = {11,13,15,17};
    assert(findMin(t3) == 11);

    vector<int> t4 = {2,1};
    assert(findMin(t4) == 1);

    cout << "All tests passed!" << endl;
    return 0;
}
