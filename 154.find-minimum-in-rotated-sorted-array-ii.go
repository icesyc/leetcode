/*
 * @lc app=leetcode id=154 lang=golang
 *
 * [154] Find Minimum in Rotated Sorted Array II
 */

// @lc code=start
/**
 * 同153，数组值可以重复
 * 原理基本相同
 * 如果nums[mid] >nums[right], 说明nums[mid..right]无序, 最小值应该在[mid+1..right]范围内
 * 如果nums[mid] < nums[right], 说明nums[mid..right]有序，最小值应该在[left..mid]里
 * 相等的话，无法判断，因为nums[mid]=nums[right]，所以right-1是安全的
 */
//6,7,8,8,9
func findMin(nums []int) int {
	left, right := 0, len(nums)-1

	for left < right {
		mid := (left + right) / 2
		if nums[mid] > nums[right] {
			left = mid + 1
		} else if nums[mid] < nums[right] {
			right = mid
		} else {
			right--
		}
	}
	return nums[left]
}

// @lc code=end

