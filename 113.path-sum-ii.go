/*
 * @lc app=leetcode id=113 lang=golang
 *
 * [113] Path Sum II
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
 * 同上一题, dfs + 递归
 */
/*
func pathSum(root *TreeNode, targetSum int) [][]int {
	res := &[][]int{}
	dfs(root, targetSum, []int{}, res)
	return *res
}
func dfs(root *TreeNode, targetSum int, cur []int, res *[][]int) {
	if root == nil {
		return
	}
	cur = append(cur, root.Val)
	if root.Left == nil && root.Right == nil && root.Val == targetSum {
		*res = append(*res, append([]int{}, cur...))
	}
	dfs(root.Left, targetSum-root.Val, cur, res)
	dfs(root.Right, targetSum-root.Val, cur, res)
}
*/
/**
 * 同上一题，使用stack处理
 * stack需要保存当前节点，剩余的sum,和祖先路径
 */
type TreeNodeWrap struct {
	node *TreeNode
	sum  int
	cur  []int
}

func pathSum(root *TreeNode, targetSum int) [][]int {
	stack := []*TreeNodeWrap{}
	if root != nil {
		rootWrap := &TreeNodeWrap{root, targetSum, []int{}}
		stack = append(stack, rootWrap)
	}
	res := [][]int{}
	for len(stack) > 0 {
		wrap := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		cur := append(wrap.cur, wrap.node.Val)
		if wrap.node.Left == nil && wrap.node.Right == nil {
			if wrap.node.Val == wrap.sum {
				res = append(res, append([]int{}, cur...))
			}
		}
		if wrap.node.Left != nil {
			lwrap := &TreeNodeWrap{wrap.node.Left, wrap.sum - wrap.node.Val, cur}
			stack = append(stack, lwrap)
		}
		if wrap.node.Right != nil {
			rwrap := &TreeNodeWrap{wrap.node.Right, wrap.sum - wrap.node.Val, cur}
			stack = append(stack, rwrap)
		}
	}
	return res
}

// @lc code=end

