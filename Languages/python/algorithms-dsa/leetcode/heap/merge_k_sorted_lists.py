"""
LeetCode #23 - Merge k Sorted Lists
Topic: Heap
Difficulty: Hard

Merge k sorted linked lists using a min-heap.

Time Complexity: O(n log k) where n = total nodes, k = number of lists
Space Complexity: O(k)
"""
import heapq


class ListNode:
    def __init__(self, val: int = 0, next: "ListNode | None" = None):
        self.val = val
        self.next = next


def merge_k_lists(lists: list[ListNode | None]) -> ListNode | None:
    heap: list[tuple[int, int, ListNode]] = []
    for i, node in enumerate(lists):
        if node:
            heapq.heappush(heap, (node.val, i, node))
    dummy = ListNode(0)
    curr = dummy
    while heap:
        val, i, node = heapq.heappop(heap)
        curr.next = node
        curr = curr.next
        if node.next:
            heapq.heappush(heap, (node.next.val, i, node.next))
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
    lists = [from_list([1, 4, 5]), from_list([1, 3, 4]), from_list([2, 6])]
    assert to_list(merge_k_lists(lists)) == [1, 1, 2, 3, 4, 4, 5, 6]
    assert to_list(merge_k_lists([])) == []
    assert to_list(merge_k_lists([None])) == []
    print("All tests passed!")
