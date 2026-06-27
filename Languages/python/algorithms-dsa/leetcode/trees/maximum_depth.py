"""
LeetCode #104 - Maximum Depth of Binary Tree
Topic: Trees
Difficulty: Easy

Find max depth using recursive DFS.

Time Complexity: O(n)
Space Complexity: O(h) where h = tree height
"""


class TreeNode:
    def __init__(self, val: int = 0, left: "TreeNode | None" = None, right: "TreeNode | None" = None):
        self.val = val
        self.left = left
        self.right = right


def max_depth(root: TreeNode | None) -> int:
    if not root:
        return 0
    return 1 + max(max_depth(root.left), max_depth(root.right))


if __name__ == "__main__":
    tree = TreeNode(3, TreeNode(9), TreeNode(20, TreeNode(15), TreeNode(7)))
    assert max_depth(tree) == 3
    assert max_depth(TreeNode(1, None, TreeNode(2))) == 2
    assert max_depth(None) == 0
    print("All tests passed!")
