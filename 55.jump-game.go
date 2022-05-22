/*
 * @lc app=leetcode id=55 lang=golang
 *
 * [55] Jump Game
 */

// @lc code=start
func canJump(nums []int) bool {
	curMax := 0
	for i := 0; i < len(nums)-1 && i <= curMax; i++ {
		if i+nums[i] > curMax {
			curMax = i + nums[i]
		}
	}
	return curMax >= len(nums)-1
}

// @lc code=end

