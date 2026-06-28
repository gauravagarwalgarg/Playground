/**
 * Pattern: Heaps / Priority Queues
 *
 * Key techniques covered:
 * 1. Top-K Elements (max-heap / min-heap of size K)
 * 2. Merge K Sorted Lists (min-heap)
 * 3. Kth Largest Element (min-heap of size K)
 * 4. Running Median (two heaps: max-heap for lower half, min-heap for upper half)
 *
 * Core insight: priority_queue gives O(log n) insert/extract.
 * - Max-heap (default): priority_queue<int>
 * - Min-heap: priority_queue<int, vector<int>, greater<int>>
 *
 * Compile: g++ -std=c++17 -o test heaps.cpp && ./test
 */
#include <iostream>
#include <vector>
#include <queue>
#include <algorithm>
#include <cassert>
using namespace std;

// ─────────────────────────────────────────────────────────────────────────────
// 1. Top-K Frequent Elements
//    Strategy: Use a min-heap of size K to keep the K largest frequencies.
//    Time: O(n log k), Space: O(n)
// ─────────────────────────────────────────────────────────────────────────────
#include <unordered_map>

vector<int> topKFrequent(vector<int>& nums, int k) {
    // Count frequencies
    unordered_map<int, int> freq;
    for (int n : nums) freq[n]++;

    // Min-heap of (frequency, element) keeps K most frequent
    using P = pair<int, int>;
    priority_queue<P, vector<P>, greater<P>> minHeap;

    for (auto& [num, count] : freq) {
        minHeap.push({count, num});
        if ((int)minHeap.size() > k) {
            minHeap.pop(); // Remove least frequent
        }
    }

    vector<int> result;
    while (!minHeap.empty()) {
        result.push_back(minHeap.top().second);
        minHeap.pop();
    }
    return result;
}

// ─────────────────────────────────────────────────────────────────────────────
// 2. Merge K Sorted Lists (simulated with vectors)
//    Strategy: Push the first element of each list into a min-heap.
//    Pop smallest, push next element from that list.
//    Time: O(N log K) where N = total elements, K = number of lists
// ─────────────────────────────────────────────────────────────────────────────
vector<int> mergeKSorted(vector<vector<int>>& lists) {
    // (value, list_index, element_index)
    using T = tuple<int, int, int>;
    priority_queue<T, vector<T>, greater<T>> minHeap;

    // Initialize heap with first element of each list
    for (int i = 0; i < (int)lists.size(); i++) {
        if (!lists[i].empty()) {
            minHeap.push({lists[i][0], i, 0});
        }
    }

    vector<int> result;
    while (!minHeap.empty()) {
        auto [val, li, ei] = minHeap.top();
        minHeap.pop();
        result.push_back(val);

        // If there's a next element in the same list, push it
        if (ei + 1 < (int)lists[li].size()) {
            minHeap.push({lists[li][ei + 1], li, ei + 1});
        }
    }
    return result;
}

// ─────────────────────────────────────────────────────────────────────────────
// 3. Kth Largest Element
//    Strategy: Maintain a min-heap of size K.
//    After processing all elements, the top is the Kth largest.
//    Time: O(n log k), Space: O(k)
// ─────────────────────────────────────────────────────────────────────────────
int kthLargest(vector<int>& nums, int k) {
    priority_queue<int, vector<int>, greater<int>> minHeap;

    for (int n : nums) {
        minHeap.push(n);
        if ((int)minHeap.size() > k) {
            minHeap.pop(); // Remove smallest, keeping K largest
        }
    }
    return minHeap.top();
}

// ─────────────────────────────────────────────────────────────────────────────
// 4. Running Median (Two Heaps)
//    Strategy:
//    - maxHeap: stores the smaller half (top = max of lower half)
//    - minHeap: stores the larger half (top = min of upper half)
//    - Balance: maxHeap.size() >= minHeap.size() (differ by at most 1)
//    Median = maxHeap.top() if odd count, else average of both tops.
//    Time per insert: O(log n), Space: O(n)
// ─────────────────────────────────────────────────────────────────────────────
class MedianFinder {
    priority_queue<int> maxHeap;                            // Lower half
    priority_queue<int, vector<int>, greater<int>> minHeap; // Upper half

public:
    void addNum(int num) {
        // Always add to maxHeap first
        maxHeap.push(num);

        // Ensure maxHeap's top <= minHeap's top
        if (!minHeap.empty() && maxHeap.top() > minHeap.top()) {
            minHeap.push(maxHeap.top());
            maxHeap.pop();
        }

        // Balance sizes: maxHeap can have at most 1 more than minHeap
        if ((int)maxHeap.size() > (int)minHeap.size() + 1) {
            minHeap.push(maxHeap.top());
            maxHeap.pop();
        } else if ((int)minHeap.size() > (int)maxHeap.size()) {
            maxHeap.push(minHeap.top());
            minHeap.pop();
        }
    }

    double findMedian() {
        if (maxHeap.size() > minHeap.size()) {
            return maxHeap.top();
        }
        return (maxHeap.top() + minHeap.top()) / 2.0;
    }
};

// ─────────────────────────────────────────────────────────────────────────────
int main() {
    // Test 1: Top-K Frequent
    {
        vector<int> nums = {1, 1, 1, 2, 2, 3};
        auto result = topKFrequent(nums, 2);
        sort(result.begin(), result.end());
        assert(result == vector<int>({1, 2}));
    }

    // Test 2: Merge K Sorted Lists
    {
        vector<vector<int>> lists = {{1, 4, 5}, {1, 3, 4}, {2, 6}};
        auto merged = mergeKSorted(lists);
        assert(merged == vector<int>({1, 1, 2, 3, 4, 4, 5, 6}));
    }

    // Test 3: Kth Largest
    {
        vector<int> nums = {3, 2, 1, 5, 6, 4};
        assert(kthLargest(nums, 2) == 5);

        vector<int> nums2 = {3, 2, 3, 1, 2, 4, 5, 5, 6};
        assert(kthLargest(nums2, 4) == 4);
    }

    // Test 4: Running Median
    {
        MedianFinder mf;
        mf.addNum(1);
        assert(mf.findMedian() == 1.0);
        mf.addNum(2);
        assert(mf.findMedian() == 1.5);
        mf.addNum(3);
        assert(mf.findMedian() == 2.0);
        mf.addNum(4);
        assert(mf.findMedian() == 2.5);
        mf.addNum(5);
        assert(mf.findMedian() == 3.0);
    }

    cout << "All heap pattern tests passed!" << endl;
    return 0;
}
