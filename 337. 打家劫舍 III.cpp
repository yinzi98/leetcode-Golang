小偷又发现一个新的可行窃的地点。 这个地区只有一个入口，称为“根”。 除了根部之外，每栋房子有且只有一个父房子。 
一番侦察之后，聪明的小偷意识到“这个地方的所有房屋形成了一棵二叉树”。 如果两个直接相连的房子在同一天晚上被打劫，房屋将自动报警。

在不触动警报的情况下，计算小偷一晚能盗取的最高金额。

示例 1:

     3
    / \
   2   3
    \   \ 
     3   1
能盗取的最高金额 = 3 + 3 + 1 = 7.

示例 2:

     3
    / \
   4   5
  / \   \ 
 1   3   1
能盗取的最高金额 = 4 + 5 = 9.



如果是我，肯定直接深搜+并查集
以后再转成Golang，这个大神的思路特别牛逼


struct TreeNode {
	int val;
	TreeNode *left;
	TreeNode *right;
	TreeNode(int x) : val(x), left(NULL), right(NULL) {}
};

class Solution {
	unordered_map<TreeNode*, int> records;
public:
    int rob(TreeNode* root) {
        if (root == NULL) return 0;
        if (records.count(root) == 0) {
        	int left = rob(root->left);
        	int right = rob(root->right);
        	int without_left = 0, without_right = 0;
        	if (root->left) without_left = rob(root->left->left) + rob(root->left->right);
        	if (root->right) without_right = rob(root->right->left) + rob(root->right->right);
        	records[root] = max(root->val+without_left+without_right, left+right);
        }
        return records[root];
    }
};
