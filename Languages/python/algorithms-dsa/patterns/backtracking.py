"""
Backtracking Pattern Templates

Four fundamental backtracking patterns:
1. Subsets generate all subsets (power set)
2. Permutations generate all orderings
3. Combinations choose k from n
4. Constraint Backtrack place elements satisfying constraints (N-Queens style)

Run: python backtracking.py
"""

from typing import List


# =============================================================================
# TEMPLATE 1: Subsets (Power Set)
# Decision at each element: include or skip.
# Every recursive state is a valid subset.
# Time: O(2^n), Space: O(n) recursion depth
# =============================================================================
def subsets(nums: List[int]) -> List[List[int]]:
    """Generate all subsets of nums."""
    result = []

    def backtrack(start: int, current: List[int]):
        result.append(current[:])  # record current state (copy)

        for i in range(start, len(nums)):
            current.append(nums[i])       # CHOOSE
            backtrack(i + 1, current)      # EXPLORE (i+1: no reuse)
            current.pop()                  # UNCHOOSE

    backtrack(0, [])
    return result


def subsets_with_dup(nums: List[int]) -> List[List[int]]:
    """Generate unique subsets when input has duplicates."""
    nums.sort()  # MUST sort to group duplicates
    result = []

    def backtrack(start: int, current: List[int]):
        result.append(current[:])

        for i in range(start, len(nums)):
            # Skip duplicate elements at the same level
            if i > start and nums[i] == nums[i - 1]:
                continue
            current.append(nums[i])
            backtrack(i + 1, current)
            current.pop()

    backtrack(0, [])
    return result


# =============================================================================
# TEMPLATE 2: Permutations
# Use each element exactly once. Track used elements.
# Complete when current has all elements.
# Time: O(n!), Space: O(n)
# =============================================================================
def permutations(nums: List[int]) -> List[List[int]]:
    """Generate all permutations of nums."""
    result = []

    def backtrack(current: List[int], used: List[bool]):
        if len(current) == len(nums):
            result.append(current[:])
            return

        for i in range(len(nums)):
            if used[i]:
                continue
            used[i] = True
            current.append(nums[i])
            backtrack(current, used)
            current.pop()
            used[i] = False

    backtrack([], [False] * len(nums))
    return result


def permutations_unique(nums: List[int]) -> List[List[int]]:
    """Generate unique permutations when input has duplicates."""
    nums.sort()
    result = []

    def backtrack(current: List[int], used: List[bool]):
        if len(current) == len(nums):
            result.append(current[:])
            return

        for i in range(len(nums)):
            if used[i]:
                continue
            # Skip duplicate: same value as previous AND previous not used at this level
            if i > 0 and nums[i] == nums[i - 1] and not used[i - 1]:
                continue

            used[i] = True
            current.append(nums[i])
            backtrack(current, used)
            current.pop()
            used[i] = False

    backtrack([], [False] * len(nums))
    return result


# =============================================================================
# TEMPLATE 3: Combinations (Choose k from n)
# Like subsets but only record when size == k.
# Time: O(C(n,k)), Space: O(k)
# =============================================================================
def combinations(n: int, k: int) -> List[List[int]]:
    """Generate all combinations of k numbers from [1..n]."""
    result = []

    def backtrack(start: int, current: List[int]):
        if len(current) == k:
            result.append(current[:])
            return

        # Pruning: need (k - len(current)) more elements
        # so we can only start from positions that leave enough remaining
        remaining_needed = k - len(current)
        for i in range(start, n - remaining_needed + 2):
            current.append(i)
            backtrack(i + 1, current)
            current.pop()

    backtrack(1, [])
    return result


def combination_sum(candidates: List[int], target: int) -> List[List[int]]:
    """Find all unique combinations that sum to target (reuse allowed)."""
    candidates.sort()
    result = []

    def backtrack(start: int, remaining: int, current: List[int]):
        if remaining == 0:
            result.append(current[:])
            return

        for i in range(start, len(candidates)):
            if candidates[i] > remaining:
                break  # prune: no point trying larger candidates

            current.append(candidates[i])
            backtrack(i, remaining - candidates[i], current)  # i not i+1: reuse OK
            current.pop()

    backtrack(0, target, [])
    return result


# =============================================================================
# TEMPLATE 4: Constraint Backtracking (N-Queens)
# Place elements on a grid/structure satisfying constraints.
# Prune aggressively using constraint sets.
# Time: depends on constraints, Space: O(n)
# =============================================================================
def solve_n_queens(n: int) -> List[List[str]]:
    """Place n queens on n×n board so no two attack each other."""
    result = []
    # Track occupied columns and diagonals
    cols = set()
    diag1 = set()  # row - col (identifies one diagonal direction)
    diag2 = set()  # row + col (identifies other diagonal direction)

    board = [['.' for _ in range(n)] for _ in range(n)]

    def backtrack(row: int):
        if row == n:
            # All queens placed successfully
            result.append([''.join(r) for r in board])
            return

        for col in range(n):
            # Check constraints (pruning)
            if col in cols or (row - col) in diag1 or (row + col) in diag2:
                continue

            # Place queen
            board[row][col] = 'Q'
            cols.add(col)
            diag1.add(row - col)
            diag2.add(row + col)

            # Explore next row
            backtrack(row + 1)

            # Remove queen (backtrack)
            board[row][col] = '.'
            cols.remove(col)
            diag1.remove(row - col)
            diag2.remove(row + col)

    backtrack(0)
    return result


# =============================================================================
if __name__ == "__main__":
    # Subsets
    print("Subsets [1,2,3]:", subsets([1, 2, 3]))
    print("Subsets with dup [1,2,2]:", subsets_with_dup([1, 2, 2]))

    # Permutations
    print("Permutations [1,2,3]:", permutations([1, 2, 3]))

    # Combinations
    print("C(4,2):", combinations(4, 2))
    print("Combination Sum [2,3,6,7] target=7:", combination_sum([2, 3, 6, 7], 7))

    # N-Queens
    solutions = solve_n_queens(4)
    print(f"\n4-Queens ({len(solutions)} solutions):")
    for sol in solutions:
        for row in sol:
            print(f"  {row}")
        print()
