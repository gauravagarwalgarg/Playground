/**
 * LeetCode 297: Serialize and Deserialize Binary Tree
 * Topic: Trees
 * Difficulty: Hard
 *
 * Design an algorithm to serialize and deserialize a binary tree.
 * BFS serialization with "null" markers for missing nodes.
 * Time: O(n), Space: O(n)
 */
#include <iostream>
#include <string>
#include <sstream>
#include <queue>
#include <cassert>
using namespace std;

struct TreeNode {
    int val;
    TreeNode* left;
    TreeNode* right;
    TreeNode(int x) : val(x), left(nullptr), right(nullptr) {}
};

string serialize(TreeNode* root) {
    if (!root) return "null";
    string res;
    queue<TreeNode*> q;
    q.push(root);
    while (!q.empty()) {
        TreeNode* node = q.front(); q.pop();
        if (node) {
            res += to_string(node->val) + ",";
            q.push(node->left);
            q.push(node->right);
        } else {
            res += "null,";
        }
    }
    return res;
}

TreeNode* deserialize(string data) {
    if (data == "null") return nullptr;
    stringstream ss(data);
    string token;
    getline(ss, token, ',');
    TreeNode* root = new TreeNode(stoi(token));
    queue<TreeNode*> q;
    q.push(root);
    while (!q.empty()) {
        TreeNode* node = q.front(); q.pop();
        if (getline(ss, token, ',') && token != "null") {
            node->left = new TreeNode(stoi(token));
            q.push(node->left);
        }
        if (getline(ss, token, ',') && token != "null") {
            node->right = new TreeNode(stoi(token));
            q.push(node->right);
        }
    }
    return root;
}

int main() {
    // Build tree: [1, 2, 3, null, null, 4, 5]
    TreeNode* root = new TreeNode(1);
    root->left = new TreeNode(2);
    root->right = new TreeNode(3);
    root->right->left = new TreeNode(4);
    root->right->right = new TreeNode(5);

    string encoded = serialize(root);
    TreeNode* decoded = deserialize(encoded);

    assert(decoded->val == 1);
    assert(decoded->left->val == 2);
    assert(decoded->right->val == 3);
    assert(decoded->right->left->val == 4);
    assert(decoded->right->right->val == 5);
    assert(decoded->left->left == nullptr);

    // Empty tree
    assert(deserialize("null") == nullptr);

    cout << "All tests passed!" << endl;
    return 0;
}
