/*
 * @lc app=leetcode id=101 lang=golang
 *
 * [101] Symmetric Tree
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
 * stack解法
 * 原理同递归，每次push到stack里一对要比较的节点
 * 注意push的顺序即可
 */
func isSymmetric(root *TreeNode) bool {
	stack := []*TreeNode{root.Left, root.Right}
	for len(stack) > 0 {
		left, right := stack[len(stack)-2], stack[len(stack)-1]
		stack = stack[:len(stack)-2]
		if (left == nil || right == nil) && left != right {
			return false
		}
		if left == nil && right == nil {
			continue
		}
		if left.Val != right.Val {
			return false
		}
		stack = append(stack, left.Left, right.Right)
		stack = append(stack, left.Right, right.Left)
	}
	return true
}

/**
 * 递归解法
 * 既然是对称，那么2层的左右两节点必然相同
 * 再往下3层的左节点和3层的右节点必然相同
 * 递归的时候交换左右参数即可
 */
/*
func isSymmetric(root *TreeNode) bool {
	return traverse(root.Left, root.Right)
}

func traverse(left, right *TreeNode) bool {
	if left == nil && right == nil {
		return true
	} else if left == nil || right == nil {
		return false
	} else if left.Val != right.Val {
		return false
	}
	return traverse(left.Left, right.Right) && traverse(left.Right, right.Left)
}
*/
// @lc code=end

