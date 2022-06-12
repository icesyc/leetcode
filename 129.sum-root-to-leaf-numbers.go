/*
 * @lc app=leetcode id=129 lang=golang
 *
 * [129] Sum Root to Leaf Numbers
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
 * 计算所有路径表示数字的加和
 * dfs搜索，每下到一层就需要进一位
 * 如果左右节点都不存在，则到达底部返回加和
 * 只有一个子节点时，需要继续遍历，然后对于空节点返回的值为0
 */
func sumNumbers(root *TreeNode) int {
	return dfs(root, 0)
}

func dfs(root *TreeNode, sum int) int {
	if root == nil {
		return 0
	}
	sum = sum*10 + root.Val
	if root.Left == nil && root.Right == nil {
		return sum
	}
	ls := dfs(root.Left, sum)
	rs := dfs(root.Right, sum)
	return ls + rs
}

// @lc code=end

