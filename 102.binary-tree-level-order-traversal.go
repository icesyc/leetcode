/*
 * @lc app=leetcode id=102 lang=golang
 *
 * [102] Binary Tree Level Order Traversal
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
 * BFS搜索, 使用队列存储节点
 * 循环每次将node的子节点加到队列中
 * 内循环处理当前level的所有节点
 */
func levelOrder(root *TreeNode) [][]int {
	return bfs(root)
}

func bfs(root *TreeNode) [][]int {
	res := [][]int{}
	queue := []*TreeNode{}
	if root != nil {
		queue = append(queue, root)
	}
	for len(queue) > 0 {
		currentLevelNum := len(queue)
		currentVals := make([]int, currentLevelNum)
		for i := 0; i < currentLevelNum; i++ {
			node := queue[i]
			currentVals[i] = node.Val
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
		queue = queue[currentLevelNum:]
		res = append(res, currentVals)
	}
	return res
}

// @lc code=end

