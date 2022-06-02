/*
 * @lc app=leetcode id=98 lang=golang
 *
 * [98] Validate Binary Search Tree
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
 * 采用stack方式，使用中序遍历
 * 使用prev保存遍历的前一个节点，
 * 由于遍历顺序是left, root, right
 * 由于left < root < right
 * 所以如果前一个节点>=当前节点，说明不是BST
 */
func isValidBST(root *TreeNode) bool {
	stack := []*TreeNode{}
	var prev *TreeNode
	for root != nil || len(stack) > 0 {
		for root != nil {
			stack = append(stack, root)
			root = root.Left
		}
		//从stack中pop最左侧的
		root = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if prev != nil && root.Val <= prev.Val {
			return false
		}
		prev = root
		//处理右节点
		root = root.Right
	}
	return true
}

/*
 * 根据bst特性，传递当前节点值所在的区间, 递归判断
 */
/*
func isValidBST(root *TreeNode) bool {
	return isValidBTree(root, math.MinInt, math.MaxInt)
}

func isValidBTree(root *TreeNode, min, max int) bool {
	if root == nil {
		return true
	}
	if root.Val >= max || root.Val <= min {
		return false
	}
	return isValidBTree(root.Left, min, root.Val) && isValidBTree(root.Right, root.Val, max)
}
*/
// @lc code=end

