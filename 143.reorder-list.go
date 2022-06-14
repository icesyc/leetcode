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

// @lc code=end

