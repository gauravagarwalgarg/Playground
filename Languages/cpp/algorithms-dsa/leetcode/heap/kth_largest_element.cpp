/*
 * LeetCode 215 - Kth Largest Element in an Array
 * Topic: Heap
 * Difficulty: Medium
 *
 * Maintain a min-heap of size k. The top is the kth largest.
 * Time: O(n log k)
 * Space: O(k)
 */
#include <iostream>
#include <vector>
#include <queue>
#include <cassert>
using namespace std;

int findKthLargest(vector<int>& nums, int k) {
    priority_queue<int, vector<int>, greater<int>> minHeap;
    for (int n : nums) {
        minHeap.push(n);
        if ((int)minHeap.size() > k) minHeap.pop();
    }
    return minHeap.top();
}

int main() {
    vector<int> t1 = {3,2,1,5,6,4};
    assert(findKthLargest(t1, 2) == 5);

    vector<int> t2 = {3,2,3,1,2,4,5,5,6};
    assert(findKthLargest(t2, 4) == 4);

    vector<int> t3 = {1};
    assert(findKthLargest(t3, 1) == 1);

    vector<int> t4 = {7,6,5,4,3,2,1};
    assert(findKthLargest(t4, 5) == 3);

    cout << "All tests passed!" << endl;
    return 0;
}
