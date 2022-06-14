/*
 * @lc app=leetcode id=138 lang=golang
 *
 * [138] Copy List with Random Pointer
 */

// @lc code=start
/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Next *Node
 *     Random *Node
 * }
 */
/**
 * 使用hashmap保存新旧节点的映射关系
 * 第二次循环时将对应的指针填充
 */
/*
func copyRandomList(head *Node) *Node {
	if head == nil {
		return nil
	}
	nodeMap := map[*Node]*Node{}
	cur := head
	for cur != nil {
		node := &Node{Val: cur.Val}
		nodeMap[cur] = node
		cur = cur.Next
	}
	cur = head
	for cur != nil {
		copy := nodeMap[cur]
		copy.Next = nodeMap[cur.Next]
		copy.Random = nodeMap[cur.Random]
		cur = cur.Next
	}
	return nodeMap[head]
}
*/

/**
 * O(1)的方案
 * 1 在旧链表上创建新的链表  1->2->3->4 => 1->1'->2->2'->3->3'->4->4'
 * 2 将新链表的random连接正确 1->1'->2->2'->3->3'->4->4'
 *                            ^-------------^
 * 3.将新链表和旧链表分离
 */
func copyRandomList(head *Node) *Node {
	if head == nil {
		return nil
	}
	cur := head
	for cur != nil {
		next := cur.Next
		cur.Next = &Node{Val: cur.Val}
		cur.Next.Next = next
		cur = next
	}

	cur = head
	for cur != nil {
		if cur.Random != nil {
			cur.Next.Random = cur.Random.Next
		}
		cur = cur.Next.Next
	}

	cur = head
	copyHead := cur.Next
	copy := copyHead

	//1->1'->2->2'->3->3'->4->4' => 1->2->3->4, 1'->2'->3'->4'
	for copy.Next != nil {
		cur.Next = cur.Next.Next
		cur = cur.Next

		copy.Next = copy.Next.Next
		copy = copy.Next
	}
	//最后节点(4)的指针恢复
	cur.Next = nil
	return copyHead
}

// @lc code=end

