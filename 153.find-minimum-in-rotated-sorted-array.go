/*
 * @lc app=leetcode id=153 lang=golang
 *
 * [153] Find Minimum in Rotated Sorted Array
 */
// @lc code=start
/**
 * 二分法搜索
 * 例如：2,3,4,5,1
 * 我们取mid=2, 值为4
 * 如果nums[mid] >nums[right]，最小值肯定在mid的右侧
 * 如果nums[mid] < nums[right], 最小值肯定在mid的左侧(包括mid)
 */
func findMin(nums []int) int {
	n := len(nums)
	left, right := 0, n-1
	//注意这里的条件不能是<=, 当left=right时，说明我们已经找到元素了
	//不然会死循环
	for left < right {
		mid := (left + right) / 2
		if nums[mid] > nums[right] {
			left = mid + 1
		} else {
			right = mid
		}
	}
	return nums[left]
}

// @lc code=end

