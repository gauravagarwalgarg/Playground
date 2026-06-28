import java.util.*;

/**
 * Interval Patterns
 * 
 * Key insight: Sort intervals by start time, then process sequentially.
 * Overlapping intervals have: a.start <= b.end && b.start <= a.end
 * 
 * Common patterns:
 * 1. Merge Intervals: Sort by start, merge overlapping
 * 2. Insert Interval: Find position, merge with overlapping
 * 3. Meeting Rooms I: Check if any intervals overlap (can attend all)
 * 4. Meeting Rooms II: Sweep line / min-heap for concurrent meetings
 * 
 * Sweep Line technique:
 * - Create events for start (+1) and end (-1) of each interval
 * - Sort events by time
 * - Track running count; maximum = peak concurrent intervals
 */
public class Intervals {

    // ==================== Merge Intervals ====================
    /**
     * Merge all overlapping intervals.
     * Strategy: Sort by start time, then greedily extend or create new intervals.
     * 
     * Time: O(n log n), Space: O(n)
     */
    public static int[][] mergeIntervals(int[][] intervals) {
        if (intervals.length <= 1) return intervals;

        // Sort by start time
        Arrays.sort(intervals, (a, b) -> a[0] - b[0]);

        List<int[]> merged = new ArrayList<>();
        int[] current = intervals[0];
        merged.add(current);

        for (int i = 1; i < intervals.length; i++) {
            if (intervals[i][0] <= current[1]) {
                // Overlapping: extend the end
                current[1] = Math.max(current[1], intervals[i][1]);
            } else {
                // Non-overlapping: start a new interval
                current = intervals[i];
                merged.add(current);
            }
        }

        return merged.toArray(new int[merged.size()][]);
    }

    // ==================== Insert Interval ====================
    /**
     * Insert a new interval into a sorted, non-overlapping list and merge if needed.
     * Strategy: Add all intervals that end before newInterval starts,
     * merge all that overlap, then add the rest.
     * 
     * Time: O(n), Space: O(n)
     */
    public static int[][] insertInterval(int[][] intervals, int[] newInterval) {
        List<int[]> result = new ArrayList<>();
        int i = 0;
        int n = intervals.length;

        // Add all intervals that come before newInterval
        while (i < n && intervals[i][1] < newInterval[0]) {
            result.add(intervals[i]);
            i++;
        }

        // Merge overlapping intervals with newInterval
        while (i < n && intervals[i][0] <= newInterval[1]) {
            newInterval[0] = Math.min(newInterval[0], intervals[i][0]);
            newInterval[1] = Math.max(newInterval[1], intervals[i][1]);
            i++;
        }
        result.add(newInterval);

        // Add remaining intervals
        while (i < n) {
            result.add(intervals[i]);
            i++;
        }

        return result.toArray(new int[result.size()][]);
    }

    // ==================== Can Attend All Meetings ====================
    /**
     * Determine if a person can attend all meetings (no overlaps).
     * Strategy: Sort by start time, check if any meeting starts before the previous ends.
     * 
     * Time: O(n log n), Space: O(1)
     */
    public static boolean canAttendAllMeetings(int[][] intervals) {
        if (intervals.length <= 1) return true;

        Arrays.sort(intervals, (a, b) -> a[0] - b[0]);

        for (int i = 1; i < intervals.length; i++) {
            if (intervals[i][0] < intervals[i - 1][1]) {
                return false; // Overlap: current starts before previous ends
            }
        }

        return true;
    }

    // ==================== Minimum Meeting Rooms (Sweep Line) ====================
    /**
     * Find the minimum number of conference rooms needed.
     * Strategy: Use sweep line - create start (+1) and end (-1) events,
     * sort by time, track maximum concurrent meetings.
     * 
     * Time: O(n log n), Space: O(n)
     */
    public static int minMeetingRooms(int[][] intervals) {
        if (intervals.length == 0) return 0;

        // Create events: [time, type] where type = +1 for start, -1 for end
        int[][] events = new int[intervals.length * 2][2];
        int idx = 0;
        for (int[] interval : intervals) {
            events[idx++] = new int[]{interval[0], 1};  // Meeting starts
            events[idx++] = new int[]{interval[1], -1}; // Meeting ends
        }

        // Sort by time; if same time, end (-1) before start (+1)
        // This handles the case where one meeting ends exactly when another starts
        Arrays.sort(events, (a, b) -> a[0] != b[0] ? a[0] - b[0] : a[1] - b[1]);

        int maxRooms = 0;
        int currentRooms = 0;

        for (int[] event : events) {
            currentRooms += event[1];
            maxRooms = Math.max(maxRooms, currentRooms);
        }

        return maxRooms;
    }

    // ==================== Helper ====================
    private static boolean arraysEqual2D(int[][] a, int[][] b) {
        if (a.length != b.length) return false;
        for (int i = 0; i < a.length; i++) {
            if (!Arrays.equals(a[i], b[i])) return false;
        }
        return true;
    }

    // ==================== Tests ====================
    public static void main(String[] args) {
        // Test Merge Intervals
        assert arraysEqual2D(
            mergeIntervals(new int[][]{{1, 3}, {2, 6}, {8, 10}, {15, 18}}),
            new int[][]{{1, 6}, {8, 10}, {15, 18}}
        ) : "Merge intervals test 1 failed";
        assert arraysEqual2D(
            mergeIntervals(new int[][]{{1, 4}, {4, 5}}),
            new int[][]{{1, 5}}
        ) : "Merge intervals test 2 failed";
        assert arraysEqual2D(
            mergeIntervals(new int[][]{{1, 4}, {0, 4}}),
            new int[][]{{0, 4}}
        ) : "Merge intervals test 3 failed";

        // Test Insert Interval
        assert arraysEqual2D(
            insertInterval(new int[][]{{1, 3}, {6, 9}}, new int[]{2, 5}),
            new int[][]{{1, 5}, {6, 9}}
        ) : "Insert interval test 1 failed";
        assert arraysEqual2D(
            insertInterval(new int[][]{{1, 2}, {3, 5}, {6, 7}, {8, 10}, {12, 16}}, new int[]{4, 8}),
            new int[][]{{1, 2}, {3, 10}, {12, 16}}
        ) : "Insert interval test 2 failed";

        // Test Can Attend All Meetings
        assert !canAttendAllMeetings(new int[][]{{0, 30}, {5, 10}, {15, 20}})
            : "Meetings test 1 failed";
        assert canAttendAllMeetings(new int[][]{{7, 10}, {2, 4}})
            : "Meetings test 2 failed";
        assert canAttendAllMeetings(new int[][]{{1, 5}, {5, 10}})
            : "Meetings test 3 failed (adjacent, no overlap)";

        // Test Min Meeting Rooms (Sweep Line)
        assert minMeetingRooms(new int[][]{{0, 30}, {5, 10}, {15, 20}}) == 2
            : "Min rooms test 1 failed";
        assert minMeetingRooms(new int[][]{{7, 10}, {2, 4}}) == 1
            : "Min rooms test 2 failed";
        assert minMeetingRooms(new int[][]{{1, 5}, {2, 6}, {3, 7}, {4, 8}}) == 4
            : "Min rooms test 3 failed";
        assert minMeetingRooms(new int[][]{{1, 5}, {5, 10}}) == 1
            : "Min rooms test 4 failed (one ends as next starts)";

        System.out.println("All interval pattern tests passed!");
    }
}
