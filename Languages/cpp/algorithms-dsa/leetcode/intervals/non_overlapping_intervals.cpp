/**
 * LeetCode 435: Non-overlapping Intervals
 * Topic: Intervals
 * Difficulty: Medium
 *
 * Return the minimum number of intervals to remove to make the rest non-overlapping.
 * Sort by end time, greedily keep intervals with earliest end.
 * Time: O(n log n), Space: O(1)
 */
#include <iostream>
#include <vector>
#include <algorithm>
#include <cassert>
using namespace std;

int eraseOverlapIntervals(vector<vector<int>>& intervals) {
    if (intervals.empty()) return 0;
    sort(intervals.begin(), intervals.end(), [](auto& a, auto& b) {
        return a[1] < b[1];
    });
    int count = 0, prevEnd = intervals[0][1];
    for (int i = 1; i < (int)intervals.size(); i++) {
        if (intervals[i][0] < prevEnd) {
            count++;
        } else {
            prevEnd = intervals[i][1];
        }
    }
    return count;
}

int main() {
    vector<vector<int>> i1 = {{1,2},{2,3},{3,4},{1,3}};
    assert(eraseOverlapIntervals(i1) == 1);

    vector<vector<int>> i2 = {{1,2},{1,2},{1,2}};
    assert(eraseOverlapIntervals(i2) == 2);

    vector<vector<int>> i3 = {{1,2},{2,3}};
    assert(eraseOverlapIntervals(i3) == 0);

    cout << "All tests passed!" << endl;
    return 0;
}
