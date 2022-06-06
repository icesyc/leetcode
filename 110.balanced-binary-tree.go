/*
 * @lc app=leetcode id=110 lang=golang
 *
 * [110] Balanced Binary Tree
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
 * dfs遍历获取子树的高度，如果子度不平衡，返回-1
 * 左右子树都平衡的话，判断左右子树高度差是否<2
 */
func isBalanced(root *TreeNode) bool {
	if root == nil {
		return true
	}
	return getHeight(root) != -1
}

func getHeight(root *TreeNode) int {
	if root == nil {
		return 0
	}
	leftHeight := getHeight(root.Left)
	//左树不平衡就没必要检查右树了
	if leftHeight == -1 {
		return -1
	}
	rightHeight := getHeight(root.Right)
	if rightHeight == -1 || abs(rightHeight-leftHeight) > 1 {
		return -1
	}
	return max(leftHeight, rightHeight) + 1
}
func abs(a int) int {
	if a < 0 {
		a = -a
	}
	return a
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// @lc code=end

