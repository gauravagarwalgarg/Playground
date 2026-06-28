#include <iostream>

using namespace std;

struct TreeNode {
		int val;
		TreeNode* left;
		TreeNode* right;
		Treenode(int data) : val(data), left(nullptr), right(nullptr) {}
};

int dfs(Treenode* node, int first_min) {
	if(!node) return numeric_limits<int>::max();

	if(node->val != first_min) return node->val;

	int left_second = dfs(node->left, first_min);
	int right_second = dfs(node->right, first_min);

	return min(left_second, right_second);
}

int findSecondMinimumValue(TreeNode* root) {
	if(!root || !root->left || !root->right) return -1;

	int second_min = dfs(root, root->val);
	return (second_min == numeric_limit<int>::max()) ? -1 : second_min;
}

int main() {
    TreeNode* root = new TreeNode(2);
    root->left = new TreeNode(2);
    root->right = new TreeNode(3);
    root->left->left = new TreeNode(5);
    root->left->right = new TreeNode(5);
    root->right->left = new TreeNode(3);
    root->right->right = new TreeNode(4);

    cout << findSecondMinimumValue(root) << endl; // Output: 3
    return 0;
}
