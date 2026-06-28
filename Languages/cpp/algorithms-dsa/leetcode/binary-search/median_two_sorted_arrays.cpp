/*
 * LeetCode 4 - Median of Two Sorted Arrays
 * Topic: Binary Search
 * Difficulty: Hard
 *
 * Binary search on the smaller array to find correct partition.
 * Time: O(log(min(m, n)))
 * Space: O(1)
 */
#include <iostream>
#include <vector>
#include <climits>
#include <cassert>
#include <cmath>
using namespace std;

double findMedianSortedArrays(vector<int>& nums1, vector<int>& nums2) {
    if (nums1.size() > nums2.size()) swap(nums1, nums2);
    int m = nums1.size(), n = nums2.size();
    int lo = 0, hi = m;
    while (lo <= hi) {
        int i = (lo + hi) / 2;
        int j = (m + n + 1) / 2 - i;
        int left1 = (i == 0) ? INT_MIN : nums1[i-1];
        int right1 = (i == m) ? INT_MAX : nums1[i];
        int left2 = (j == 0) ? INT_MIN : nums2[j-1];
        int right2 = (j == n) ? INT_MAX : nums2[j];
        if (left1 <= right2 && left2 <= right1) {
            if ((m + n) % 2 == 0)
                return (max(left1, left2) + min(right1, right2)) / 2.0;
            return max(left1, left2);
        } else if (left1 > right2) hi = i - 1;
        else lo = i + 1;
    }
    return 0.0;
}

int main() {
    vector<int> a1 = {1,3}, b1 = {2};
    assert(fabs(findMedianSortedArrays(a1, b1) - 2.0) < 1e-5);

    vector<int> a2 = {1,2}, b2 = {3,4};
    assert(fabs(findMedianSortedArrays(a2, b2) - 2.5) < 1e-5);

    vector<int> a3 = {}, b3 = {1};
    assert(fabs(findMedianSortedArrays(a3, b3) - 1.0) < 1e-5);

    cout << "All tests passed!" << endl;
    return 0;
}
