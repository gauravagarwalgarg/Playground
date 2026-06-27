"""
LeetCode #102 - Binary Tree Level Order Traversal
Topic: Trees
Difficulty: Medium

BFS level order traversal using deque.

Time Complexity: O(n)
Space Complexity: O(n)
"""
from collections import deque


class TreeNode:
    def __init__(self, val: int = 0, left: "TreeNode | None" = None, right: "TreeNode | None" = None):
        self.val = val
        self.left = left
        self.right = right


def level_order(root: TreeNode | None) -> list[list[int]]:
    if not root:
        return []
    result: list[list[int]] = []
    queue = deque([root])
    while queue:
        level: list[int] = []
        for _ in range(len(queue)):
            node = queue.popleft()
            level.append(node.val)
            if node.left:
                queue.append(node.left)
            if node.right:
                queue.append(node.right)
        result.append(level)
    return result


if __name__ == "__main__":
    tree = TreeNode(3, TreeNode(9), TreeNode(20, TreeNode(15), TreeNode(7)))
    assert level_order(tree) == [[3], [9, 20], [15, 7]]
    assert level_order(TreeNode(1)) == [[1]]
    assert level_order(None) == []
    print("All tests passed!")
