/*
 * @lc app=leetcode id=61 lang=golang
 *
 * [61] Rotate List
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func rotateRight(head *ListNode, k int) *ListNode {
	if head == nil {
		return head
	}
	last, n := head, 1
	for last.Next != nil {
		last = last.Next
		n++
	}
	k = k % n
	if k > 0 {
		last.Next = head
		tmp := head
		for i := 0; i < n-k-1; i++ {
			tmp = tmp.Next
		}
		head = tmp.Next
		tmp.Next = nil
	}
	return head
}

// @lc code=end

