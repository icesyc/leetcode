/*
 * @lc app=leetcode id=173 lang=golang
 *
 * [173] Binary Search Tree Iterator
 */

// @lc code=start
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
/**
 * 中序遍历
 * 思路，使用一个stack保存所有最左侧的节点
 * 每次调用next时弹出一个节点
 * 同时检查该节点的right，如果存在，继续将right的所有最左子节点入栈
 * 这样保证next()的平均时间复杂度为O(1),空间复杂度为O(h)
 * h为树的高度, stack最多保存h个节点
 */
type BSTIterator struct {
	stack []*TreeNode
}

func Constructor(root *TreeNode) BSTIterator {
	bst := BSTIterator{}
	cur := root
	for cur != nil {
		bst.stack = append(bst.stack, cur)
		cur = cur.Left
	}
	return bst
}

func (this *BSTIterator) Next() int {
	node := this.stack[len(this.stack)-1]
	this.stack = this.stack[:len(this.stack)-1]
	cur := node.Right
	for cur != nil {
		this.stack = append(this.stack, cur)
		cur = cur.Left
	}
	return node.Val
}

func (this *BSTIterator) HasNext() bool {
	return len(this.stack) > 0
}

/**
 * Your BSTIterator object will be instantiated and called as such:
 * obj := Constructor(root);
 * param_1 := obj.Next();
 * param_2 := obj.HasNext();
 */
// @lc code=end

