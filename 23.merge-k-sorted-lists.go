/*
 * @lc app=leetcode id=23 lang=golang
 *
 * [23] Merge k Sorted Lists
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
 * 解法2，归并法
 */
func mergeKLists(lists []*ListNode) *ListNode {
	if len(lists) == 0 {
		return nil
	}
	if len(lists) == 1 {
		return lists[0]
	}
	mid := len(lists) / 2
	left := mergeKLists(lists[0:mid])
	right := mergeKLists(lists[mid:])
	head := &ListNode{}
	cur := head
	for left != nil && right != nil {
		if left.Val < right.Val {
			cur.Next = left
			left = left.Next
		} else {
			cur.Next = right
			right = right.Next
		}
		cur = cur.Next
	}
	if left == nil {
		cur.Next = right
	} else {
		cur.Next = left
	}
	return head.Next
}

//*/
/**
 * 解法1
 *
func mergeKLists(lists []*ListNode) *ListNode {
	head := &ListNode{}
	cur := head
	for {
		minIdx, endNum := -1, 0
		for i := 0; i < len(lists); i++ {
			if lists[i] == nil {
				endNum++
				continue
			}
			if minIdx == -1 || lists[i].Val < lists[minIdx].Val {
				minIdx = i
			}
		}
		if endNum == len(lists) {
			break
		}
		cur.Next = lists[minIdx]
		cur = cur.Next
		lists[minIdx] = lists[minIdx].Next
	}
	return head.Next
}

//*/
// @lc code=end

