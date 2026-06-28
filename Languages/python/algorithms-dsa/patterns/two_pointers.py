"""
Two Pointers Pattern Templates

Three variants covering most two-pointer problems:
1. Opposite Ends converging from both sides (sorted arrays, pair sums)
2. Fast/Slow cycle detection, finding midpoints
3. Same Direction partitioning, removing elements in-place

Run: python two_pointers.py
"""

from typing import List, Optional


# =============================================================================
# TEMPLATE 1: Opposite Ends (Converging Pointers)
# Use when: sorted array, find pair with target sum, container problems
# Example: Find two numbers in sorted array that sum to target
# Time: O(n), Space: O(1)
# =============================================================================
def two_sum_sorted(nums: List[int], target: int) -> List[int]:
    """Find indices of two numbers that add up to target in a SORTED array."""
    left, right = 0, len(nums) - 1

    while left < right:
        current_sum = nums[left] + nums[right]

        if current_sum == target:
            return [left, right]
        elif current_sum < target:
            left += 1       # need larger sum → move left pointer right
        else:
            right -= 1      # need smaller sum → move right pointer left

    return []  # no pair found


def container_with_most_water(heights: List[int]) -> int:
    """Find max area between two lines (container with most water)."""
    left, right = 0, len(heights) - 1
    max_area = 0

    while left < right:
        width = right - left
        height = min(heights[left], heights[right])
        max_area = max(max_area, width * height)

        # Move the shorter side inward (it's the bottleneck)
        if heights[left] < heights[right]:
            left += 1
        else:
            right -= 1

    return max_area


# =============================================================================
# TEMPLATE 2: Fast/Slow Pointers (Floyd's Tortoise and Hare)
# Use when: cycle detection, finding middle, happy number
# Example: Detect cycle in linked list
# Time: O(n), Space: O(1)
# =============================================================================
class ListNode:
    def __init__(self, val=0, next=None):
        self.val = val
        self.next = next


def has_cycle(head: Optional[ListNode]) -> bool:
    """Detect if linked list has a cycle using Floyd's algorithm."""
    slow = fast = head

    while fast and fast.next:
        slow = slow.next          # moves 1 step
        fast = fast.next.next     # moves 2 steps

        if slow == fast:
            return True  # they met → cycle exists

    return False  # fast reached end → no cycle


def find_middle(head: Optional[ListNode]) -> Optional[ListNode]:
    """Find middle node. For even length, returns second middle."""
    slow = fast = head

    while fast and fast.next:
        slow = slow.next
        fast = fast.next.next

    return slow  # when fast reaches end, slow is at middle


def find_cycle_start(head: Optional[ListNode]) -> Optional[ListNode]:
    """Find the node where the cycle begins (Floyd's Phase 2)."""
    slow = fast = head

    # Phase 1: detect cycle
    while fast and fast.next:
        slow = slow.next
        fast = fast.next.next
        if slow == fast:
            break
    else:
        return None  # no cycle

    # Phase 2: find entry point
    # Move one pointer to head, advance both at same speed
    slow = head
    while slow != fast:
        slow = slow.next
        fast = fast.next

    return slow  # meeting point = cycle start


# =============================================================================
# TEMPLATE 3: Same Direction (Partition / Remove In-Place)
# Use when: removing elements, partitioning, Dutch National Flag
# Example: Move all zeroes to end while maintaining order
# Time: O(n), Space: O(1)
# =============================================================================
def move_zeroes(nums: List[int]) -> None:
    """Move all zeroes to end, maintaining relative order of non-zero elements."""
    # slow points to where next non-zero should be placed
    slow = 0

    for fast in range(len(nums)):
        if nums[fast] != 0:
            nums[slow], nums[fast] = nums[fast], nums[slow]
            slow += 1


def remove_duplicates_sorted(nums: List[int]) -> int:
    """Remove duplicates in-place from sorted array. Returns new length."""
    if not nums:
        return 0

    slow = 0  # last unique element position

    for fast in range(1, len(nums)):
        if nums[fast] != nums[slow]:
            slow += 1
            nums[slow] = nums[fast]

    return slow + 1  # length of unique portion


# =============================================================================
if __name__ == "__main__":
    # Opposite ends: two sum
    print("Two Sum Sorted [1,2,3,4,6] target=6:", two_sum_sorted([1, 2, 3, 4, 6], 6))

    # Container with most water
    print("Max water [1,8,6,2,5,4,8,3,7]:", container_with_most_water([1, 8, 6, 2, 5, 4, 8, 3, 7]))

    # Same direction: move zeroes
    arr = [0, 1, 0, 3, 12]
    move_zeroes(arr)
    print("Move zeroes [0,1,0,3,12]:", arr)

    # Remove duplicates
    arr2 = [1, 1, 2, 2, 3, 4, 4]
    new_len = remove_duplicates_sorted(arr2)
    print(f"Remove dups [1,1,2,2,3,4,4]: {arr2[:new_len]}")
