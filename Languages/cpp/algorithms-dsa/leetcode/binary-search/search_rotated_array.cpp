/*
 * LeetCode 33 - Search in Rotated Sorted Array
 * Topic: Binary Search
 * Difficulty: Medium
 *
 * Modified binary search: determine which half is sorted, then decide direction.
 * Time: O(log n)
 * Space: O(1)
 */
#include <iostream>
#include <vector>
#include <cassert>
using namespace std;

int search(vector<int>& nums, int target) {
    int lo = 0, hi = nums.size() - 1;
    while (lo <= hi) {
        int mid = lo + (hi - lo) / 2;
        if (nums[mid] == target) return mid;
        if (nums[lo] <= nums[mid]) {
            if (target >= nums[lo] && target < nums[mid]) hi = mid - 1;
            else lo = mid + 1;
        } else {
            if (target > nums[mid] && target <= nums[hi]) lo = mid + 1;
            else hi = mid - 1;
        }
    }
    return -1;
}

int main() {
    vector<int> t1 = {4,5,6,7,0,1,2};
    assert(search(t1, 0) == 4);
    assert(search(t1, 3) == -1);

    vector<int> t2 = {1};
    assert(search(t2, 0) == -1);

    vector<int> t3 = {3,1};
    assert(search(t3, 1) == 1);

    cout << "All tests passed!" << endl;
    return 0;
}
