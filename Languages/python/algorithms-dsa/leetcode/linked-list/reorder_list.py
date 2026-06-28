"""
LeetCode #143 - Reorder List
Topic: Linked List
Difficulty: Medium

Reorder list by finding middle, reversing second half, and interleaving.

Time Complexity: O(n)
Space Complexity: O(1)
"""


class ListNode:
    def __init__(self, val: int = 0, next: "ListNode | None" = None):
        self.val = val
        self.next = next


def reorder_list(head: ListNode | None) -> None:
    if not head or not head.next:
        return
    # Find middle
    slow, fast = head, head
    while fast.next and fast.next.next:
        slow = slow.next
        fast = fast.next.next
    # Reverse second half
    prev, curr = None, slow.next
    slow.next = None
    while curr:
        nxt = curr.next
        curr.next = prev
        prev = curr
        curr = nxt
    # Interleave
    first, second = head, prev
    while second:
        tmp1, tmp2 = first.next, second.next
        first.next = second
        second.next = tmp1
        first = tmp1
        second = tmp2


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
    head = from_list([1, 2, 3, 4])
    reorder_list(head)
    assert to_list(head) == [1, 4, 2, 3]

    head2 = from_list([1, 2, 3, 4, 5])
    reorder_list(head2)
    assert to_list(head2) == [1, 5, 2, 4, 3]
    print("All tests passed!")
