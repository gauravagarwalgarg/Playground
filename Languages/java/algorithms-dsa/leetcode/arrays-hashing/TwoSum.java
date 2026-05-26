import java.util.HashMap;
import java.util.Arrays;

/**
 * LeetCode 1: Two Sum
 * Given an array of integers nums and an integer target,
 * return indices of the two numbers such that they add up to target.
 * 
 * Time: O(n), Space: O(n)
 */
public class TwoSum {

    public static int[] twoSum(int[] nums, int target) {
        HashMap<Integer, Integer> seen = new HashMap<>();
        for (int i = 0; i < nums.length; i++) {
            int complement = target - nums[i];
            if (seen.containsKey(complement)) {
                return new int[]{seen.get(complement), i};
            }
            seen.put(nums[i], i);
        }
        return new int[]{};
    }

    public static void main(String[] args) {
        assert Arrays.equals(twoSum(new int[]{2, 7, 11, 15}, 9), new int[]{0, 1})
            : "Test 1 failed";
        assert Arrays.equals(twoSum(new int[]{3, 2, 4}, 6), new int[]{1, 2})
            : "Test 2 failed";
        assert Arrays.equals(twoSum(new int[]{3, 3}, 6), new int[]{0, 1})
            : "Test 3 failed";

        System.out.println("All tests passed!");
    }
}
