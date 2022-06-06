/*
 * @lc app=leetcode id=112 lang=golang
 *
 * [112] Path Sum
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
 * stack dfs
 * 找出从root到leaf的节点之和是否和targetSum相等
 * dfs遍历， 每次经过一个节点，减去该节点的值
 * 如果到leaf(node.Left, node.Right全是nil)
 * 看leaf的值是否与当前剩余的sum相同即可
 */
type TreeNodeWrap struct {
	node *TreeNode
	sum  int
}

func hasPathSum(root *TreeNode, targetSum int) bool {
	stack := []*TreeNodeWrap{}
	if root != nil {
		rootWrap := &TreeNodeWrap{root, targetSum}
		stack = append(stack, rootWrap)
	}
	for len(stack) > 0 {
		wrap := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if wrap.node.Left == nil && wrap.node.Right == nil && wrap.node.Val == wrap.sum {
			return true
		}
		if wrap.node.Left != nil {
			lwrap := &TreeNodeWrap{wrap.node.Left, wrap.sum - wrap.node.Val}
			stack = append(stack, lwrap)
		}
		if wrap.node.Right != nil {
			rwrap := &TreeNodeWrap{wrap.node.Right, wrap.sum - wrap.node.Val}
			stack = append(stack, rwrap)
		}
	}
	return false
}

/**
 * 递归dfs
 */
/*
func hasPathSum(root *TreeNode, targetSum int) bool {
	if root == nil {
		return false
	}
	if root.Left == nil && root.Right == nil {
		return targetSum == root.Val
	}
	res := hasPathSum(root.Left, targetSum-root.Val)
	res = res || hasPathSum(root.Right, targetSum-root.Val)
	return res
}
*/
// @lc code=end

