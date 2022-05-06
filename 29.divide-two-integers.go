/*
 * @lc app=leetcode id=29 lang=golang
 *
 * [29] Divide Two Integers
 */

// @lc code=start
func divide(dividend int, divisor int) int {
	positive := true
	if dividend < 0 == divisor < 0 {
		dividend = -dividend
		positive = !positive
	}
	if divisor < 0 {
		divisor = -divisor
		positive = !positive
	}
	temp := divisor
	added := 1
	for dividend >= divisor {
		dividend -= temp
		n += added
		temp = temp << 1
		added = added << 1
	}
	dividend += temp
	temp = temp >> 1
	for dividend >= divisor {
		temp = temp >> 1

	}
	if !positive {
		n = -n
	}
	if n > (1<<31)-1 {
		return (1 << 31) - 1
	}
	return n
}

// @lc code=end

