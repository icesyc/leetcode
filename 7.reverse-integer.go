/*
 * @lc app=leetcode id=7 lang=golang
 *
 * [7] Reverse Integer
 */

// @lc code=start
func reverse(x int) int {
	r := 0
	for x != 0 {
		r = r*10 + x%10
		x = x / 10
	}
	if r > 1<<31-1 || r < -(1<<31) {
		return 0
	}
	return r
}

// @lc code=end

