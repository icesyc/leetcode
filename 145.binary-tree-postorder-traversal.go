/*
 * @lc app=leetcode id=145 lang=golang
 *
 * [145] Binary Tree Postorder Traversal
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
 * 后序遍历 递归
 */
/*
func postorderTraversal(root *TreeNode) []int {
	res := []int{}
	if root != nil {
		res = append(res, postorderTraversal(root.Left)...)
		res = append(res, postorderTraversal(root.Right)...)
		res = append(res, root.Val)
	}
	return res
}
*/

/**
 * 后序遍历 使用stack, 每次将元素插入到结果前面即可
 */
func postorderTraversal(root *TreeNode) []int {
	res := []int{}
	stack := []*TreeNode{root}
	for len(stack) > 0 {
		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if node != nil {
			res = append([]int{node.Val}, res...)
			stack = append(stack, node.Left)
			stack = append(stack, node.Right)
		}
	}
	return res
}

// @lc code=end

