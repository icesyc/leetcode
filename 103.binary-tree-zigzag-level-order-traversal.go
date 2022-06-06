/*
 * @lc app=leetcode id=103 lang=golang
 *
 * [103] Binary Tree Zigzag Level Order Traversal
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
 * dfs遍历, 增加leftToRight标识，判断方向
 * 如果是leftToRight 数组填充索引为为0,1,2... 否则为len-1, len-2...
 */
func zigzagLevelOrder(root *TreeNode) [][]int {
	res := [][]int{}
	queue := []*TreeNode{}
	if root != nil {
		queue = append(queue, root)
	}
	leftToRight := true
	for len(queue) > 0 {
		currentLevelNum := len(queue)
		currentVals := make([]int, currentLevelNum)
		for i := 0; i < currentLevelNum; i++ {
			node := queue[i]
			idx := i
			if !leftToRight {
				idx = currentLevelNum - 1 - i
			}
			currentVals[idx] = node.Val
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
		leftToRight = !leftToRight
		queue = queue[currentLevelNum:]
		res = append(res, currentVals)
	}
	return res
}

// @lc code=end

