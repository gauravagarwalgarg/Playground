import java.util.*;

/**
 * Prefix Sum Patterns
 * 
 * Key insight: Precompute cumulative sums to answer range queries in O(1).
 * prefix[i] = sum of elements from index 0 to i-1
 * sum(i, j) = prefix[j+1] - prefix[i]
 * 
 * Common patterns:
 * 1. Range Sum Query: Precompute prefix sums for O(1) range queries
 * 2. Subarray Sum Equals K: prefix sum + hashmap for count
 * 3. Product Except Self: Left and right running products
 * 4. Pivot Index: Total sum - left sum = left sum (find equilibrium)
 * 5. Kadane's Algorithm: Maximum subarray sum via running sum with reset
 */
public class PrefixSums {

    // ==================== Range Sum Query ====================
    /**
     * Precompute prefix sums to answer range sum queries in O(1).
     * prefix[i] = nums[0] + nums[1] + ... + nums[i-1]
     * sum(left, right) = prefix[right+1] - prefix[left]
     * 
     * Build: O(n), Query: O(1), Space: O(n)
     */
    static int[] buildPrefixSum(int[] nums) {
        int[] prefix = new int[nums.length + 1];
        for (int i = 0; i < nums.length; i++) {
            prefix[i + 1] = prefix[i] + nums[i];
        }
        return prefix;
    }

    static int rangeSum(int[] prefix, int left, int right) {
        return prefix[right + 1] - prefix[left];
    }

    // ==================== Subarray Sum Equals K ====================
    /**
     * Count the number of contiguous subarrays that sum to k.
     * Strategy: If prefix[j] - prefix[i] == k, then subarray (i, j] sums to k.
     * Use a hashmap to count how many previous prefix sums equal (currentSum - k).
     * 
     * Time: O(n), Space: O(n)
     */
    public static int subarraySum(int[] nums, int k) {
        Map<Integer, Integer> prefixCount = new HashMap<>();
        prefixCount.put(0, 1); // Empty prefix has sum 0

        int currentSum = 0;
        int count = 0;

        for (int num : nums) {
            currentSum += num;
            // How many previous prefix sums equal (currentSum - k)?
            count += prefixCount.getOrDefault(currentSum - k, 0);
            prefixCount.put(currentSum, prefixCount.getOrDefault(currentSum, 0) + 1);
        }

        return count;
    }

    // ==================== Product of Array Except Self ====================
    /**
     * For each index, compute the product of all other elements (no division).
     * Strategy: Two passes - left products and right products.
     * result[i] = product of all elements to the left * product of all to the right
     * 
     * Time: O(n), Space: O(1) extra (output array not counted)
     */
    public static int[] productExceptSelf(int[] nums) {
        int n = nums.length;
        int[] result = new int[n];

        // Left pass: result[i] = product of nums[0..i-1]
        result[0] = 1;
        for (int i = 1; i < n; i++) {
            result[i] = result[i - 1] * nums[i - 1];
        }

        // Right pass: multiply by product of nums[i+1..n-1]
        int rightProduct = 1;
        for (int i = n - 2; i >= 0; i--) {
            rightProduct *= nums[i + 1];
            result[i] *= rightProduct;
        }

        return result;
    }

    // ==================== Find Pivot Index ====================
    /**
     * Find the index where the sum of elements to the left equals the sum to the right.
     * Strategy: leftSum == totalSum - leftSum - nums[i]
     * Simplifies to: 2 * leftSum + nums[i] == totalSum
     * 
     * Time: O(n), Space: O(1)
     */
    public static int pivotIndex(int[] nums) {
        int totalSum = 0;
        for (int num : nums) totalSum += num;

        int leftSum = 0;
        for (int i = 0; i < nums.length; i++) {
            // Right sum = totalSum - leftSum - nums[i]
            if (leftSum == totalSum - leftSum - nums[i]) {
                return i;
            }
            leftSum += nums[i];
        }

        return -1; // No pivot index found
    }

    // ==================== Kadane's Algorithm (Max Subarray) ====================
    /**
     * Find the contiguous subarray with the maximum sum.
     * Strategy: Maintain a running sum. If it drops below 0, reset to 0
     * (starting a new subarray is better than carrying negative sum).
     * Track the maximum seen at any point.
     * 
     * Time: O(n), Space: O(1)
     */
    public static int maxSubarraySum(int[] nums) {
        int maxSum = nums[0];
        int currentSum = nums[0];

        for (int i = 1; i < nums.length; i++) {
            // Either extend the current subarray or start fresh
            currentSum = Math.max(nums[i], currentSum + nums[i]);
            maxSum = Math.max(maxSum, currentSum);
        }

        return maxSum;
    }

    // ==================== Tests ====================
    public static void main(String[] args) {
        // Test Range Sum Query
        int[] nums = {1, 2, 3, 4, 5};
        int[] prefix = buildPrefixSum(nums);
        assert rangeSum(prefix, 0, 4) == 15 : "Range sum test 1 failed (full array)";
        assert rangeSum(prefix, 1, 3) == 9 : "Range sum test 2 failed (2+3+4)";
        assert rangeSum(prefix, 2, 2) == 3 : "Range sum test 3 failed (single element)";
        assert rangeSum(prefix, 0, 0) == 1 : "Range sum test 4 failed";

        // Test Subarray Sum Equals K
        assert subarraySum(new int[]{1, 1, 1}, 2) == 2
            : "Subarray sum test 1 failed";
        assert subarraySum(new int[]{1, 2, 3}, 3) == 2
            : "Subarray sum test 2 failed (both [1,2] and [3])";
        assert subarraySum(new int[]{1, -1, 0}, 0) == 3
            : "Subarray sum test 3 failed (negative numbers)";
        assert subarraySum(new int[]{3, 4, 7, 2, -3, 1, 4, 2}, 7) == 4
            : "Subarray sum test 4 failed";

        // Test Product Except Self
        assert Arrays.equals(
            productExceptSelf(new int[]{1, 2, 3, 4}),
            new int[]{24, 12, 8, 6}
        ) : "Product except self test 1 failed";
        assert Arrays.equals(
            productExceptSelf(new int[]{-1, 1, 0, -3, 3}),
            new int[]{0, 0, 9, 0, 0}
        ) : "Product except self test 2 failed";

        // Test Pivot Index
        assert pivotIndex(new int[]{1, 7, 3, 6, 5, 6}) == 3
            : "Pivot index test 1 failed";
        assert pivotIndex(new int[]{1, 2, 3}) == -1
            : "Pivot index test 2 failed (no pivot)";
        assert pivotIndex(new int[]{2, 1, -1}) == 0
            : "Pivot index test 3 failed (pivot at start)";

        // Test Kadane's Maximum Subarray
        assert maxSubarraySum(new int[]{-2, 1, -3, 4, -1, 2, 1, -5, 4}) == 6
            : "Kadane test 1 failed (subarray [4,-1,2,1])";
        assert maxSubarraySum(new int[]{1}) == 1
            : "Kadane test 2 failed (single element)";
        assert maxSubarraySum(new int[]{5, 4, -1, 7, 8}) == 23
            : "Kadane test 3 failed (entire array)";
        assert maxSubarraySum(new int[]{-1, -2, -3}) == -1
            : "Kadane test 4 failed (all negative)";

        System.out.println("All prefix sum pattern tests passed!");
    }
}
