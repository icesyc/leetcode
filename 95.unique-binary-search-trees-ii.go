/*
 * @lc app=leetcode id=95 lang=golang
 *
 * [95] Unique Binary Search Trees II
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
 * DP[i]表示从Val从1-i对应的所有tree
 * DP[0] = [[]]
 * DP[1] = [[1]]
 * DP[i] = for j in [1:i] Node[Val=k, Left in DP[j-1], Right in DP[j+1:i]]
 * 对于数字为i所有的tree, 是以1:i每个数字做为root的情况
 * 当j做为root时, j为Val, dp[j-1]为所有可能的左tree
 * 对于j+1:i所表示的右树,需要做一下换算
 * 因为dp[i]表示的是1:i的范围，我们将右侧的范围长度两端都减去j
 * 得到1:i-j，也就是dp[i-j]就是我们计算的基础
 * 那么在生成右树的时候，我们需要把原来的右树的值都加上这个j，才是正确的
 *
 */
func generateTrees(n int) []*TreeNode {
	dp := make([][]*TreeNode, n+1)
	//为了使用leftTree, rightTree的循环正确执行，需要放置一个nil节点
	//否则计算dp[1]的时候就跳过循环了
	dp[0] = []*TreeNode{nil}
	for i := 1; i <= n; i++ {
		for j := 1; j <= i; j++ {
			for _, leftTree := range dp[j-1] {
				for _, rightTree := range dp[i-j] {
					node := &TreeNode{Val: j}
					node.Left = leftTree
					node.Right = copyTree(rightTree, j)
					dp[i] = append(dp[i], node)
				}
			}
		}
	}
	return dp[n]
}

func copyTree(root *TreeNode, offset int) *TreeNode {
	if root == nil {
		return nil
	}
	newRoot := &TreeNode{Val: root.Val + offset}
	newRoot.Left = copyTree(root.Left, offset)
	newRoot.Right = copyTree(root.Right, offset)
	return newRoot
}

/**
 * 递归
 */
/*
func generateTrees(n int) []*TreeNode {
	return helper(1, n)
}

func helper(start, end int) []*TreeNode {
	allNodes := []*TreeNode{}
	if start > end {
		allNodes = append(allNodes, nil)
		return allNodes
	}
	for i := start; i <= end; i++ {
		leftSubTree := helper(start, i-1)
		rightSubTree := helper(i+1, end)
		for j := 0; j < len(leftSubTree); j++ {
			for k := 0; k < len(rightSubTree); k++ {
				curNode := &TreeNode{Val: i}
				curNode.Left = leftSubTree[j]
				curNode.Right = rightSubTree[k]
				allNodes = append(allNodes, curNode)
			}
		}
	}
	return allNodes
}
*/
// @lc code=end

