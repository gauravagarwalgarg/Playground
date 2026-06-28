import java.util.*;

/**
 * Heap / Priority Queue Patterns
 * 
 * Key insight: Heaps give O(log n) insert and O(1) access to min/max element.
 * Use a min-heap when you need the smallest, max-heap when you need the largest.
 * 
 * Common patterns:
 * 1. Top-K elements: Use a min-heap of size K (keeps K largest)
 * 2. Merge K sorted: Use a min-heap to always pick the smallest next element
 * 3. Kth largest: Maintain a min-heap of size K; root is the answer
 * 4. Running median: Use two heaps (max-heap for lower half, min-heap for upper half)
 * 
 * Time complexities:
 * - Top-K: O(n log k)
 * - Merge K sorted: O(n log k) where n = total elements, k = number of lists
 * - Running median: O(log n) per insert, O(1) per findMedian
 */
public class Heaps {

    // ==================== Top-K Frequent Elements ====================
    /**
     * Find the k most frequent elements in an array.
     * Strategy: Count frequencies, then use a min-heap of size k.
     * The heap evicts the least frequent, leaving the k most frequent.
     * 
     * Time: O(n log k), Space: O(n)
     */
    public static int[] topKFrequent(int[] nums, int k) {
        // Step 1: Count frequencies
        Map<Integer, Integer> freq = new HashMap<>();
        for (int num : nums) {
            freq.put(num, freq.getOrDefault(num, 0) + 1);
        }

        // Step 2: Min-heap by frequency, keep size k
        PriorityQueue<Integer> minHeap = new PriorityQueue<>(
            (a, b) -> freq.get(a) - freq.get(b)
        );

        for (int num : freq.keySet()) {
            minHeap.offer(num);
            if (minHeap.size() > k) {
                minHeap.poll(); // Remove least frequent
            }
        }

        // Step 3: Extract results
        int[] result = new int[k];
        for (int i = 0; i < k; i++) {
            result[i] = minHeap.poll();
        }
        return result;
    }

    // ==================== Merge K Sorted Lists ====================
    /**
     * Merge k sorted arrays into one sorted array.
     * Strategy: Use a min-heap holding one element from each list.
     * Always extract the minimum and advance that list's pointer.
     * 
     * Time: O(n log k), Space: O(k) for the heap
     */
    public static int[] mergeKSorted(int[][] lists) {
        // Min-heap stores: [value, listIndex, elementIndex]
        PriorityQueue<int[]> minHeap = new PriorityQueue<>(
            (a, b) -> a[0] - b[0]
        );

        // Initialize: add first element of each list
        int totalSize = 0;
        for (int i = 0; i < lists.length; i++) {
            if (lists[i].length > 0) {
                minHeap.offer(new int[]{lists[i][0], i, 0});
                totalSize += lists[i].length;
            }
        }

        int[] result = new int[totalSize];
        int idx = 0;

        while (!minHeap.isEmpty()) {
            int[] curr = minHeap.poll();
            result[idx++] = curr[0];

            int listIdx = curr[1];
            int elemIdx = curr[2] + 1;

            // If there's a next element in that list, add it
            if (elemIdx < lists[listIdx].length) {
                minHeap.offer(new int[]{lists[listIdx][elemIdx], listIdx, elemIdx});
            }
        }

        return result;
    }

    // ==================== Kth Largest Element ====================
    /**
     * Find the kth largest element in an unsorted array.
     * Strategy: Maintain a min-heap of size k.
     * After processing all elements, the root is the kth largest.
     * 
     * Time: O(n log k), Space: O(k)
     */
    public static int kthLargest(int[] nums, int k) {
        PriorityQueue<Integer> minHeap = new PriorityQueue<>();

        for (int num : nums) {
            minHeap.offer(num);
            if (minHeap.size() > k) {
                minHeap.poll(); // Remove smallest, keep k largest
            }
        }

        return minHeap.peek(); // Root = kth largest
    }

    // ==================== Running Median (MedianFinder) ====================
    /**
     * MedianFinder maintains a running median using two heaps:
     * - maxHeap: stores the smaller half (gives quick access to max of lower half)
     * - minHeap: stores the larger half (gives quick access to min of upper half)
     * 
     * Invariant: maxHeap.size() == minHeap.size() or maxHeap.size() == minHeap.size() + 1
     * Median = maxHeap.peek() if odd count, else average of both peeks.
     */
    static class MedianFinder {
        private PriorityQueue<Integer> maxHeap; // Lower half
        private PriorityQueue<Integer> minHeap; // Upper half

        public MedianFinder() {
            maxHeap = new PriorityQueue<>(Collections.reverseOrder());
            minHeap = new PriorityQueue<>();
        }

        public void addNum(int num) {
            // Always add to maxHeap first
            maxHeap.offer(num);

            // Balance: ensure maxHeap's max <= minHeap's min
            if (!minHeap.isEmpty() && maxHeap.peek() > minHeap.peek()) {
                minHeap.offer(maxHeap.poll());
            }

            // Balance sizes: maxHeap can have at most 1 more than minHeap
            if (maxHeap.size() > minHeap.size() + 1) {
                minHeap.offer(maxHeap.poll());
            } else if (minHeap.size() > maxHeap.size()) {
                maxHeap.offer(minHeap.poll());
            }
        }

        public double findMedian() {
            if (maxHeap.size() > minHeap.size()) {
                return maxHeap.peek();
            }
            return (maxHeap.peek() + minHeap.peek()) / 2.0;
        }
    }

    // ==================== Tests ====================
    public static void main(String[] args) {
        // Test Top-K Frequent
        int[] topK = topKFrequent(new int[]{1, 1, 1, 2, 2, 3}, 2);
        Arrays.sort(topK);
        assert Arrays.equals(topK, new int[]{1, 2}) : "Top-K test 1 failed";

        int[] topK2 = topKFrequent(new int[]{1}, 1);
        assert topK2[0] == 1 : "Top-K test 2 failed";

        // Test Merge K Sorted
        int[][] lists = {{1, 4, 7}, {2, 5, 8}, {3, 6, 9}};
        int[] merged = mergeKSorted(lists);
        assert Arrays.equals(merged, new int[]{1, 2, 3, 4, 5, 6, 7, 8, 9})
            : "Merge K sorted test failed";

        int[][] lists2 = {{1, 3, 5}, {2, 4, 6}, {0, 7}};
        int[] merged2 = mergeKSorted(lists2);
        assert Arrays.equals(merged2, new int[]{0, 1, 2, 3, 4, 5, 6, 7})
            : "Merge K sorted test 2 failed";

        // Test Kth Largest
        assert kthLargest(new int[]{3, 2, 1, 5, 6, 4}, 2) == 5
            : "Kth largest test 1 failed";
        assert kthLargest(new int[]{3, 2, 3, 1, 2, 4, 5, 5, 6}, 4) == 4
            : "Kth largest test 2 failed";

        // Test MedianFinder
        MedianFinder mf = new MedianFinder();
        mf.addNum(1);
        assert mf.findMedian() == 1.0 : "Median test 1 failed";
        mf.addNum(2);
        assert mf.findMedian() == 1.5 : "Median test 2 failed";
        mf.addNum(3);
        assert mf.findMedian() == 2.0 : "Median test 3 failed";
        mf.addNum(4);
        assert mf.findMedian() == 2.5 : "Median test 4 failed";
        mf.addNum(5);
        assert mf.findMedian() == 3.0 : "Median test 5 failed";

        System.out.println("All heap pattern tests passed!");
    }
}
