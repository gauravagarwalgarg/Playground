/**
 * LeetCode 56: Merge Intervals
 * Topic: Intervals
 * Difficulty: Medium
 *
 * Given an array of intervals, merge all overlapping intervals.
 * Sort by start, then merge consecutive overlapping ones.
 * Time: O(n log n), Space: O(n)
 */
#include <iostream>
#include <vector>
#include <algorithm>
#include <cassert>
using namespace std;

vector<vector<int>> merge(vector<vector<int>>& intervals) {
    sort(intervals.begin(), intervals.end());
    vector<vector<int>> res;
    for (auto& iv : intervals) {
        if (!res.empty() && res.back()[1] >= iv[0]) {
            res.back()[1] = max(res.back()[1], iv[1]);
        } else {
            res.push_back(iv);
        }
    }
    return res;
}

int main() {
    vector<vector<int>> i1 = {{1,3},{2,6},{8,10},{15,18}};
    auto r1 = merge(i1);
    assert(r1 == vector<vector<int>>({{1,6},{8,10},{15,18}}));

    vector<vector<int>> i2 = {{1,4},{4,5}};
    auto r2 = merge(i2);
    assert(r2 == vector<vector<int>>({{1,5}}));

    vector<vector<int>> i3 = {{1,4},{0,4}};
    auto r3 = merge(i3);
    assert(r3 == vector<vector<int>>({{0,4}}));

    cout << "All tests passed!" << endl;
    return 0;
}
