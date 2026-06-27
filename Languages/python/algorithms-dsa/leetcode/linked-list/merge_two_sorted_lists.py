"""
LeetCode #21 - Merge Two Sorted Lists
Topic: Linked List
Difficulty: Easy

Merge two sorted linked lists iteratively.

Time Complexity: O(n + m)
Space Complexity: O(1)
"""


class ListNode:
    def __init__(self, val: int = 0, next: "ListNode | None" = None):
        self.val = val
        self.next = next


def merge_two_lists(list1: ListNode | None, list2: ListNode | None) -> ListNode | None:
    dummy = ListNode(0)
    curr = dummy
    while list1 and list2:
        if list1.val <= list2.val:
            curr.next = list1
            list1 = list1.next
        else:
            curr.next = list2
            list2 = list2.next
        curr = curr.next
    curr.next = list1 or list2
    return dummy.next


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
    assert to_list(merge_two_lists(from_list([1, 2, 4]), from_list([1, 3, 4]))) == [1, 1, 2, 3, 4, 4]
    assert to_list(merge_two_lists(from_list([]), from_list([]))) == []
    assert to_list(merge_two_lists(from_list([]), from_list([0]))) == [0]
    print("All tests passed!")
