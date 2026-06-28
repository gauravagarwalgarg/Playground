/**
 * LeetCode 102: Binary Tree Level Order Traversal
 * Topic: Trees
 * Difficulty: Medium
 *
 * Return the level order traversal of a binary tree (BFS with queue).
 * Time: O(n), Space: O(n)
 */
#include <iostream>
#include <vector>
#include <queue>
#include <cassert>
using namespace std;

struct TreeNode {
    int val;
    TreeNode* left;
    TreeNode* right;
    TreeNode(int x) : val(x), left(nullptr), right(nullptr) {}
};

vector<vector<int>> levelOrder(TreeNode* root) {
    vector<vector<int>> res;
    if (!root) return res;
    queue<TreeNode*> q;
    q.push(root);
    while (!q.empty()) {
        int sz = q.size();
        vector<int> level;
        for (int i = 0; i < sz; i++) {
            TreeNode* node = q.front(); q.pop();
            level.push_back(node->val);
            if (node->left) q.push(node->left);
            if (node->right) q.push(node->right);
        }
        res.push_back(level);
    }
    return res;
}

int main() {
    // Tree: [3, 9, 20, null, null, 15, 7]
    TreeNode* root = new TreeNode(3);
    root->left = new TreeNode(9);
    root->right = new TreeNode(20);
    root->right->left = new TreeNode(15);
    root->right->right = new TreeNode(7);

    auto res = levelOrder(root);
    assert(res.size() == 3);
    assert(res[0] == vector<int>({3}));
    assert(res[1] == vector<int>({9, 20}));
    assert(res[2] == vector<int>({15, 7}));

    // Empty tree
    assert(levelOrder(nullptr).empty());

    cout << "All tests passed!" << endl;
    return 0;
}
