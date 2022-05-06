/*
 * @lc app=leetcode id=26 lang=golang
 *
 * [26] Remove Duplicates from Sorted Array
 */

// @lc code=start
func removeDuplicates(nums []int) int {
	cnt, n := 0, len(nums)
	for i := 1; i < n; i++ {
		if nums[i] == nums[i-1] {
			cnt++
		} else {
			nums[i-cnt] = nums[i]
		}
	}
	return n - cnt
}

/*
func removeDuplicates(nums []int) int {
	i, n := 0, len(nums)
	for j := i + 1; j < n; j++ {
		if nums[i] != nums[j] {
			nums[i+1] = nums[j]
			i++
		}
	}
	return i + 1
}
*/

// @lc code=end

