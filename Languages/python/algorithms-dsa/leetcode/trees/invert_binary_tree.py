"""
LeetCode 226: Invert Binary Tree
Given the root of a binary tree, invert the tree and return its root.

Time: O(n), Space: O(h) where h is height
"""

from typing import Optional


class TreeNode:
    def __init__(self, val: int = 0, left=None, right=None):
        self.val = val
        self.left = left
        self.right = right


def invert_tree(root: Optional[TreeNode]) -> Optional[TreeNode]:
    if root is None:
        return None
    root.left, root.right = invert_tree(root.right), invert_tree(root.left)
    return root


def tree_to_list(root: Optional[TreeNode]) -> list:
    """Level-order traversal for testing."""
    if not root:
        return []
    result = []
    queue = [root]
    while queue:
        node = queue.pop(0)
        if node:
            result.append(node.val)
            queue.append(node.left)
            queue.append(node.right)
        else:
            result.append(None)
    # Remove trailing Nones
    while result and result[-1] is None:
        result.pop()
    return result


if __name__ == "__main__":
    # Tree: [4, 2, 7, 1, 3, 6, 9]
    root = TreeNode(4,
        TreeNode(2, TreeNode(1), TreeNode(3)),
        TreeNode(7, TreeNode(6), TreeNode(9))
    )
    inverted = invert_tree(root)
    assert tree_to_list(inverted) == [4, 7, 2, 9, 6, 3, 1]

    # Empty tree
    assert invert_tree(None) is None

    # Single node
    single = TreeNode(1)
    assert tree_to_list(invert_tree(single)) == [1]

    print("All tests passed!")
