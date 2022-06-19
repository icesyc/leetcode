/*
 * @lc app=leetcode id=172 lang=golang
 *
 * [172] Factorial Trailing Zeroes
 */

// @lc code=start
/**
 * 2*5=10
 * 可以认为每X5的倍数都会多1个0
 * 5的指数可提供多个0，5^2可提供2个0, 5^3可提供3个0
 * 算n包含多少个5的倍数
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

