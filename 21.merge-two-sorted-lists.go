/*
 * @lc app=leetcode id=21 lang=golang
 *
 * [21] Merge Two Sorted Lists
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	mergedNode := &ListNode{}
	curNode := mergedNode
	for list1 != nil && list2 != nil {
		if list1.Val < list2.Val {
			curNode.Next = list1
			list1 = list1.Next
		} else {
			curNode.Next = list2
			list2 = list2.Next
		}
		curNode = curNode.Next
	}
	if list1 == nil {
		curNode.Next = list2
	} else {
		curNode.Next = list1
	}
	return mergedNode.Next
}

// @lc code=end

