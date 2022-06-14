/*
 * @lc app=leetcode id=142 lang=golang
 *
 * [142] Linked List Cycle II
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
 * 返回有环列表中环形开始的结点
 * 参考图片
 * https://farm6.staticflickr.com/5758/22715587283_bdb4ba8434.jpg
 * fast和slow指针同时从head出发
 * fast每次走2步，slow走一步
 * 如果fast和slow相遇在p点，
 * 那么fast走的距离为a+2b+c, slow走的距离为a+b
 * 由于fast是slow的2倍速度，所以a+2b+c=2(a+b)
 * 得到a=c, 我们再让一个slow2指针从head出发，
 * slow从p点出发, 由到a=c,他们的相遇结点必然为q
 * q即为我们要返回的结果
 */
func detectCycle(head *ListNode) *ListNode {
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
		if slow == fast {
			slow2 := head
			for slow != slow2 {
				slow = slow.Next
				slow2 = slow2.Next
			}
			return slow
		}
	}
	return nil
}

// @lc code=end

