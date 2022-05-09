/*
 * @lc app=leetcode id=29 lang=golang
 *
 * [29] Divide Two Integers
 */

// @lc code=start
func divide(dividend int, divisor int) int {
	if dividend == -1<<31 && divisor == -1 {
		return (1 << 31) - 1
	}
	positive := true
	if dividend < 0 {
		dividend = -dividend
		positive = !positive
	}
	if divisor < 0 {
		divisor = -divisor
		positive = !positive
	}

	res := 0
	for dividend >= divisor {
		temp, m := divisor, 1
		for temp<<1 <= dividend {
			temp = temp << 1
			m = m << 1
		}
		dividend -= temp
		res += m
	}
	if positive {
		return res
	}
	return -res
}

// @lc code=end

