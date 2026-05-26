/**
 * LeetCode 217: Contains Duplicate
 * Given an integer array nums, return 1 if any value appears at least twice.
 *
 * Time: O(n^2) - simple approach without hash set
 * Space: O(1)
 */

#include <stdio.h>
#include <assert.h>

int contains_duplicate(int* nums, int nums_size) {
    for (int i = 0; i < nums_size; i++) {
        for (int j = i + 1; j < nums_size; j++) {
            if (nums[i] == nums[j]) {
                return 1;
            }
        }
    }
    return 0;
}

int main(void) {
    int nums1[] = {1, 2, 3, 1};
    assert(contains_duplicate(nums1, 4) == 1);

    int nums2[] = {1, 2, 3, 4};
    assert(contains_duplicate(nums2, 4) == 0);

    int nums3[] = {1, 1, 1, 3, 3, 4, 3, 2, 4, 2};
    assert(contains_duplicate(nums3, 10) == 1);

    // Empty array
    assert(contains_duplicate(NULL, 0) == 0);

    printf("All tests passed!\n");
    return 0;
}
