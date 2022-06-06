/*
 * @lc app=leetcode id=104 lang=golang
 *
 * [104] Maximum Depth of Binary Tree
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
 * DFS 递归
 */
func maxDepth(root *TreeNode) int {
	return dfs(root)
}
func dfs(root *TreeNode) int {
	if root == nil {
		return 0
	}
	max := dfs(root.Left)
	maxRight := dfs(root.Right)
	if maxRight > max {
		max = maxRight
	}
	//当前层级加1
	return max + 1
}

/**
 * 循环bfs
 */
/*
func maxDepth(root *TreeNode) int {
	return bfs(root)
}

func bfs(root *TreeNode) int {
	queue := []*TreeNode{}
	if root != nil {
		queue = append(queue, root)
	}
	depth := 0
	for len(queue) > 0 {
		nums := len(queue)
		for i := 0; i < nums; i++ {
			node := queue[i]
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
		queue = queue[nums:]
		depth++
	}
	return depth
}
*/

// @lc code=end

