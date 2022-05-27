/*
 * @lc app=leetcode id=81 lang=golang
 *
 * [81] Search in Rotated Sorted Array II
 */

// @lc code=start
func search(nums []int, target int) bool {
	low, high := 0, len(nums)-1

	for low <= high {
		//去除重复元素
		for low < high && nums[low] == nums[low+1] {
			low++
		}
		for low < high && nums[high] == nums[high-1] {
			high--
		}
		mid := (low + high) / 2
		if nums[mid] == target {
			return true
		}
		//左侧为有序
		if nums[low] <= nums[mid] {
			if target >= nums[low] && target < nums[mid] {
				high = mid - 1
			} else {
				low = mid + 1
			}
		} else {
			//右侧为有序
			if target > nums[mid] && target <= nums[high] {
				low = mid + 1
			} else {
				high = mid - 1
			}
		}
	}
	return false
}

// @lc code=end

