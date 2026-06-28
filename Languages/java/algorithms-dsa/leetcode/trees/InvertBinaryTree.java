import java.util.ArrayList;
import java.util.LinkedList;
import java.util.List;
import java.util.Queue;

/**
 * LeetCode 226: Invert Binary Tree
 * Given the root of a binary tree, invert the tree and return its root.
 * 
 * Time: O(n), Space: O(h)
 */
public class InvertBinaryTree {

    static class TreeNode {
        int val;
        TreeNode left, right;

        TreeNode(int val) {
            this.val = val;
        }

        TreeNode(int val, TreeNode left, TreeNode right) {
            this.val = val;
            this.left = left;
            this.right = right;
        }
    }

    public static TreeNode invertTree(TreeNode root) {
        if (root == null) return null;
        TreeNode temp = root.left;
        root.left = invertTree(root.right);
        root.right = invertTree(temp);
        return root;
    }

    public static List<Integer> levelOrder(TreeNode root) {
        List<Integer> result = new ArrayList<>();
        if (root == null) return result;
        Queue<TreeNode> queue = new LinkedList<>();
        queue.add(root);
        while (!queue.isEmpty()) {
            TreeNode node = queue.poll();
            result.add(node.val);
            if (node.left != null) queue.add(node.left);
            if (node.right != null) queue.add(node.right);
        }
        return result;
    }

    public static void main(String[] args) {
        // Tree: [4, 2, 7, 1, 3, 6, 9]
        TreeNode root = new TreeNode(4,
            new TreeNode(2, new TreeNode(1), new TreeNode(3)),
            new TreeNode(7, new TreeNode(6), new TreeNode(9))
        );

        TreeNode inverted = invertTree(root);
        List<Integer> result = levelOrder(inverted);
        List<Integer> expected = List.of(4, 7, 2, 9, 6, 3, 1);
        assert result.equals(expected) : "Tree inversion failed: " + result;

        // Null tree
        assert invertTree(null) == null : "Null tree test failed";

        System.out.println("All tests passed!");
    }
}
