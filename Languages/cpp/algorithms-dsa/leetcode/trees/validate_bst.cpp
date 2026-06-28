/**
 * LeetCode 98: Validate Binary Search Tree
 * Topic: Trees
 * Difficulty: Medium
 *
 * Given the root of a binary tree, determine if it is a valid BST.
 * Inorder bounds check: each node must be within (low, high) range.
 * Time: O(n), Space: O(h)
 */
#include <iostream>
#include <climits>
#include <cassert>
using namespace std;

struct TreeNode {
    int val;
    TreeNode* left;
    TreeNode* right;
    TreeNode(int x) : val(x), left(nullptr), right(nullptr) {}
};

bool validate(TreeNode* node, long low, long high) {
    if (!node) return true;
    if (node->val <= low || node->val >= high) return false;
    return validate(node->left, low, node->val) &&
           validate(node->right, node->val, high);
}

bool isValidBST(TreeNode* root) {
    return validate(root, LONG_MIN, LONG_MAX);
}

int main() {
    // Valid BST: [2, 1, 3]
    TreeNode* r1 = new TreeNode(2);
    r1->left = new TreeNode(1);
    r1->right = new TreeNode(3);
    assert(isValidBST(r1) == true);

    // Invalid BST: [5, 1, 4, null, null, 3, 6]
    TreeNode* r2 = new TreeNode(5);
    r2->left = new TreeNode(1);
    r2->right = new TreeNode(4);
    r2->right->left = new TreeNode(3);
    r2->right->right = new TreeNode(6);
    assert(isValidBST(r2) == false);

    // Single node
    TreeNode* r3 = new TreeNode(1);
    assert(isValidBST(r3) == true);

    cout << "All tests passed!" << endl;
    return 0;
}
