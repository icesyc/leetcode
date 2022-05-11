/*
 * @lc app=leetcode id=34 lang=golang
 *
 * [34] Find First and Last Position of Element in Sorted Array
 */

// @lc code=start
func searchRange(nums []int, target int) []int {
	low, high := 0, len(nums)-1
	res := []int{-1, -1}
	///查找最左
	for low < high {
		mid := (low + high) / 2
		if nums[mid] < target {
			low = mid + 1
		} else {
			high = mid
		}
	}
	//没找到
	if len(nums) == 0 || nums[low] != target {
		return res
	}
	res[0] = low
	//查找最右
	high = len(nums) - 1
	for low < high {
		//mid base to the right
		mid := (low+high)/2 + 1
		if nums[mid] > target {
			high = mid - 1
		} else {
			low = mid
		}
	}
	res[1] = high
	return res
}

// @lc code=end

