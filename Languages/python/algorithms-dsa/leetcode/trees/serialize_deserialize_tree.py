"""
LeetCode #297 - Serialize and Deserialize Binary Tree
Topic: Trees
Difficulty: Hard

Serialize/deserialize binary tree using BFS level-order.

Time Complexity: O(n)
Space Complexity: O(n)
"""
from collections import deque


class TreeNode:
    def __init__(self, val: int = 0, left: "TreeNode | None" = None, right: "TreeNode | None" = None):
        self.val = val
        self.left = left
        self.right = right


def serialize(root: TreeNode | None) -> str:
    if not root:
        return ""
    result: list[str] = []
    queue = deque([root])
    while queue:
        node = queue.popleft()
        if node:
            result.append(str(node.val))
            queue.append(node.left)
            queue.append(node.right)
        else:
            result.append("null")
    return ",".join(result)


def deserialize(data: str) -> TreeNode | None:
    if not data:
        return None
    vals = data.split(",")
    root = TreeNode(int(vals[0]))
    queue = deque([root])
    i = 1
    while queue and i < len(vals):
        node = queue.popleft()
        if vals[i] != "null":
            node.left = TreeNode(int(vals[i]))
            queue.append(node.left)
        i += 1
        if i < len(vals) and vals[i] != "null":
            node.right = TreeNode(int(vals[i]))
            queue.append(node.right)
        i += 1
    return root


if __name__ == "__main__":
    tree = TreeNode(1, TreeNode(2), TreeNode(3, TreeNode(4), TreeNode(5)))
    s = serialize(tree)
    restored = deserialize(s)
    assert serialize(restored) == s
    assert deserialize("") is None
    print("All tests passed!")
