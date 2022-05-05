/*
 * @lc app=leetcode id=24 lang=golang
 *
 * [24] Swap Nodes in Pairs
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func swapPairs(head *ListNode) *ListNode {
	tmpHead := &ListNode{0, head}
	cur := tmpHead
	for cur != nil {
		if cur.Next == nil || cur.Next.Next == nil {
			break
		}
		first := cur.Next
		second := first.Next
		tmp := second.Next
		cur.Next = second
		second.Next = first
		first.Next = tmp
		cur = first
	}
	return tmpHead.Next
}

// @lc code=end

