/*
 * @lc app=leetcode id=148 lang=golang
 *
 * [148] Sort List
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
 * 链表排序, 插入排序见上一题实现
 * 使用归并排序
 */
func sortList(head *ListNode) *ListNode {
	return divideAndConquer(head)
}

func divideAndConquer(head *ListNode) *ListNode {
	//1个节点
	if head == nil || head.Next == nil {
		return head
	}
	fast, slow := head, head
	for fast.Next != nil && fast.Next.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	mid := slow.Next
	slow.Next = nil
	left := divideAndConquer(head)
	right := divideAndConquer(mid)
	return mergeList(left, right)
}

func mergeList(left, right *ListNode) *ListNode {
	dummy := &ListNode{}
	cur := dummy
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
	//左右个数不相等时，有一个必然还有剩余
	if left != nil {
		cur.Next = left
	}
	if right != nil {
		cur.Next = right
	}
	return dummy.Next
}

// @lc code=end

