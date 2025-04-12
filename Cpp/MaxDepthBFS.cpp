#include <iostream>
#include <queue>

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
			queue<TreeNode*> q;
				
			q.push(root);

			int depth = 0;

			while(!q.empty())
			{
				int size = q.size();
				depth++;

				for(int i = 0; i < size; i++)
				{
					TreeNode* node = q.front();
					q.pop();
					
					if(node->left) q.push(node->left);
					if(node->right) q.push(node->right);
				}
			}

			return depth;
		}
};

int main() {
	TreeNode* root = new TreeNode(1);
	root->left = new TreeNode(2);
	root->right = new TreeNode(3);
	root->left->left = new TreeNode(4);
	root->left->right = new TreeNode(5);

	Solution sol;
	cout << "Max Depth: " << sol.maxDepth(root) << endl;

	return 0;
}

