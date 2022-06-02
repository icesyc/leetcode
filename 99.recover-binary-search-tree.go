/*
 * @lc app=leetcode id=99 lang=golang
 *
 * [99] Recover Binary Search Tree
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
 * morris中序遍历, O(1) space
 */
func recoverTree(root *TreeNode) {
	prev := &TreeNode{Val: math.MinInt}
	var first, second *TreeNode
	traverse(root, func(node *TreeNode) {
		if prev.Val >= node.Val {
			if first == nil {
				first = prev
			}
			second = node
		}
		prev = node
	})
	first.Val, second.Val = second.Val, first.Val
}

func traverse(root *TreeNode, visit func(*TreeNode)) {
	for root != nil {
		if root.Left == nil {
			visit(root)
			root = root.Right
		} else {
			prev := root.Left
			//如果链接过，那么prev.Right最终会=root
			for prev.Right != root && prev.Right != nil {
				prev = prev.Right
			}
			//连接左树的最右子节点到root上, 继续处理root.Left
			if prev.Right == nil {
				prev.Right = root
				root = root.Left
			} else {
				//移除链接
				prev.Right = nil
				visit(root)
				root = root.Right
			}
		}
	}
}

/**
 * 使用stack 中序遍历
 */
/*
func recoverTree(root *TreeNode) {
	stack := []*TreeNode{}
	var first, second *TreeNode
	prev := &TreeNode{Val: math.MinInt}
	for root != nil || len(stack) > 0 {
		for root != nil {
			stack = append(stack, root)
			root = root.Left
		}
		root = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		//假设遍历顺序为6,3,4,5,2
		//那么第一个错误节点为prev, 第二个为当前节点
		if prev.Val >= root.Val {
			if first == nil {
				first = prev
			}
			second = root
		}
		prev = root
		root = root.Right
	}
	first.Val, second.Val = second.Val, first.Val
}
*/
/**
 * 使用递归中序遍历
 */
/*
func recoverTree(root *TreeNode) {
	helperNode := make([]*TreeNode, 3)
	helperNode[0] = &TreeNode{Val: math.MinInt}
	traverse(root, helperNode)
	helperNode[1].Val, helperNode[2].Val = helperNode[2].Val, helperNode[1].Val
}

func traverse(root *TreeNode, helperNode []*TreeNode) {
	if root == nil {
		return
	}
	traverse(root.Left, helperNode)
	if helperNode[0].Val >= root.Val {
		if helperNode[1] == nil {
			helperNode[1] = helperNode[0]
		}
		helperNode[2] = root
	}
	helperNode[0] = root
	traverse(root.Right, helperNode)
}
*/

// @lc code=end

