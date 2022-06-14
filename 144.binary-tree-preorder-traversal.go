/*
 * @lc app=leetcode id=144 lang=golang
 *
 * [144] Binary Tree Preorder Traversal
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
 * 前序节点遍历 递归
 */
/*
func preorderTraversal(root *TreeNode) []int {
	res := []int{}
	if root != nil {
		res = append(res, root.Val)
		res = append(res, preorderTraversal(root.Left)...)
		res = append(res, preorderTraversal(root.Right)...)
	}
	return res
}
*/
/**
 * 使用stack
 */
func preorderTraversal(root *TreeNode) []int {
	res := []int{}
	stack := []*TreeNode{root}
	for len(stack) > 0 {
		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if node != nil {
			res = append(res, node.Val)
			stack = append(stack, node.Right)
			stack = append(stack, node.Left)
		}
	}
	return res
}

// @lc code=end

