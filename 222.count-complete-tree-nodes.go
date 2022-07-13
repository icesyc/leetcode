/*
 * @lc app=leetcode id=222 lang=golang
 *
 * [222] Count Complete Tree Nodes
 */

// @lc code=start
/**
 * 完全二叉树的节点个数
 * 对于满节点的二叉树，节点数为 (1 << h) - 1
 * 如果左树和右树高度相同
 * 说明最后一层的最后一个节点落在右侧, 左节点为高度h-1的满二叉树
 * 这时的节点数为左节点的节点数+根结点+右节点数
 * 1 << (h-1) - 1 + 1 + countNodes(root.Right)
 * 如果高度不同, 说明最后一层的最后一个节点落在左侧，右节点为高度h-2的满二叉树
 * 总结点数为1 << (h-2) - 1 + 1 + countNodes(root.Left)
 */
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func countNodes(root *TreeNode) int {
	if root == nil {
		return 0
	}
	lh := height(root.Left)
	rh := height(root.Right)
	if lh == rh {
		return 1<<lh + countNodes(root.Right)
	}
	return 1<<rh + countNodes(root.Left)
}

func height(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return 1 + height(root.Left)
}

// @lc code=end

