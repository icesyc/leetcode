/*
 * @lc app=leetcode id=114 lang=golang
 *
 * [114] Flatten Binary Tree to Linked List
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
 * dfs preorder遍历
 * O(1) space
 * 算法核心思想
 * 把root的左子树的最右节点连到右子树上, 循环处理
 * 处理之前cur=root, 下一次循环cur=cur.Right
 *       root             root
 *      /    \              \
 *   left    right  =>      left
 *   /   \                   /  \
 * left1 left2           left1  left2
 *                                \
 *                               right
 */
func flatten(root *TreeNode) {
	cur := root
	for cur != nil {
		if cur.Left != nil {
			tmp := cur.Left
			for tmp.Right != nil {
				tmp = tmp.Right
			}
			tmp.Right = cur.Right
			cur.Right = cur.Left
			cur.Left = nil
		}
		cur = cur.Right
	}
}

/**
 * dfs preorder遍历
 * O(n) space
 */
/*
func flatten(root *TreeNode) {
	stack := []*TreeNode{}
	if root != nil {
		stack = append(stack, root)
	}
	dummy := &TreeNode{}
	cur := dummy
	for len(stack) > 0 {
		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		cur.Right = &TreeNode{Val: node.Val}
		cur = cur.Right
		//先push右节点 pop的时候先pop左节点
		if node.Right != nil {
			stack = append(stack, node.Right)
		}
		if node.Left != nil {
			stack = append(stack, node.Left)
		}
	}
	if dummy.Right != nil {
		root.Left = nil
		root.Right = dummy.Right.Right
	}
}
*/

// @lc code=end

