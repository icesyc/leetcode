/*
 * @lc app=leetcode id=105 lang=golang
 *
 * [105] Construct Binary Tree from Preorder and Inorder Traversal
 */

// @lc code=start
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
/**
 * 给定的2个数组，一个是preorder,一个是inorder
 *           root
 *         /      \
 *      left1   right1
 *       /   \
 *    left2 right2
 * 根据preorder的定义，访问顺序为root->left1->left2->right2->right1...
 * 根据inorder的定义，访问顺序为left2->left1->right2->root->right1...
 * preorder[0]为root, 我们根据root可以从inorder中获取到左子树的长度
 * 通过遍历inorder直到发现值为root，根据左树的长度，可以将preorder和inorder划分为两个子树
 * 这时候就是变成构建leftTree和rightTree的两个子问题
 * 递归即可
 */
func buildTree(preorder []int, inorder []int) *TreeNode {
	return build(preorder, inorder)
}

func build(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 {
		return nil
	}
	root := &TreeNode{Val: preorder[0]}
	//找到所有左侧的子树
	i := 0
	for i < len(inorder) && inorder[i] != root.Val {
		i++
	}
	root.Left = buildTree(preorder[1:i+1], inorder[:i])
	root.Right = buildTree(preorder[i+1:], inorder[i+1:])
	return root
}

// @lc code=end

