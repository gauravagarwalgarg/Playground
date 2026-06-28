/**
 * Pattern: Intervals
 *
 * Key techniques covered:
 * 1. Merge Intervals (sort by start, merge overlapping)
 * 2. Insert Interval (three-phase: before, overlap, after)
 * 3. Can Attend All Meetings (sort, check overlaps)
 * 4. Minimum Meeting Rooms (sweep line / event counting)
 *
 * Core insight: Sort intervals by start time. Overlapping intervals
 * have the property: current.start <= prev.end.
 *
 * Compile: g++ -std=c++17 -o test intervals.cpp && ./test
 */
#include <iostream>
#include <vector>
#include <algorithm>
#include <cassert>
using namespace std;

// ─────────────────────────────────────────────────────────────────────────────
// 1. Merge Intervals
//    Merge all overlapping intervals.
//    Strategy: Sort by start. If current overlaps with last merged, extend it.
//    Time: O(n log n), Space: O(n)
// ─────────────────────────────────────────────────────────────────────────────
vector<vector<int>> mergeIntervals(vector<vector<int>>& intervals) {
    if (intervals.empty()) return {};

    sort(intervals.begin(), intervals.end());
    vector<vector<int>> merged = {intervals[0]};

    for (int i = 1; i < (int)intervals.size(); i++) {
        auto& last = merged.back();
        if (intervals[i][0] <= last[1]) {
            // Overlapping extend the end
            last[1] = max(last[1], intervals[i][1]);
        } else {
            // Non-overlapping add new interval
            merged.push_back(intervals[i]);
        }
    }
    return merged;
}

// ─────────────────────────────────────────────────────────────────────────────
// 2. Insert Interval
//    Insert a new interval into a sorted list and merge if necessary.
//    Strategy: Three phases add all before, merge overlapping, add all after.
//    Time: O(n), Space: O(n)
// ─────────────────────────────────────────────────────────────────────────────
vector<vector<int>> insertInterval(vector<vector<int>>& intervals, vector<int> newInterval) {
    vector<vector<int>> result;
    int i = 0, n = intervals.size();

    // Phase 1: Add all intervals that come before newInterval
    while (i < n && intervals[i][1] < newInterval[0]) {
        result.push_back(intervals[i]);
        i++;
    }

    // Phase 2: Merge overlapping intervals with newInterval
    while (i < n && intervals[i][0] <= newInterval[1]) {
        newInterval[0] = min(newInterval[0], intervals[i][0]);
        newInterval[1] = max(newInterval[1], intervals[i][1]);
        i++;
    }
    result.push_back(newInterval);

    // Phase 3: Add all intervals that come after
    while (i < n) {
        result.push_back(intervals[i]);
        i++;
    }
    return result;
}

// ─────────────────────────────────────────────────────────────────────────────
// 3. Can Attend All Meetings
//    Check if a person can attend all meetings (no overlaps).
//    Strategy: Sort by start time, check if any meeting starts before previous ends.
//    Time: O(n log n), Space: O(1)
// ─────────────────────────────────────────────────────────────────────────────
bool canAttendAll(vector<vector<int>>& intervals) {
    sort(intervals.begin(), intervals.end());

    for (int i = 1; i < (int)intervals.size(); i++) {
        if (intervals[i][0] < intervals[i - 1][1]) {
            return false; // Overlap detected
        }
    }
    return true;
}

// ─────────────────────────────────────────────────────────────────────────────
// 4. Minimum Meeting Rooms (Sweep Line / Event Counting)
//    Find the minimum number of conference rooms needed.
//    Strategy: Create start events (+1) and end events (-1).
//    Sort all events. The maximum running sum is the answer.
//    Time: O(n log n), Space: O(n)
//
//    Why sweep line? It reduces 2D interval overlap counting to a 1D scan.
// ─────────────────────────────────────────────────────────────────────────────
int minMeetingRooms(vector<vector<int>>& intervals) {
    vector<pair<int, int>> events; // (time, +1 for start / -1 for end)

    for (auto& interval : intervals) {
        events.push_back({interval[0], +1}); // Meeting starts
        events.push_back({interval[1], -1}); // Meeting ends
    }

    // Sort by time; if same time, process ends (-1) before starts (+1)
    // This prevents counting a room as used when one meeting ends
    // at the same time another begins.
    sort(events.begin(), events.end(), [](auto& a, auto& b) {
        if (a.first != b.first) return a.first < b.first;
        return a.second < b.second; // -1 before +1
    });

    int rooms = 0, maxRooms = 0;
    for (auto& [time, delta] : events) {
        rooms += delta;
        maxRooms = max(maxRooms, rooms);
    }
    return maxRooms;
}

// ─────────────────────────────────────────────────────────────────────────────
int main() {
    // Test 1: Merge Intervals
    {
        vector<vector<int>> intervals = {{1, 3}, {2, 6}, {8, 10}, {15, 18}};
        auto result = mergeIntervals(intervals);
        assert(result == vector<vector<int>>({{1, 6}, {8, 10}, {15, 18}}));

        vector<vector<int>> intervals2 = {{1, 4}, {4, 5}};
        auto result2 = mergeIntervals(intervals2);
        assert(result2 == vector<vector<int>>({{1, 5}}));
    }

    // Test 2: Insert Interval
    {
        vector<vector<int>> intervals = {{1, 3}, {6, 9}};
        auto result = insertInterval(intervals, {2, 5});
        assert(result == vector<vector<int>>({{1, 5}, {6, 9}}));

        vector<vector<int>> intervals2 = {{1, 2}, {3, 5}, {6, 7}, {8, 10}, {12, 16}};
        auto result2 = insertInterval(intervals2, {4, 8});
        assert(result2 == vector<vector<int>>({{1, 2}, {3, 10}, {12, 16}}));
    }

    // Test 3: Can Attend All Meetings
    {
        vector<vector<int>> meetings1 = {{0, 30}, {5, 10}, {15, 20}};
        assert(canAttendAll(meetings1) == false);

        vector<vector<int>> meetings2 = {{7, 10}, {2, 4}};
        assert(canAttendAll(meetings2) == true);
    }

    // Test 4: Minimum Meeting Rooms
    {
        vector<vector<int>> meetings1 = {{0, 30}, {5, 10}, {15, 20}};
        assert(minMeetingRooms(meetings1) == 2);

        vector<vector<int>> meetings2 = {{7, 10}, {2, 4}};
        assert(minMeetingRooms(meetings2) == 1);

        vector<vector<int>> meetings3 = {{1, 5}, {2, 6}, {3, 7}, {4, 8}};
        assert(minMeetingRooms(meetings3) == 4);
    }

    cout << "All interval pattern tests passed!" << endl;
    return 0;
}
