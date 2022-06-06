/*
 * @lc app=leetcode id=109 lang=golang
 *
 * [109] Convert Sorted List to Binary Search Tree
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
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
/**
 * 使用快慢指针遍历链表，如果fast走完，那么slow为中值
 * 递归处理时需要传递end结点做为循环结束的条件
 * 递归结束的条件：只有一个结点时就直接返回
 */
func sortedListToBST(head *ListNode) *TreeNode {
	return helper(head, nil)
}
func helper(head *ListNode, endNode *ListNode) *TreeNode {
	if head == endNode {
		return nil
	}
	fast, slow := head, head
	for fast != endNode && fast.Next != endNode {
		slow = slow.Next
		fast = fast.Next.Next
	}
	root := &TreeNode{Val: slow.Val}
	root.Left = helper(head, slow)
	root.Right = helper(slow.Next, endNode)
	return root
}

// @lc code=end

