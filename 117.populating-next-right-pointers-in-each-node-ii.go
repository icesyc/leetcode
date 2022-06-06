/*
 * @lc app=leetcode id=117 lang=golang
 *
 * [117] Populating Next Right Pointers in Each Node II
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
 * 同上一题类似，问题的区别是树不是完全BTree, 节点不是完全对称的
 * 先根据root.Next链表查找右侧要连接的节点
 * 循环结束条件为root.Next != nil
 * nextNode = root.Next.Left || roog.Next.Right
 * 同时在连接时要判断Left和Right是否存在
 * 设置正确的Next链接
 * 在处理子结点时，要先处理右节点,保证右链是完整的
 */
/*
func connect(root *Node) *Node {
	if root == nil {
		return root
	}
	var nextNode *Node
	parent := root
	for parent.Next != nil && nextNode == nil {
		nextNode = parent.Next.Left
		if nextNode == nil {
			nextNode = parent.Next.Right
		}
		parent = parent.Next
	}
	//右节点存在
	if root.Right != nil {
		root.Right.Next = nextNode
		nextNode = root.Right
	}
	//左节点存在
	if root.Left != nil {
		root.Left.Next = nextNode
	}
	//一定要先处理右节点，保证链表的完整性
	connect(root.Right)
	connect(root.Left)
	return root
}
*/
/**
 * 使用循环bfs遍历树
 * 关键点, 使用dummy做为临时节点
 * pre保存链表当前指针，处理完一层后进入下一层(dummy.Next)
 * 链接过程见图
 * https://raw.githubusercontent.com/icesyc/diagram/main/117.populating-next-right-pointers-in-each-node-ii.drawio.svg
 */
func connect(root *Node) *Node {
	dummy := &Node{}
	//保存上一个节点
	pre, cur := dummy, root
	for cur != nil {
		//左节点不为空，链接左节点，向前移动链表头部
		if cur.Left != nil {
			pre.Next = cur.Left
			pre = pre.Next
		}
		//右节点不为空，链链右节点，向前移动链表头部
		if cur.Right != nil {
			pre.Next = cur.Right
			pre = pre.Next
		}
		//横向移动当前节点
		cur = cur.Next
		//当前层次遍历完成，移向下一层
		//由于dummy和pre开始指向同一个节点
		//当前层次遍历完成后dummy.Next即为下一层的第一个节点
		if cur == nil {
			cur = dummy.Next
			//重置临时节点
			pre = dummy
			dummy.Next = nil
		}
	}
	return root
}

// @lc code=end

