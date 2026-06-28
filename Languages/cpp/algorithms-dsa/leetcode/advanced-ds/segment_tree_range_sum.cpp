/*
 * LeetCode 307 - Range Sum Query - Mutable
 * Topic: Advanced Data Structures (Segment Tree)
 * Difficulty: Medium
 *
 * Segment tree supporting point updates and range sum queries.
 * Tree stored in array: parent i has children 2i and 2i+1.
 *
 * Time: O(log n) per update and query
 * Space: O(n)
 */
#include <iostream>
#include <vector>
#include <cassert>
using namespace std;

class NumArray {
    vector<int> tree;
    int n;

    void build(vector<int>& nums, int node, int start, int end) {
        if (start == end) {
            tree[node] = nums[start];
            return;
        }
        int mid = (start + end) / 2;
        build(nums, 2 * node, start, mid);
        build(nums, 2 * node + 1, mid + 1, end);
        tree[node] = tree[2 * node] + tree[2 * node + 1];
    }

    void updateHelper(int node, int start, int end, int idx, int val) {
        if (start == end) {
            tree[node] = val;
            return;
        }
        int mid = (start + end) / 2;
        if (idx <= mid)
            updateHelper(2 * node, start, mid, idx, val);
        else
            updateHelper(2 * node + 1, mid + 1, end, idx, val);
        tree[node] = tree[2 * node] + tree[2 * node + 1];
    }

    int queryHelper(int node, int start, int end, int l, int r) {
        if (r < start || end < l) return 0;          // No overlap
        if (l <= start && end <= r) return tree[node]; // Total overlap
        int mid = (start + end) / 2;
        return queryHelper(2 * node, start, mid, l, r) +
               queryHelper(2 * node + 1, mid + 1, end, l, r);
    }

public:
    NumArray(vector<int>& nums) {
        n = nums.size();
        tree.resize(4 * n, 0);
        if (n > 0) build(nums, 1, 0, n - 1);
    }

    void update(int index, int val) {
        updateHelper(1, 0, n - 1, index, val);
    }

    int sumRange(int left, int right) {
        return queryHelper(1, 0, n - 1, left, right);
    }
};

int main() {
    // Test 1: Basic operations
    vector<int> nums = {1, 3, 5};
    NumArray obj(nums);
    assert(obj.sumRange(0, 2) == 9);  // 1 + 3 + 5
    obj.update(1, 2);                  // nums = [1, 2, 5]
    assert(obj.sumRange(0, 2) == 8);  // 1 + 2 + 5

    // Test 2: Single element queries
    assert(obj.sumRange(0, 0) == 1);
    assert(obj.sumRange(1, 1) == 2);
    assert(obj.sumRange(2, 2) == 5);

    // Test 3: Partial range
    assert(obj.sumRange(0, 1) == 3);  // 1 + 2
    assert(obj.sumRange(1, 2) == 7);  // 2 + 5

    // Test 4: Update and verify
    obj.update(0, 10);                // nums = [10, 2, 5]
    assert(obj.sumRange(0, 2) == 17);

    // Test 5: Larger array
    vector<int> nums2 = {1, 2, 3, 4, 5, 6, 7, 8};
    NumArray obj2(nums2);
    assert(obj2.sumRange(0, 7) == 36);
    assert(obj2.sumRange(2, 5) == 18); // 3+4+5+6
    obj2.update(3, 10);               // nums = [1,2,3,10,5,6,7,8]
    assert(obj2.sumRange(2, 5) == 24); // 3+10+5+6

    cout << "All tests passed!" << endl;
    return 0;
}
