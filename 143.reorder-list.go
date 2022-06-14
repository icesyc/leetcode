/*
 * @lc app=leetcode id=143 lang=golang
 *
 * [143] Reorder List
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
 * 使用一个数组保存所有节点
 * 二次循环，每次从数组尾部弹出一个元素，链接到当前位置的后面
 * 循环结束条件为cur == tail || cur.Next == tail
 * 对于奇数个数和尾数个数的条件有所区别
 *
 *  奇数
 *  1,   2,   3,   4,   5
 * cur next            tail
 *
 *  1,   5    2,   3,   4
 *           cur  next tail
 *
 *  1,   4    2,   5,   3
 *                     cur,next,tail
 *
 *  偶数
 *  1,   4,    2,   3
 *            cur  next tail
 */
/*
func reorderList(head *ListNode) {
	nodeList := []*ListNode{}
	cur := head
	for cur != nil {
		nodeList = append(nodeList, cur)
		cur = cur.Next
	}
	last := len(nodeList) - 1
	cur = head
	for {
		tail := nodeList[last]
		next := cur.Next
		if cur == tail || next == tail {
			tail.Next = nil
			break
		}
		cur.Next = tail
		tail.Next = next
		cur = next
		last--
	}
}
*/
/**
 * O(1)的解法
 * 1.找出中间结点
 * 2.反转后半个链表
 * 3.合并前后两个链表
 *
 * 1,2,3,4,5,6 => 1,2,3 | 4,5,6
 * 1,2,3 | 4,5,6 => 1,2,3 | 6,5,4
 * 1,2,3 | 6,5,4 => 1,6,2,5,3,4
 *
 */
func reorderList(head *ListNode) {
	middle := split(head)
	right := reverse(middle)
	merge(head, right)
}

//1,2,3,4,5,6 => 1,2,3 | 4,5,6
func split(head *ListNode) *ListNode {
	slow, fast := head, head
	for fast.Next != nil && fast.Next.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	middle := slow.Next
	slow.Next = nil
	return middle
}

//1,2,3 | 4,5,6 => 1,2,3 | 6,5,4
func reverse(head *ListNode) *ListNode {
	var pre *ListNode
	cur := head
	for cur != nil {
		tmp := cur.Next
		cur.Next = pre
		pre = cur
		cur = tmp
	}
	return pre
}

//1,2,3 | 6,5,4 => 1,6,2,5,3,4
func merge(left *ListNode, right *ListNode) *ListNode {
	head := left
	for right != nil {
		fmt.Println(left.Val, right.Val)
		lnext := left.Next
		rnext := right.Next
		left.Next = right
		right.Next = lnext
		left = lnext
		right = rnext
	}
	return head
}

// @lc code=end

