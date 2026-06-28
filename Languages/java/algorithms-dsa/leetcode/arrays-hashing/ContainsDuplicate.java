import java.util.HashSet;

/**
 * LeetCode 217: Contains Duplicate
 * Given an integer array nums, return true if any value appears at least twice.
 * 
 * Time: O(n), Space: O(n)
 */
public class ContainsDuplicate {

    public static boolean containsDuplicate(int[] nums) {
        HashSet<Integer> seen = new HashSet<>();
        for (int num : nums) {
            if (!seen.add(num)) {
                return true;
            }
        }
        return false;
    }

    public static void main(String[] args) {
        assert containsDuplicate(new int[]{1, 2, 3, 1}) == true : "Test 1 failed";
        assert containsDuplicate(new int[]{1, 2, 3, 4}) == false : "Test 2 failed";
        assert containsDuplicate(new int[]{1, 1, 1, 3, 3, 4, 3, 2, 4, 2}) == true : "Test 3 failed";
        assert containsDuplicate(new int[]{}) == false : "Test 4 failed";

        System.out.println("All tests passed!");
    }
}
