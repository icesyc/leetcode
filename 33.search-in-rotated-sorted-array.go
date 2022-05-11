/*
 * @lc app=leetcode id=33 lang=golang
 *
 * [33] Search in Rotated Sorted Array
 */

/**
 * 二分之后的数组，一定是一边有序的
 * 如果左边有序，那么值如果在这个区间，就移动高位到mid-1处;
 * 如果不在这个区间，那肯定在右侧，移动低位到mid+1处
 * 右侧同理
 */
// @lc code=start
func search(nums []int, target int) int {
	low := 0
	high := len(nums) - 1
	for low <= high {
		mid := (low + high) / 2
		if nums[mid] == target {
			return mid
		}
		//左侧为有序
		if nums[low] <= nums[mid] {
			if target >= nums[low] && target < nums[mid] {
				high = mid - 1
			} else {
				low = mid + 1
			}
		} else {
			if target > nums[mid] && target <= nums[high] {
				low = mid + 1
			} else {
				high = mid - 1
			}
		}
	}
	return -1
}

// @lc code=end

