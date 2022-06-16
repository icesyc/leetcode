/*
 * @lc app=leetcode id=147 lang=golang
 *
 * [147] Insertion Sort List
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
 * 链表的插入排序
 * 每次从头开始扫描，找到合适的位置插入
 * 注意：声明的哨兵节点(dummy)没有链接到原来的head上
 * 这样可以变成将一个链表插入到一个空链表中，减少程序的复杂度
 */
func insertionSortList(head *ListNode) *ListNode {
	cur := head
	dummy := &ListNode{}
	prev := dummy
	for cur != nil {
		next := cur.Next
		//优化点，每次prev指向已经排好序的倒数第2个节点
		//如果prev的值小于当前值，可以从prev后面开始扫描
		//而不是从头开始扫描
		if prev.Val > cur.Val {
			prev = dummy
		}
		for prev.Next != nil && prev.Next.Val < cur.Val {
			prev = prev.Next
		}
		//在prev和prev.Next中间插入
		cur.Next = prev.Next
		prev.Next = cur
		cur = next
	}
	return dummy.Next
}

// @lc code=end

