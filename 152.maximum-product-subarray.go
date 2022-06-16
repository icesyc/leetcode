/*
 * @lc app=leetcode id=152 lang=golang
 *
 * [152] Maximum Product Subarray
 */

// @lc code=start
/**
 * 同53题类似，但这里我们需要两个dp数组
 * 一个保存最大值，一个保存最小值
 * 因为负负得正的原理，我们只保存一个值无法确定当前是否为最大值
 * 每次循环的时候需要从三个数中重新计算出最大和最小值
 */
func maxProduct(nums []int) int {
	minDp := make([]int, len(nums))
	maxDp := make([]int, len(nums))
	minDp[0], maxDp[0] = nums[0], nums[0]
	res := nums[0]
	for i := 1; i < len(nums); i++ {
		maxDp[i] = max(max(maxDp[i-1]*nums[i], minDp[i-1]*nums[i]), nums[i])
		minDp[i] = min(min(maxDp[i-1]*nums[i], minDp[i-1]*nums[i]), nums[i])
		res = max(res, maxDp[i])
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// @lc code=end

