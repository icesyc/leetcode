/*
 * @lc app=leetcode id=27 lang=golang
 *
 * [27] Remove Element
 */

// @lc code=start
func removeElement(nums []int, val int) int {
	cnt, n := 0, len(nums)
	for i := 0; i < n; i++ {
		if nums[i] == val {
			cnt++
		} else {
			nums[i-cnt] = nums[i]
		}
	}
	return n - cnt
}

// @lc code=end

