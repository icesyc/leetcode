/*
 * @lc app=leetcode id=124 lang=golang
 *
 * [124] Binary Tree Maximum Path Sum
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
 * 对于一颗二叉树，路径之和的最大值有以下几种情况
 * 1. root本身
 * 2. root+left
 * 3. root+right
 * 4. left + root + right
 * 我们使用dfs postorder遍历, 返回该节点子树的最大和
 * 子树的最大和只能是(root, root+left,root+right)
 * 以例2图为例
 * 右节点20的子树最大值为20+15或20+17的最大值
 * 如果把左右节点都加上，那么上层节点的路径是不能同时包括15和7的
 */
func maxPathSum(root *TreeNode) int {
	res := math.MinInt
	traverse(root, &res)
	return res
}

func traverse(root *TreeNode, res *int) int {
	if root == nil {
		return 0
	}
	maxLeft := traverse(root.Left, res)
	maxRight := traverse(root.Right, res)
	currentMax := max(max(root.Val, root.Val+maxLeft), root.Val+maxRight)
	maxAll := max(currentMax, root.Val+maxLeft+maxRight)
	if maxAll > *res {
		*res = maxAll
	}
	return currentMax
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// @lc code=end

