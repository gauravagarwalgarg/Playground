"""
LeetCode #206 - Reverse Linked List
Topic: Linked List
Difficulty: Easy

Reverse a singly linked list iteratively.

Time Complexity: O(n)
Space Complexity: O(1)
"""


class ListNode:
    def __init__(self, val: int = 0, next: "ListNode | None" = None):
        self.val = val
        self.next = next


def reverse_list(head: ListNode | None) -> ListNode | None:
    prev = None
    curr = head
    while curr:
        nxt = curr.next
        curr.next = prev
        prev = curr
        curr = nxt
    return prev


def to_list(head: ListNode | None) -> list[int]:
    result = []
    while head:
        result.append(head.val)
        head = head.next
    return result


def from_list(vals: list[int]) -> ListNode | None:
    dummy = ListNode(0)
    curr = dummy
    for v in vals:
        curr.next = ListNode(v)
        curr = curr.next
    return dummy.next


if __name__ == "__main__":
    assert to_list(reverse_list(from_list([1, 2, 3, 4, 5]))) == [5, 4, 3, 2, 1]
    assert to_list(reverse_list(from_list([1, 2]))) == [2, 1]
    assert to_list(reverse_list(from_list([]))) == []
    print("All tests passed!")
