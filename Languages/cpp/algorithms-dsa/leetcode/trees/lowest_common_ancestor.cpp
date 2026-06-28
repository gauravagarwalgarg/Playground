/**
 * LeetCode 236: Lowest Common Ancestor of a Binary Tree
 * Topic: Trees
 * Difficulty: Medium
 *
 * Given a binary tree and two nodes, find their lowest common ancestor.
 * Recursive split: if p and q are on different sides, current node is LCA.
 * Time: O(n), Space: O(h)
 */
#include <iostream>
#include <cassert>
using namespace std;

struct TreeNode {
    int val;
    TreeNode* left;
    TreeNode* right;
    TreeNode(int x) : val(x), left(nullptr), right(nullptr) {}
};

TreeNode* lowestCommonAncestor(TreeNode* root, TreeNode* p, TreeNode* q) {
    if (!root || root == p || root == q) return root;
    TreeNode* left = lowestCommonAncestor(root->left, p, q);
    TreeNode* right = lowestCommonAncestor(root->right, p, q);
    if (left && right) return root;
    return left ? left : right;
}

int main() {
    // Tree: [3, 5, 1, 6, 2, 0, 8, null, null, 7, 4]
    TreeNode* root = new TreeNode(3);
    root->left = new TreeNode(5);
    root->right = new TreeNode(1);
    root->left->left = new TreeNode(6);
    root->left->right = new TreeNode(2);
    root->right->left = new TreeNode(0);
    root->right->right = new TreeNode(8);
    root->left->right->left = new TreeNode(7);
    root->left->right->right = new TreeNode(4);

    TreeNode* p = root->left;        // node 5
    TreeNode* q = root->right;       // node 1
    assert(lowestCommonAncestor(root, p, q) == root);

    TreeNode* p2 = root->left;       // node 5
    TreeNode* q2 = root->left->right->right; // node 4
    assert(lowestCommonAncestor(root, p2, q2) == root->left);

    cout << "All tests passed!" << endl;
    return 0;
}
