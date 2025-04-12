#include <iostream>
#include <stack>

using namespace std;

struct TreeNode {
    int val;
    TreeNode* left;
    TreeNode* right;
    TreeNode(int x) : val(x), left(nullptr), right(nullptr) {}
};

class Solution {
public:
    int maxDepth(TreeNode* root) {
	    if(!root) return 0;

	    stack<pair<TreeNode*, int>> s;

	    int maxDepth = 0;

	    s.push({root, 1});

	    while(!s.empty())
	    {
		    auto [node, depth] = s.top();
		    s.pop();

		    maxDepth = max(maxDepth, depth);

		    if(node->left) s.push({node->left,  depth + 1});
		    if(node->right) s.push({node->right,  depth + 1});
	    }

	    return maxDepth;
   }
};

// Example usage
int main() {
    TreeNode* root = new TreeNode(1);
    root->left = new TreeNode(2);
    root->right = new TreeNode(3);
    root->left->left = new TreeNode(4);
    root->left->right = new TreeNode(5);

    Solution sol;
    cout << "Max Depth: " << sol.maxDepth(root) << endl; // Output: 3

    return 0;
}

