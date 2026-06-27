"""
Dynamic Programming Pattern Templates

Four foundational DP patterns:
1. Fibonacci-style (1D linear)
2. 0/1 Knapsack (take or skip with capacity)
3. LCS (two-sequence 2D)
4. LIS (longest increasing subsequence)

Each shows recurrence, base case, and space optimization notes.

Run: python dynamic_programming.py
"""

from typing import List
import bisect


# =============================================================================
# TEMPLATE 1: Fibonacci-Style (1D Linear DP)
# Pattern: dp[i] depends on dp[i-1], dp[i-2] (constant lookback)
# Example: House Robber max money robbing non-adjacent houses
# Time: O(n), Space: O(1) with optimization
# =============================================================================
def house_robber(nums: List[int]) -> int:
    """Max sum of non-adjacent elements (can't rob two adjacent houses)."""
    if not nums:
        return 0
    if len(nums) <= 2:
        return max(nums)

    # dp[i] = max money robbing houses [0..i]
    # dp[i] = max(dp[i-1], dp[i-2] + nums[i])
    #   - skip house i: take dp[i-1]
    #   - rob house i: take dp[i-2] + nums[i]

    # Space-optimized: only need two previous values
    prev2 = nums[0]            # dp[i-2]
    prev1 = max(nums[0], nums[1])  # dp[i-1]

    for i in range(2, len(nums)):
        curr = max(prev1, prev2 + nums[i])
        prev2 = prev1
        prev1 = curr

    return prev1


def climbing_stairs(n: int) -> int:
    """Count ways to climb n stairs taking 1 or 2 steps at a time."""
    if n <= 2:
        return n
    a, b = 1, 2
    for _ in range(3, n + 1):
        a, b = b, a + b
    return b


# =============================================================================
# TEMPLATE 2: 0/1 Knapsack
# Pattern: For each item, TAKE or SKIP. Track remaining capacity.
# Example: Maximize value within weight capacity
# Time: O(n * W), Space: O(W) with optimization
# =============================================================================
def knapsack_01(weights: List[int], values: List[int], capacity: int) -> int:
    """Classic 0/1 knapsack: max value within given capacity."""
    # dp[w] = max value achievable with capacity w
    dp = [0] * (capacity + 1)

    for i in range(len(weights)):
        # CRITICAL: iterate in REVERSE to avoid using same item twice
        # (Forward iteration gives unbounded knapsack)
        for w in range(capacity, weights[i] - 1, -1):
            dp[w] = max(dp[w],                          # skip item i
                        dp[w - weights[i]] + values[i])  # take item i

    return dp[capacity]


def subset_sum(nums: List[int], target: int) -> bool:
    """Can we select a subset that sums to exactly target? (Knapsack variant)"""
    dp = [False] * (target + 1)
    dp[0] = True  # empty subset sums to 0

    for num in nums:
        for t in range(target, num - 1, -1):  # reverse!
            dp[t] = dp[t] or dp[t - num]

    return dp[target]


def partition_equal_subset(nums: List[int]) -> bool:
    """Can array be partitioned into two subsets with equal sum?"""
    total = sum(nums)
    if total % 2 != 0:
        return False
    return subset_sum(nums, total // 2)


# =============================================================================
# TEMPLATE 3: Longest Common Subsequence (2D Two-Sequence DP)
# Pattern: dp[i][j] = answer for first i of seq1 and first j of seq2
# Example: Find LCS of two strings
# Time: O(m * n), Space: O(min(m, n)) with optimization
# =============================================================================
def lcs(text1: str, text2: str) -> int:
    """Length of longest common subsequence of two strings."""
    m, n = len(text1), len(text2)

    # dp[i][j] = LCS of text1[0..i-1] and text2[0..j-1]
    # Space optimization: only need previous row
    prev = [0] * (n + 1)

    for i in range(1, m + 1):
        curr = [0] * (n + 1)
        for j in range(1, n + 1):
            if text1[i - 1] == text2[j - 1]:
                curr[j] = prev[j - 1] + 1    # chars match → extend
            else:
                curr[j] = max(prev[j],        # skip char from text1
                              curr[j - 1])    # skip char from text2
        prev = curr

    return prev[n]


def edit_distance(word1: str, word2: str) -> int:
    """Minimum operations (insert/delete/replace) to convert word1 → word2."""
    m, n = len(word1), len(word2)
    # dp[i][j] = edit distance of word1[:i] and word2[:j]
    prev = list(range(n + 1))  # base: converting "" to word2[:j] costs j

    for i in range(1, m + 1):
        curr = [i] + [0] * n  # converting word1[:i] to "" costs i
        for j in range(1, n + 1):
            if word1[i - 1] == word2[j - 1]:
                curr[j] = prev[j - 1]  # no operation needed
            else:
                curr[j] = 1 + min(prev[j],      # delete from word1
                                  curr[j - 1],   # insert into word1
                                  prev[j - 1])   # replace
        prev = curr

    return prev[n]


# =============================================================================
# TEMPLATE 4: Longest Increasing Subsequence (LIS)
# Two approaches: O(n²) DP and O(n log n) patience sort
# Time: O(n²) or O(n log n), Space: O(n)
# =============================================================================
def lis_dp(nums: List[int]) -> int:
    """LIS using O(n²) DP clearer, useful for reconstruction."""
    if not nums:
        return 0

    # dp[i] = length of LIS ending at index i
    dp = [1] * len(nums)  # every element alone is LIS of length 1

    for i in range(1, len(nums)):
        for j in range(i):
            if nums[j] < nums[i]:
                dp[i] = max(dp[i], dp[j] + 1)

    return max(dp)


def lis_binary_search(nums: List[int]) -> int:
    """LIS using O(n log n) patience sorting optimal for length only."""
    # tails[i] = smallest tail of all increasing subsequences of length i+1
    tails = []

    for x in nums:
        pos = bisect.bisect_left(tails, x)  # first tail >= x
        if pos == len(tails):
            tails.append(x)    # x extends longest subsequence
        else:
            tails[pos] = x     # x improves subsequence of this length

    return len(tails)


# =============================================================================
if __name__ == "__main__":
    # Fibonacci-style
    print("House Robber [2,7,9,3,1]:", house_robber([2, 7, 9, 3, 1]))
    print("Climb 5 stairs:", climbing_stairs(5))

    # 0/1 Knapsack
    print("Knapsack w=[2,3,4,5] v=[3,4,5,6] cap=8:", knapsack_01([2, 3, 4, 5], [3, 4, 5, 6], 8))
    print("Subset sum [3,34,4,12,5,2] target=9:", subset_sum([3, 34, 4, 12, 5, 2], 9))
    print("Partition [1,5,11,5]:", partition_equal_subset([1, 5, 11, 5]))

    # LCS
    print("LCS 'abcde' 'ace':", lcs("abcde", "ace"))
    print("Edit distance 'horse' 'ros':", edit_distance("horse", "ros"))

    # LIS
    seq = [10, 9, 2, 5, 3, 7, 101, 18]
    print(f"LIS {seq} (O(n²)): {lis_dp(seq)}")
    print(f"LIS {seq} (O(nlogn)): {lis_binary_search(seq)}")
