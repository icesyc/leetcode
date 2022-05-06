/*
 * @lc app=leetcode id=25 lang=golang
 *
 * [25] Reverse Nodes in k-Group
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
/**
 * 解法2
 */
func reverseKGroup(head *ListNode, k int) *ListNode {
	left, right := head, head
	dummy := &ListNode{}
	jump := dummy
	for {
		i := 0
		for i = 0; i < k && right != nil; i++ {
			right = right.Next
		}
		if i < k {
			return dummy.Next
		}
		prev, cur := right, left
		for i := 0; i < k; i++ {
			tmp := cur.Next
			cur.Next = prev
			prev = cur
			cur = tmp
		}
		jump.Next, jump, left = prev, left, right

	}
}

//*/
/**
 * 解法1
 *
func reverseKGroup(head *ListNode, k int) *ListNode {
	var i int
	cur := head
	for i = 0; i < k; i++ {
		if cur == nil {
			return head
		}
		cur = cur.Next
	}
	prev := reverseKGroup(cur, k)
	for i = 0; i < k; i++ {
		tmp := head.Next
		head.Next = prev
		prev = head
		head = tmp
	}
	return prev
}

//*/
// @lc code=end

