/*
 * @lc app=leetcode id=199 lang=golang
 *
 * [199] Binary Tree Right Side View
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
 * bfs遍历，只处理每层最右侧的节点
 */
/*
func rightSideView(root *TreeNode) []int {
	res := []int{}
	queue := []*TreeNode{}
	if root != nil {
		queue = append(queue, root)
	}
	for len(queue) > 0 {
		levelNum := len(queue)
		//只添加最右铡的节点到结果
		res = append(res, queue[0].Val)
		for i := 0; i < levelNum; i++ {
			node := queue[i]
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
		}
		queue = queue[levelNum:]
	}
	return res
}
*/
/**
 * dfs遍历，每层只添加一个节点
 */
func rightSideView(root *TreeNode) []int {
	res := &[]int{}
	dfs(root, res, 0)
	return *res
}

func dfs(root *TreeNode, res *[]int, depth int) {
	if root == nil {
		return
	}
	if len(*res) == depth {
		*res = append(*res, root.Val)
	}
	//先处理右节点
	dfs(root.Right, res, depth+1)
	dfs(root.Left, res, depth+1)
}

// @lc code=end

