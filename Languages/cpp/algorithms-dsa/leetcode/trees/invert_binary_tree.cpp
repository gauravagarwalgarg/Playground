/**
 * LeetCode 226: Invert Binary Tree
 * Topic: Trees
 * Difficulty: Easy
 *
 * Invert a binary tree (mirror it).
 * Time: O(n), Space: O(h) where h = height
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

TreeNode* invertTree(TreeNode* root) {
    if (!root) return nullptr;
    swap(root->left, root->right);
    invertTree(root->left);
    invertTree(root->right);
    return root;
}

int main() {
    // Build: [4, 2, 7, 1, 3, 6, 9]
    TreeNode* root = new TreeNode(4);
    root->left = new TreeNode(2);
    root->right = new TreeNode(7);
    root->left->left = new TreeNode(1);
    root->left->right = new TreeNode(3);
    root->right->left = new TreeNode(6);
    root->right->right = new TreeNode(9);

    invertTree(root);

    // After invert: root->left should be 7, root->right should be 2
    assert(root->left->val == 7);
    assert(root->right->val == 2);
    assert(root->left->left->val == 9);
    assert(root->right->right->val == 1);

    cout << "All tests passed!" << endl;
    return 0;
}
