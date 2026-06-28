import java.util.*;

/**
 * Bit Manipulation Patterns
 * 
 * Essential bit operations:
 * - x & 1       : Check if x is odd
 * - x & (x-1)  : Turn off the rightmost set bit
 * - x ^ x = 0  : XOR cancels duplicates
 * - x | (1<<i) : Set bit i
 * - x & (1<<i) : Check bit i
 * - x ^ (1<<i) : Toggle bit i
 * 
 * Common patterns:
 * 1. Single Number: XOR all elements, duplicates cancel
 * 2. Power of Two: n & (n-1) == 0 (only one bit set)
 * 3. Count Bits: Brian Kernighan's trick (n &= n-1 removes lowest set bit)
 * 4. Hamming Distance: XOR + count set bits
 * 5. Subsets via Bitmask: Enumerate all 2^n subsets
 * 6. Missing Number: XOR with indices cancels all except missing
 */
public class BitManipulation {

    // ==================== Single Number ====================
    /**
     * Find the element that appears only once (all others appear twice).
     * Strategy: XOR all elements. Since a ^ a = 0 and a ^ 0 = a,
     * all duplicates cancel out, leaving the unique element.
     * 
     * Time: O(n), Space: O(1)
     */
    public static int singleNumber(int[] nums) {
        int result = 0;
        for (int num : nums) {
            result ^= num;
        }
        return result;
    }

    // ==================== Power of Two ====================
    /**
     * Check if n is a power of two.
     * Strategy: A power of two has exactly one bit set (e.g., 8 = 1000).
     * n & (n-1) turns off the lowest set bit.
     * If the result is 0, there was only one bit set.
     * 
     * Time: O(1), Space: O(1)
     */
    public static boolean isPowerOfTwo(int n) {
        return n > 0 && (n & (n - 1)) == 0;
    }

    // ==================== Count Set Bits (Hamming Weight) ====================
    /**
     * Count the number of 1-bits in an integer.
     * Strategy: Brian Kernighan's algorithm - n & (n-1) removes the lowest
     * set bit. Count how many times we can do this.
     * 
     * Time: O(number of set bits), Space: O(1)
     */
    public static int countBits(int n) {
        int count = 0;
        while (n != 0) {
            n &= (n - 1); // Remove lowest set bit
            count++;
        }
        return count;
    }

    /**
     * Count bits for every number from 0 to n.
     * Uses dynamic programming: bits[i] = bits[i >> 1] + (i & 1)
     * The number of bits in i equals bits in i/2 plus the last bit.
     * 
     * Time: O(n), Space: O(n)
     */
    public static int[] countBitsRange(int n) {
        int[] bits = new int[n + 1];
        for (int i = 1; i <= n; i++) {
            bits[i] = bits[i >> 1] + (i & 1);
        }
        return bits;
    }

    // ==================== Hamming Distance ====================
    /**
     * Count the number of positions where corresponding bits differ.
     * Strategy: XOR gives 1 where bits differ, then count set bits.
     * 
     * Time: O(1), Space: O(1)
     */
    public static int hammingDistance(int x, int y) {
        return countBits(x ^ y);
    }

    // ==================== Subsets via Bitmask ====================
    /**
     * Generate all subsets of an array using bitmask enumeration.
     * Strategy: For n elements, there are 2^n subsets.
     * Each number from 0 to 2^n - 1 represents a subset where
     * bit i indicates whether element i is included.
     * 
     * Time: O(n * 2^n), Space: O(n * 2^n)
     */
    public static List<List<Integer>> subsets(int[] nums) {
        int n = nums.length;
        int totalSubsets = 1 << n; // 2^n
        List<List<Integer>> result = new ArrayList<>();

        for (int mask = 0; mask < totalSubsets; mask++) {
            List<Integer> subset = new ArrayList<>();
            for (int i = 0; i < n; i++) {
                if ((mask & (1 << i)) != 0) {
                    subset.add(nums[i]);
                }
            }
            result.add(subset);
        }

        return result;
    }

    // ==================== Missing Number ====================
    /**
     * Find the missing number in [0, n] given n numbers.
     * Strategy: XOR all numbers with all indices 0..n.
     * Every number except the missing one appears in both sets,
     * so they cancel out, leaving only the missing number.
     * 
     * Time: O(n), Space: O(1)
     */
    public static int missingNumber(int[] nums) {
        int n = nums.length;
        int xor = n; // Start with n (since indices go 0..n-1)
        for (int i = 0; i < n; i++) {
            xor ^= i ^ nums[i];
        }
        return xor;
    }

    // ==================== Tests ====================
    public static void main(String[] args) {
        // Test Single Number
        assert singleNumber(new int[]{2, 2, 1}) == 1 : "Single number test 1 failed";
        assert singleNumber(new int[]{4, 1, 2, 1, 2}) == 4 : "Single number test 2 failed";
        assert singleNumber(new int[]{1}) == 1 : "Single number test 3 failed";

        // Test Power of Two
        assert isPowerOfTwo(1) : "Power of 2 test: 1 failed";
        assert isPowerOfTwo(2) : "Power of 2 test: 2 failed";
        assert isPowerOfTwo(16) : "Power of 2 test: 16 failed";
        assert !isPowerOfTwo(3) : "Power of 2 test: 3 failed";
        assert !isPowerOfTwo(0) : "Power of 2 test: 0 failed";
        assert !isPowerOfTwo(-4) : "Power of 2 test: -4 failed";

        // Test Count Bits
        assert countBits(0) == 0 : "Count bits test: 0 failed";
        assert countBits(1) == 1 : "Count bits test: 1 failed";
        assert countBits(7) == 3 : "Count bits test: 7 (111) failed";
        assert countBits(255) == 8 : "Count bits test: 255 failed";

        // Test Count Bits Range
        assert Arrays.equals(countBitsRange(5), new int[]{0, 1, 1, 2, 1, 2})
            : "Count bits range test failed";

        // Test Hamming Distance
        assert hammingDistance(1, 4) == 2 : "Hamming test 1 failed (001 vs 100)";
        assert hammingDistance(3, 1) == 1 : "Hamming test 2 failed (11 vs 01)";
        assert hammingDistance(0, 0) == 0 : "Hamming test 3 failed";

        // Test Subsets
        List<List<Integer>> subs = subsets(new int[]{1, 2, 3});
        assert subs.size() == 8 : "Subsets test: expected 8, got " + subs.size();
        // Verify empty set and full set are present
        assert subs.contains(new ArrayList<>()) : "Subsets should contain empty set";
        assert subs.contains(Arrays.asList(1, 2, 3)) : "Subsets should contain full set";

        List<List<Integer>> subs2 = subsets(new int[]{0});
        assert subs2.size() == 2 : "Subsets test 2 failed";

        // Test Missing Number
        assert missingNumber(new int[]{3, 0, 1}) == 2 : "Missing number test 1 failed";
        assert missingNumber(new int[]{0, 1}) == 2 : "Missing number test 2 failed";
        assert missingNumber(new int[]{9, 6, 4, 2, 3, 5, 7, 0, 1}) == 8
            : "Missing number test 3 failed";
        assert missingNumber(new int[]{0}) == 1 : "Missing number test 4 failed";

        System.out.println("All bit manipulation pattern tests passed!");
    }
}
