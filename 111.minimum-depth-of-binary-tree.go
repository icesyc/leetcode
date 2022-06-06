/*
 * @lc app=leetcode id=111 lang=golang
 *
 * [111] Minimum Depth of Binary Tree
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
 * 最小深度，使用dfs
 */
func minDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	leftDepth := minDepth(root.Left)
	rightDepth := minDepth(root.Right)
	//子节点有一个为空，返回另一个节点的深度+1
	if leftDepth == 0 || rightDepth == 0 {
		return leftDepth + rightDepth + 1
	} else {
		minChildDepth := leftDepth
		if minChildDepth > rightDepth {
			minChildDepth = rightDepth
		}
		return minChildDepth + 1
	}
}

/**
 * 最小深度，使用bfs
 */
/*
func minDepth(root *TreeNode) int {
	queue := []*TreeNode{}
	if root != nil {
		queue = append(queue, root)
	}
	minDepth := 0
	for len(queue) > 0 {
		minDepth++
		levelNums := len(queue)
		for i := 0; i < levelNums; i++ {
			node := queue[i]
			if node.Left == nil && node.Right == nil {
				return minDepth
			}
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
		queue = queue[levelNums:]
	}
	return minDepth
}
*/
// @lc code=end

