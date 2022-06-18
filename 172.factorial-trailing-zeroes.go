/*
 * @lc app=leetcode id=172 lang=golang
 *
 * [172] Factorial Trailing Zeroes
 */

// @lc code=start
/**
 * 算n包含多少个5
 */
func trailingZeroes(n int) int {
	cnt := 0
	for n > 0 {
		cnt += n / 5
		n /= 5
	}
	return cnt
}

// @lc code=end

