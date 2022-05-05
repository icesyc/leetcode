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
	n := 0
	for cur := head; cur != nil, cur = cur.Next {
		n++
	}
	prev := nil
	for i := 0; i < n / k; i++ {
		for j := 0; j < k; j++ {
			tmp := head.Next
			head.Next = prev
			prev = head
			head = tmp
		}
	}
	return prev
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

