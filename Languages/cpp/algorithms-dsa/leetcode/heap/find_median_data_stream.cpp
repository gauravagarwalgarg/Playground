/*
 * LeetCode 295 - Find Median from Data Stream
 * Topic: Heap
 * Difficulty: Hard
 *
 * Two heaps: max-heap for lower half, min-heap for upper half. Balance sizes.
 * Time: O(log n) per addNum, O(1) per findMedian
 * Space: O(n)
 */
#include <iostream>
#include <queue>
#include <cassert>
#include <cmath>
using namespace std;

class MedianFinder {
    priority_queue<int> lo;                              // max-heap
    priority_queue<int, vector<int>, greater<int>> hi;   // min-heap
public:
    void addNum(int num) {
        lo.push(num);
        hi.push(lo.top()); lo.pop();
        if (hi.size() > lo.size()) { lo.push(hi.top()); hi.pop(); }
    }
    double findMedian() {
        return lo.size() > hi.size() ? lo.top() : (lo.top() + hi.top()) / 2.0;
    }
};

int main() {
    MedianFinder mf;
    mf.addNum(1);
    assert(fabs(mf.findMedian() - 1.0) < 1e-5);
    mf.addNum(2);
    assert(fabs(mf.findMedian() - 1.5) < 1e-5);
    mf.addNum(3);
    assert(fabs(mf.findMedian() - 2.0) < 1e-5);
    mf.addNum(4);
    assert(fabs(mf.findMedian() - 2.5) < 1e-5);

    cout << "All tests passed!" << endl;
    return 0;
}
