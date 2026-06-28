"""
LeetCode #253 - Meeting Rooms II
Topic: Intervals
Difficulty: Medium

Find minimum number of meeting rooms required using sweep line / min-heap.

Time Complexity: O(n log n)
Space Complexity: O(n)
"""
import heapq


def min_meeting_rooms(intervals: list[list[int]]) -> int:
    if not intervals:
        return 0
    intervals.sort(key=lambda x: x[0])
    heap: list[int] = []  # end times
    for start, end in intervals:
        if heap and heap[0] <= start:
            heapq.heappop(heap)
        heapq.heappush(heap, end)
    return len(heap)


if __name__ == "__main__":
    assert min_meeting_rooms([[0, 30], [5, 10], [15, 20]]) == 2
    assert min_meeting_rooms([[7, 10], [2, 4]]) == 1
    assert min_meeting_rooms([[1, 5], [2, 6], [3, 7]]) == 3
    assert min_meeting_rooms([]) == 0
    print("All tests passed!")
