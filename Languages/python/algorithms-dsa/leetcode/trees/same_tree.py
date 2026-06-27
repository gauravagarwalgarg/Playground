"""
LeetCode #100 - Same Tree
Topic: Trees
Difficulty: Easy

Check if two binary trees are structurally identical using recursive comparison.

Time Complexity: O(n)
Space Complexity: O(h)
"""


class TreeNode:
    def __init__(self, val: int = 0, left: "TreeNode | None" = None, right: "TreeNode | None" = None):
        self.val = val
        self.left = left
        self.right = right


def is_same_tree(p: TreeNode | None, q: TreeNode | None) -> bool:
    if not p and not q:
        return True
    if not p or not q:
        return False
    return p.val == q.val and is_same_tree(p.left, q.left) and is_same_tree(p.right, q.right)


if __name__ == "__main__":
    t1 = TreeNode(1, TreeNode(2), TreeNode(3))
    t2 = TreeNode(1, TreeNode(2), TreeNode(3))
    assert is_same_tree(t1, t2) is True

    t3 = TreeNode(1, TreeNode(2))
    t4 = TreeNode(1, None, TreeNode(2))
    assert is_same_tree(t3, t4) is False

    assert is_same_tree(None, None) is True
    print("All tests passed!")
