/**
 * LeetCode 1: Two Sum
 * Given an array of integers nums and an integer target,
 * return indices of the two numbers such that they add up to target.
 *
 * Time: O(n^2) - simple approach without hash map
 * Space: O(1)
 */

#include <stdio.h>
#include <assert.h>
#include <stdlib.h>

/**
 * Returns a malloc'd array of 2 indices, or NULL if not found.
 * Caller must free the result.
 */
int* two_sum(int* nums, int nums_size, int target, int* return_size) {
    for (int i = 0; i < nums_size; i++) {
        for (int j = i + 1; j < nums_size; j++) {
            if (nums[i] + nums[j] == target) {
                int* result = (int*)malloc(2 * sizeof(int));
                result[0] = i;
                result[1] = j;
                *return_size = 2;
                return result;
            }
        }
    }
    *return_size = 0;
    return NULL;
}

int main(void) {
    int return_size;

    // Test 1
    int nums1[] = {2, 7, 11, 15};
    int* result1 = two_sum(nums1, 4, 9, &return_size);
    assert(return_size == 2);
    assert(result1[0] == 0 && result1[1] == 1);
    free(result1);

    // Test 2
    int nums2[] = {3, 2, 4};
    int* result2 = two_sum(nums2, 3, 6, &return_size);
    assert(return_size == 2);
    assert(result2[0] == 1 && result2[1] == 2);
    free(result2);

    // Test 3
    int nums3[] = {3, 3};
    int* result3 = two_sum(nums3, 2, 6, &return_size);
    assert(return_size == 2);
    assert(result3[0] == 0 && result3[1] == 1);
    free(result3);

    // Test 4: no solution
    int nums4[] = {1, 2, 3};
    int* result4 = two_sum(nums4, 3, 10, &return_size);
    assert(return_size == 0);
    assert(result4 == NULL);

    printf("All tests passed!\n");
    return 0;
}
