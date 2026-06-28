/**
 * Pattern: Prefix Sums
 *
 * Key techniques covered:
 * 1. Range Sum Query (prefix sum array)
 * 2. Subarray Sum Equals K (prefix sum + unordered_map)
 * 3. Product of Array Except Self (left/right product arrays)
 * 4. Find Pivot Index (total sum - left sum technique)
 * 5. Kadane's Maximum Subarray (running max technique)
 *
 * Core insight: prefix[i] = sum of elements from index 0 to i-1.
 * Range sum [l, r] = prefix[r+1] - prefix[l].
 *
 * Compile: g++ -std=c++17 -o test prefix_sums.cpp && ./test
 */
#include <iostream>
#include <vector>
#include <unordered_map>
#include <cassert>
#include <numeric>
using namespace std;

// ─────────────────────────────────────────────────────────────────────────────
// 1. Range Sum Query (Immutable)
//    Precompute prefix sums for O(1) range sum queries.
//    Time: O(n) build, O(1) query. Space: O(n)
// ─────────────────────────────────────────────────────────────────────────────
class NumArray {
    vector<int> prefix;

public:
    NumArray(vector<int>& nums) {
        int n = nums.size();
        prefix.resize(n + 1, 0);
        for (int i = 0; i < n; i++) {
            prefix[i + 1] = prefix[i] + nums[i];
        }
    }

    // Sum of elements in [left, right] inclusive
    int sumRange(int left, int right) {
        return prefix[right + 1] - prefix[left];
    }
};

// ─────────────────────────────────────────────────────────────────────────────
// 2. Subarray Sum Equals K
//    Count the number of contiguous subarrays that sum to K.
//    Strategy: If prefix[j] - prefix[i] == k, then subarray [i+1..j] sums to k.
//    Use a hashmap to count prefix sum occurrences.
//    Time: O(n), Space: O(n)
// ─────────────────────────────────────────────────────────────────────────────
int subarraySum(vector<int>& nums, int k) {
    unordered_map<int, int> prefixCount;
    prefixCount[0] = 1; // Empty prefix has sum 0
    int sum = 0, count = 0;

    for (int num : nums) {
        sum += num;
        // If (sum - k) was seen before, those are valid starting points
        if (prefixCount.count(sum - k)) {
            count += prefixCount[sum - k];
        }
        prefixCount[sum]++;
    }
    return count;
}

// ─────────────────────────────────────────────────────────────────────────────
// 3. Product of Array Except Self
//    For each index, compute product of all elements except that index.
//    Strategy: Two passes left products and right products.
//    Time: O(n), Space: O(1) extra (output array doesn't count)
// ─────────────────────────────────────────────────────────────────────────────
vector<int> productExceptSelf(vector<int>& nums) {
    int n = nums.size();
    vector<int> result(n, 1);

    // Left pass: result[i] = product of all elements to the left of i
    int leftProduct = 1;
    for (int i = 0; i < n; i++) {
        result[i] = leftProduct;
        leftProduct *= nums[i];
    }

    // Right pass: multiply by product of all elements to the right
    int rightProduct = 1;
    for (int i = n - 1; i >= 0; i--) {
        result[i] *= rightProduct;
        rightProduct *= nums[i];
    }
    return result;
}

// ─────────────────────────────────────────────────────────────────────────────
// 4. Find Pivot Index
//    Find the index where left sum == right sum.
//    Strategy: leftSum == totalSum - leftSum - nums[i]
//    Time: O(n), Space: O(1)
// ─────────────────────────────────────────────────────────────────────────────
int pivotIndex(vector<int>& nums) {
    int totalSum = accumulate(nums.begin(), nums.end(), 0);
    int leftSum = 0;

    for (int i = 0; i < (int)nums.size(); i++) {
        // rightSum = totalSum - leftSum - nums[i]
        if (leftSum == totalSum - leftSum - nums[i]) {
            return i;
        }
        leftSum += nums[i];
    }
    return -1;
}

// ─────────────────────────────────────────────────────────────────────────────
// 5. Kadane's Algorithm Maximum Subarray
//    Find the contiguous subarray with the largest sum.
//    Strategy: At each position, either extend the current subarray or start fresh.
//    maxEndingHere = max(nums[i], maxEndingHere + nums[i])
//    Time: O(n), Space: O(1)
// ─────────────────────────────────────────────────────────────────────────────
int maxSubArray(vector<int>& nums) {
    int maxSum = nums[0];
    int currentSum = nums[0];

    for (int i = 1; i < (int)nums.size(); i++) {
        // Either start a new subarray at i, or extend the previous one
        currentSum = max(nums[i], currentSum + nums[i]);
        maxSum = max(maxSum, currentSum);
    }
    return maxSum;
}

// ─────────────────────────────────────────────────────────────────────────────
int main() {
    // Test 1: Range Sum Query
    {
        vector<int> nums = {-2, 0, 3, -5, 2, -1};
        NumArray na(nums);
        assert(na.sumRange(0, 2) == 1);   // -2 + 0 + 3 = 1
        assert(na.sumRange(2, 5) == -1);  // 3 + (-5) + 2 + (-1) = -1
        assert(na.sumRange(0, 5) == -3);  // entire array
    }

    // Test 2: Subarray Sum Equals K
    {
        vector<int> nums = {1, 1, 1};
        assert(subarraySum(nums, 2) == 2);

        vector<int> nums2 = {1, 2, 3};
        assert(subarraySum(nums2, 3) == 2); // [1,2] and [3]

        vector<int> nums3 = {1, -1, 0};
        assert(subarraySum(nums3, 0) == 3); // [1,-1], [-1,0], [1,-1,0]
    }

    // Test 3: Product Except Self
    {
        vector<int> nums = {1, 2, 3, 4};
        auto result = productExceptSelf(nums);
        assert(result == vector<int>({24, 12, 8, 6}));

        vector<int> nums2 = {-1, 1, 0, -3, 3};
        auto result2 = productExceptSelf(nums2);
        assert(result2 == vector<int>({0, 0, 9, 0, 0}));
    }

    // Test 4: Pivot Index
    {
        vector<int> nums = {1, 7, 3, 6, 5, 6};
        assert(pivotIndex(nums) == 3); // left: 1+7+3=11, right: 5+6=11

        vector<int> nums2 = {1, 2, 3};
        assert(pivotIndex(nums2) == -1);

        vector<int> nums3 = {2, 1, -1};
        assert(pivotIndex(nums3) == 0); // left: 0, right: 1+(-1)=0
    }

    // Test 5: Kadane's Maximum Subarray
    {
        vector<int> nums = {-2, 1, -3, 4, -1, 2, 1, -5, 4};
        assert(maxSubArray(nums) == 6); // [4,-1,2,1]

        vector<int> nums2 = {1};
        assert(maxSubArray(nums2) == 1);

        vector<int> nums3 = {-1, -2, -3};
        assert(maxSubArray(nums3) == -1); // Single element
    }

    cout << "All prefix sum pattern tests passed!" << endl;
    return 0;
}
