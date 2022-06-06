/*
 * @lc app=leetcode id=107 lang=golang
 *
 * [107] Binary Tree Level Order Traversal II
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
 * bfs遍历，每次将当前level的所有子节点push到queue里
 * 处理完后，从queue中删除当前level的所有节点
 * 添加结果的时候添加到最前面
 */
func levelOrderBottom(root *TreeNode) [][]int {
	res := [][]int{}
	queue := []*TreeNode{}
	if root != nil {
		queue = append(queue, root)
	}
	for len(queue) > 0 {
		levelNums := len(queue)
		levelVals := []int{}
		for i := 0; i < levelNums; i++ {
			node := queue[i]
			levelVals = append(levelVals, node.Val)
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
		queue = queue[levelNums:]
		tmpRes := append([][]int{}, levelVals)
		res = append(tmpRes, res...)
	}
	return res
}

// @lc code=end

