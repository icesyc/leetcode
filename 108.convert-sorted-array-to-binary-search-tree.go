/*
 * @lc app=leetcode id=108 lang=golang
 *
 * [108] Convert Sorted Array to Binary Search Tree
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
 * 有序数组，为保证左右树平衡，中值为root
 * 再使用递归处理左右树
 */
func sortedArrayToBST(nums []int) *TreeNode {
	return helper(nums)
}
func helper(nums []int) *TreeNode {
	if len(nums) == 0 {
		return nil
	}
	mid := len(nums) / 2
	root := &TreeNode{Val: nums[mid]}
	root.Left = helper(nums[:mid])
	root.Right = helper(nums[mid+1:])
	return root
}

// @lc code=end

