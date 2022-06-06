/*
 * @lc app=leetcode id=106 lang=golang
 *
 * [106] Construct Binary Tree from Inorder and Postorder Traversal
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
 * 给定的2个数组，一个是inorder,一个是postorder
 *           root
 *         /      \
 *      left1   right1
 *       /   \
 *    left2 right2
 * 根据inorder的定义，访问顺序为left2->left1->right2->root->right1...
 * 根据preorder的定义，访问顺序为root->left1->left2->right2->right1...
 * 根据postorder的定义，访问顺序为left2->right2->left1->right1->root...
 * 可知道postorder为preorder(先处理右child)的相反顺序
 * 同上一题原理相同
 * 处理的时候从inorder和postorder尾部处理即可
 * 递归即可
 *
 * 优化方案，增加一个map保存每个root所在位置, 不用每次遍历inorder
 * 根据postorder的定义，已知postorder的尾元素即root节点
 * 考虑以下输入inorder=[9,3,15,20,7], postorder=[9,15,7,20,3]
 * root=3, 我们在inorder搜索root时，根据map查找到rootIndex=1
 * right节点的长度rightLen = end - rootIndex
 * 那么root.Right=[15,20,7], root.Left=[9]
 * right=[rootIndex+1:end], left=[begin:rootIndex]
 * 那么构建root.Right的区间=[2:], postorder=postorder[:len(postorder)-1]
 * 构建root.Left的区间=[0:1], postorder=postorder[:len(postorder)-rightLen]
 * 递归处理即可, 结束条件为begin > end 说明没有子节点
 */
func buildTree(inorder []int, postorder []int) *TreeNode {
	inorderMap := map[int]int{}
	for i, v := range inorder {
		inorderMap[v] = i
	}
	return build(0, len(inorder)-1, inorderMap, postorder)
}
func build(begin, end int, inorderMap map[int]int, postorder []int) *TreeNode {
	if begin > end {
		return nil
	}
	root := &TreeNode{Val: postorder[len(postorder)-1]}
	rootIdx := inorderMap[root.Val]
	rightLen := end - rootIdx
	//减去root节点
	postorder = postorder[:len(postorder)-1]
	root.Right = build(rootIdx+1, end, inorderMap, postorder)
	//减去right节点的长度
	postorder = postorder[:len(postorder)-rightLen]
	root.Left = build(begin, rootIdx-1, inorderMap, postorder)
	return root
}

/*
func buildTree(inorder []int, postorder []int) *TreeNode {
	return build(inorder, postorder)
}
func build(inorder []int, postorder []int) *TreeNode {
	if len(postorder) == 0 {
		return nil
	}
	root := &TreeNode{Val: postorder[len(postorder)-1]}
	i := len(inorder) - 1
	for i >= 0 && inorder[i] != root.Val {
		i--
	}
	rightLen := len(inorder) - 1 - i
	rightStart := len(postorder) - 1 - rightLen
	root.Right = build(inorder[i+1:], postorder[rightStart:len(postorder)-1])
	root.Left = build(inorder[0:i], postorder[:rightStart])
	return root
}
*/
// @lc code=end

