/*
 * @lc app=leetcode id=92 lang=golang
 *
 * [92] Reverse Linked List II
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseBetween(head *ListNode, left int, right int) *ListNode {
	dummy := &ListNode{Next: head}
	prev, cur := dummy, head
	//先移动cur到left的节点上
	for i := 1; i < left; i++ {
		prev = cur
		cur = cur.Next
	}
	leftPrev, leftStart := prev, cur
	//再遍历left:right区间,注意cur在right上时也要反转
	for left <= right {
		tmp := cur.Next
		cur.Next = prev
		prev = cur
		cur = tmp
		left++
	}
	leftPrev.Next = prev
	leftStart.Next = cur
	return dummy.Next
}

// @lc code=end

