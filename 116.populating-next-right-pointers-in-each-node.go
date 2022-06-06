/*
 * @lc app=leetcode id=116 lang=golang
 *
 * [116] Populating Next Right Pointers in Each Node
 */

// @lc code=start
/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Left *Node
 *     Right *Node
 *     Next *Node
 * }
 */
/**
 * dfs + queue, 每次处理一层的节点
 * 入队从右子节点开始，循环的时候每次将当前结点指向上次循环的节点
 */
/*
func connect(root *Node) *Node {
	queue := []*Node{}
	if root == nil {
		return root
	}
	queue = append(queue, root)
	for len(queue) > 0 {
		levelNum := len(queue)
		var right *Node
		for i := 0; i < levelNum; i++ {
			node := queue[i]
			node.Next = right
			right = node
			//先push right节点
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
		}
		queue = queue[levelNum:]
	}
	return root
}
*/
/**
 * Follow up要求使用O(1) space, 可使用递归
 */
/*
func connect(root *Node) *Node {
	if root == nil {
		return root
	}
	if root.Left != nil {
		root.Left.Next = root.Right
		if root.Next != nil {
			root.Right.Next = root.Next.Left
		}
		connect(root.Left)
		connect(root.Right)
	}
	return root
}
*/
/**
 * Follow up要求使用O(1) space, 不使用递归
 * 关键点 root.Left = root.Right
 *       root.Right = root.Next.Left
 * 处理前先保存当前层次的最左节点，处理完一层后恢复指针
 * 由于从上往下处理，所以可保证上层的Next是先处理完成的
 */
func connect(root *Node) *Node {
	cur := root
	for cur != nil && cur.Left != nil {
		//先将下一层的最左节点临时保存
		nextLevel := cur.Left
		for cur != nil {
			cur.Left.Next = cur.Right
			if cur.Next != nil {
				cur.Right.Next = cur.Next.Left
			}
			cur = cur.Next
		}
		//恢复指针，继续处理下一层
		cur = nextLevel
	}
	return root
}

// @lc code=end

