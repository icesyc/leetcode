/*
 * @lc app=leetcode id=167 lang=golang
 *
 * [167] Two Sum II - Input Array Is Sorted
 */
// @lc code=start
/**
 * 两侧双指针，根据和的大小
 * 不断缩小区间
 */
func twoSum(numbers []int, target int) []int {
	left, right := 0, len(numbers)-1
	for left < right {
		sum := numbers[left] + numbers[right]
		if sum > target {
			right--
		} else if sum < target {
			left++
		} else {
			break
		}
	}
	return []int{left + 1, right + 1}
}

// @lc code=end

