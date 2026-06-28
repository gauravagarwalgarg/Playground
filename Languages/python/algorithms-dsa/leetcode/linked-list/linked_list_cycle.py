"""
LeetCode #141 - Linked List Cycle
Topic: Linked List
Difficulty: Easy

Detect a cycle using Floyd's Tortoise and Hare algorithm.

Time Complexity: O(n)
Space Complexity: O(1)
"""


class ListNode:
    def __init__(self, val: int = 0, next: "ListNode | None" = None):
        self.val = val
        self.next = next


def has_cycle(head: ListNode | None) -> bool:
    slow = fast = head
    while fast and fast.next:
        slow = slow.next
        fast = fast.next.next
        if slow is fast:
            return True
    return False


if __name__ == "__main__":
    # Test 1: cycle exists
    node1 = ListNode(3)
    node2 = ListNode(2)
    node3 = ListNode(0)
    node4 = ListNode(-4)
    node1.next = node2
    node2.next = node3
    node3.next = node4
    node4.next = node2  # cycle
    assert has_cycle(node1) is True

    # Test 2: no cycle
    a = ListNode(1, ListNode(2, ListNode(3)))
    assert has_cycle(a) is False

    # Test 3: empty list
    assert has_cycle(None) is False
    print("All tests passed!")
