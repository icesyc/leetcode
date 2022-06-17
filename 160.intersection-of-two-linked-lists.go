/*
 * @lc app=leetcode id=160 lang=golang
 *
 * [160] Intersection of Two Linked Lists
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
 * 查找链表交点，使用map保存node, 时间复杂度O(m+n)
 */
/*
func getIntersectionNode(headA, headB *ListNode) *ListNode {
	nodeMap := map[*ListNode]bool{}
	cur := headA
	for cur != nil {
		nodeMap[cur] = true
		cur = cur.Next
	}
	cur = headB
	for cur != nil {
		if nodeMap[cur] {
			return cur
		}
		cur = cur.Next
	}
	return nil
}
*/
/**
 * 两个链表相同的部分为c
 * 那么A = a+c, B = b+c
 * AB = a+c+b+c BA=b+c+a+c
 * 可以看到同时循环A,B，那么他们将在c相遇
 * 如果A,B没有交点，并且长度不同
 * 那么循环后curA=curB=nil, 都走了m+n长度
 */
func getIntersectionNode(headA, headB *ListNode) *ListNode {
	curA, curB := headA, headB
	for curA != curB {
		if curA == nil {
			curA = headB
		} else {
			curA = curA.Next
		}
		if curB == nil {
			curB = headA
		} else {
			curB = curB.Next
		}
	}
	return curA

}

// @lc code=end

