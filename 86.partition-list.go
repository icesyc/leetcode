/*
 * @lc app=leetcode id=86 lang=golang
 *
 * [86] Partition List
 */

// @lc code=start
/**
 * 双链表, 一个保存小于x的节点，一个保存>=x的节点
 * 最后连接即可
 */
func partition(head *ListNode, x int) *ListNode {
	p1, p2 := &ListNode{}, &ListNode{}
	dummy1, dummy2 := p1, p2
	for head != nil {
		if head.Val < x {
			p1.Next = head
			p1 = p1.Next
		} else {
			p2.Next = head
			p2 = p2.Next
		}
		head = head.Next
	}
	//避免链表成环
	p2.Next = nil
	p1.Next = dummy2.Next
	return dummy1.Next
}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
/**
 * 三个指针 prev指向当前节点的前一个, cur当前节点中
 * pivot指向小于x的最后一个节点
 * 当前节点小于x，判断pivot的Next是不是大于x
 * 是的话，重新设置poviot,将节点正确连接
 */
/*
func partition(head *ListNode, x int) *ListNode {
	dummy := &ListNode{-300, head}
	prev, pivot, cur := dummy, dummy, head
	for cur != nil {
		if cur.Val < x {
			next := cur.Next
			if pivot.Next.Val >= x {
				prev.Next = next
				cur.Next = pivot.Next
			}
			pivot.Next = cur
			pivot = cur
			cur = next
			continue
		}
		prev = cur
		cur = cur.Next
	}
	return dummy.Next
}
*/
// @lc code=end

