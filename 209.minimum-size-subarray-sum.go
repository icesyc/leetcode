/*
 * @lc app=leetcode id=209 lang=golang
 *
 * [209] Minimum Size Subarray Sum
 */

// @lc code=start
/**
 * slide window方法
 * 右指针不断向右移动，直到累计的结果>=target
 * 然后向前移动左指针缩小窗口，直到累计的结果<target
 * 重复上面的逻辑，直到逻辑结果，记录其中最小的满足条件的
 * minLen = right-left
 */
func minSubArrayLen(target int, nums []int) int {
	left, right := 0, 0
	minLen := math.MaxInt
	total := 0
	for right < len(nums) {
		total += nums[right]
		right++
		for total >= target {
			minLen = min(minLen, right-left)
			total -= nums[left]
			left++
		}
	}
	if minLen == math.MaxInt {
		return 0
	}
	return minLen
}
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// @lc code=end

