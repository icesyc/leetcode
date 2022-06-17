/*
 * @lc app=leetcode id=162 lang=golang
 *
 * [162] Find Peak Element
 */

// @lc code=start
/**
 * 二分法
 * 如果nums[mid] < nums[mid+1], 说明峰值在mid+1..n中
 * 否则峰值在0..mid中
 * 当left=right时，我们找到结果
 */
func findPeakElement(nums []int) int {
	left, right := 0, len(nums)-1
	for left < right {
		mid := (left + right) / 2
		//右边有峰值
		if nums[mid] < nums[mid+1] {
			left = mid + 1
		} else {
			//左边有峰值
			right = mid
		}
	}
	return left
}

// @lc code=end

