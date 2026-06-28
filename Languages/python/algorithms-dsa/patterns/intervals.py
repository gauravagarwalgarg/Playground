"""
Intervals Pattern Templates

Core interval-based patterns for interview problems:
1. Merge Overlapping Intervals sort by start, merge greedily
2. Insert Interval find position, merge with overlaps
3. Meeting Rooms I can a person attend all meetings?
4. Meeting Rooms II minimum number of rooms needed

Key insight: Sort intervals by start time, then process linearly.
Two intervals [a,b] and [c,d] overlap when a <= d and c <= b.

Run: python intervals.py
"""

from typing import List


# =============================================================================
# TEMPLATE 1: Merge Overlapping Intervals
# Use when: consolidate overlapping ranges
# Time: O(n log n), Space: O(n)
# =============================================================================
def merge_intervals(intervals: List[List[int]]) -> List[List[int]]:
    """Merge all overlapping intervals."""
    if not intervals:
        return []

    # Sort by start time
    intervals.sort(key=lambda x: x[0])
    merged = [intervals[0]]

    for current in intervals[1:]:
        last = merged[-1]

        if current[0] <= last[1]:
            # Overlapping extend the end
            last[1] = max(last[1], current[1])
        else:
            # Non-overlapping add new interval
            merged.append(current)

    return merged


# =============================================================================
# TEMPLATE 2: Insert Interval
# Use when: insert a new interval into a sorted, non-overlapping list
# Time: O(n), Space: O(n)
# =============================================================================
def insert_interval(
    intervals: List[List[int]], new_interval: List[int]
) -> List[List[int]]:
    """Insert a new interval and merge if necessary."""
    result = []
    i = 0
    n = len(intervals)

    # 1. Add all intervals that come before new_interval
    while i < n and intervals[i][1] < new_interval[0]:
        result.append(intervals[i])
        i += 1

    # 2. Merge all overlapping intervals with new_interval
    while i < n and intervals[i][0] <= new_interval[1]:
        new_interval[0] = min(new_interval[0], intervals[i][0])
        new_interval[1] = max(new_interval[1], intervals[i][1])
        i += 1
    result.append(new_interval)

    # 3. Add remaining intervals
    while i < n:
        result.append(intervals[i])
        i += 1

    return result


# =============================================================================
# TEMPLATE 3: Meeting Rooms I (Can Attend All)
# Use when: check if any intervals overlap
# Time: O(n log n), Space: O(1)
# =============================================================================
def can_attend_all_meetings(intervals: List[List[int]]) -> bool:
    """Return True if a person can attend all meetings (no overlaps)."""
    intervals.sort(key=lambda x: x[0])

    for i in range(1, len(intervals)):
        # If current meeting starts before previous ends → conflict
        if intervals[i][0] < intervals[i - 1][1]:
            return False

    return True


# =============================================================================
# TEMPLATE 4: Meeting Rooms II (Minimum Rooms)
# Use when: find max number of concurrent overlapping intervals
# Approach 1: Sort start/end times separately (sweep line)
# Time: O(n log n), Space: O(n)
# =============================================================================
def min_meeting_rooms(intervals: List[List[int]]) -> int:
    """Find minimum number of meeting rooms required."""
    if not intervals:
        return 0

    # Sweep line approach: track events
    starts = sorted(i[0] for i in intervals)
    ends = sorted(i[1] for i in intervals)

    rooms_needed = 0
    max_rooms = 0
    s, e = 0, 0

    while s < len(starts):
        if starts[s] < ends[e]:
            # A meeting starts before the earliest one ends → need a room
            rooms_needed += 1
            max_rooms = max(max_rooms, rooms_needed)
            s += 1
        else:
            # A meeting ended → free up a room
            rooms_needed -= 1
            e += 1

    return max_rooms


def min_meeting_rooms_heap(intervals: List[List[int]]) -> int:
    """Alternative: use a min-heap to track end times of active meetings."""
    import heapq

    if not intervals:
        return 0

    intervals.sort(key=lambda x: x[0])
    # Heap stores end times of ongoing meetings
    heap = []

    for interval in intervals:
        # If earliest ending meeting has ended before this one starts
        if heap and heap[0] <= interval[0]:
            heapq.heappop(heap)  # reuse that room

        heapq.heappush(heap, interval[1])  # add current meeting's end time

    return len(heap)  # heap size = rooms in use


# =============================================================================
def main():
    # --- Merge Intervals ---
    assert merge_intervals([[1, 3], [2, 6], [8, 10], [15, 18]]) == [
        [1, 6], [8, 10], [15, 18]
    ]
    assert merge_intervals([[1, 4], [4, 5]]) == [[1, 5]]
    assert merge_intervals([[1, 4], [0, 4]]) == [[0, 4]]
    assert merge_intervals([]) == []
    print("PASS: merge_intervals")

    # --- Insert Interval ---
    assert insert_interval([[1, 3], [6, 9]], [2, 5]) == [[1, 5], [6, 9]]
    assert insert_interval(
        [[1, 2], [3, 5], [6, 7], [8, 10], [12, 16]], [4, 8]
    ) == [[1, 2], [3, 10], [12, 16]]
    assert insert_interval([], [5, 7]) == [[5, 7]]
    print("PASS: insert_interval")

    # --- Meeting Rooms I ---
    assert can_attend_all_meetings([[0, 30], [5, 10], [15, 20]]) is False
    assert can_attend_all_meetings([[7, 10], [2, 4]]) is True
    assert can_attend_all_meetings([[1, 5], [5, 10]]) is True  # no overlap at boundary
    print("PASS: can_attend_all_meetings")

    # --- Meeting Rooms II (sweep line) ---
    assert min_meeting_rooms([[0, 30], [5, 10], [15, 20]]) == 2
    assert min_meeting_rooms([[7, 10], [2, 4]]) == 1
    assert min_meeting_rooms([[1, 5], [2, 6], [3, 7], [4, 8]]) == 4
    assert min_meeting_rooms([]) == 0
    print("PASS: min_meeting_rooms (sweep line)")

    # --- Meeting Rooms II (heap) ---
    assert min_meeting_rooms_heap([[0, 30], [5, 10], [15, 20]]) == 2
    assert min_meeting_rooms_heap([[7, 10], [2, 4]]) == 1
    assert min_meeting_rooms_heap([[1, 5], [2, 6], [3, 7], [4, 8]]) == 4
    print("PASS: min_meeting_rooms_heap")

    print("\n✓ All interval pattern tests passed!")


if __name__ == "__main__":
    main()
