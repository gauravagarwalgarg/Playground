/**
 * LeetCode 253: Meeting Rooms II
 * Topic: Intervals
 * Difficulty: Medium
 *
 * Given meeting time intervals, find the minimum number of conference rooms required.
 * Sweep line approach: sort start/end events separately.
 * Time: O(n log n), Space: O(n)
 */
#include <iostream>
#include <vector>
#include <algorithm>
#include <cassert>
using namespace std;

int minMeetingRooms(vector<vector<int>>& intervals) {
    vector<int> starts, ends;
    for (auto& iv : intervals) {
        starts.push_back(iv[0]);
        ends.push_back(iv[1]);
    }
    sort(starts.begin(), starts.end());
    sort(ends.begin(), ends.end());

    int rooms = 0, maxRooms = 0, e = 0;
    for (int s = 0; s < (int)starts.size(); s++) {
        if (starts[s] < ends[e]) {
            rooms++;
        } else {
            e++;
        }
        maxRooms = max(maxRooms, rooms);
    }
    return maxRooms;
}

int main() {
    vector<vector<int>> m1 = {{0,30},{5,10},{15,20}};
    assert(minMeetingRooms(m1) == 2);

    vector<vector<int>> m2 = {{7,10},{2,4}};
    assert(minMeetingRooms(m2) == 1);

    vector<vector<int>> m3 = {{0,5},{5,10},{10,15}};
    assert(minMeetingRooms(m3) == 1);

    vector<vector<int>> m4 = {{1,5},{2,6},{3,7}};
    assert(minMeetingRooms(m4) == 3);

    cout << "All tests passed!" << endl;
    return 0;
}
