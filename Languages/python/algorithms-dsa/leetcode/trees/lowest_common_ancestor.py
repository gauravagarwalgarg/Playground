"""
LeetCode #236 - Lowest Common Ancestor of a Binary Tree
Topic: Trees
Difficulty: Medium

Find LCA using recursive post-order traversal.

Time Complexity: O(n)
Space Complexity: O(h)
"""


class TreeNode:
    def __init__(self, val: int = 0, left: "TreeNode | None" = None, right: "TreeNode | None" = None):
        self.val = val
        self.left = left
        self.right = right


def lowest_common_ancestor(root: TreeNode | None, p: TreeNode, q: TreeNode) -> TreeNode | None:
    if not root or root is p or root is q:
        return root
    left = lowest_common_ancestor(root.left, p, q)
    right = lowest_common_ancestor(root.right, p, q)
    if left and right:
        return root
    return left or right


if __name__ == "__main__":
    # Tree: [3,5,1,6,2,0,8,null,null,7,4]
    n7 = TreeNode(7)
    n4 = TreeNode(4)
    n6 = TreeNode(6)
    n2 = TreeNode(2, n7, n4)
    n5 = TreeNode(5, n6, n2)
    n0 = TreeNode(0)
    n8 = TreeNode(8)
    n1 = TreeNode(1, n0, n8)
    root = TreeNode(3, n5, n1)

    assert lowest_common_ancestor(root, n5, n1) is root
    assert lowest_common_ancestor(root, n5, n4) is n5
    print("All tests passed!")
