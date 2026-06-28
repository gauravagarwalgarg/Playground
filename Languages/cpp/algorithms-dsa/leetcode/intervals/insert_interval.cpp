/**
 * LeetCode 57: Insert Interval
 * Topic: Intervals
 * Difficulty: Medium
 *
 * Insert a new interval into a sorted list of non-overlapping intervals,
 * merging if necessary. Find overlap region and merge.
 * Time: O(n), Space: O(n)
 */
#include <iostream>
#include <vector>
#include <cassert>
using namespace std;

vector<vector<int>> insert(vector<vector<int>>& intervals, vector<int>& newInterval) {
    vector<vector<int>> res;
    int i = 0, n = intervals.size();

    // Add all intervals before newInterval
    while (i < n && intervals[i][1] < newInterval[0])
        res.push_back(intervals[i++]);

    // Merge overlapping intervals
    while (i < n && intervals[i][0] <= newInterval[1]) {
        newInterval[0] = min(newInterval[0], intervals[i][0]);
        newInterval[1] = max(newInterval[1], intervals[i][1]);
        i++;
    }
    res.push_back(newInterval);

    // Add remaining intervals
    while (i < n)
        res.push_back(intervals[i++]);

    return res;
}

int main() {
    vector<vector<int>> i1 = {{1,3},{6,9}};
    vector<int> n1 = {2,5};
    auto r1 = insert(i1, n1);
    assert(r1 == vector<vector<int>>({{1,5},{6,9}}));

    vector<vector<int>> i2 = {{1,2},{3,5},{6,7},{8,10},{12,16}};
    vector<int> n2 = {4,8};
    auto r2 = insert(i2, n2);
    assert(r2 == vector<vector<int>>({{1,2},{3,10},{12,16}}));

    vector<vector<int>> i3 = {};
    vector<int> n3 = {5,7};
    auto r3 = insert(i3, n3);
    assert(r3 == vector<vector<int>>({{5,7}}));

    cout << "All tests passed!" << endl;
    return 0;
}
