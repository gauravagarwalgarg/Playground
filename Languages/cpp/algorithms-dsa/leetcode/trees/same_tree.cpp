/**
 * LeetCode 100: Same Tree
 * Topic: Trees
 * Difficulty: Easy
 *
 * Given the roots of two binary trees, check if they are the same.
 * Recursive comparison of structure and values.
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

bool isSameTree(TreeNode* p, TreeNode* q) {
    if (!p && !q) return true;
    if (!p || !q) return false;
    return p->val == q->val &&
           isSameTree(p->left, q->left) &&
           isSameTree(p->right, q->right);
}

int main() {
    // Tree 1: [1, 2, 3]
    TreeNode* p1 = new TreeNode(1);
    p1->left = new TreeNode(2);
    p1->right = new TreeNode(3);

    TreeNode* q1 = new TreeNode(1);
    q1->left = new TreeNode(2);
    q1->right = new TreeNode(3);
    assert(isSameTree(p1, q1) == true);

    // Different structure
    TreeNode* p2 = new TreeNode(1);
    p2->left = new TreeNode(2);

    TreeNode* q2 = new TreeNode(1);
    q2->right = new TreeNode(2);
    assert(isSameTree(p2, q2) == false);

    // Both null
    assert(isSameTree(nullptr, nullptr) == true);

    cout << "All tests passed!" << endl;
    return 0;
}
