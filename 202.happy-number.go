/*
 * @lc app=leetcode id=202 lang=golang
 *
 * [202] Happy Number
 */

// @lc code=start
/**
 * 使用hashmap保存计算过的结果，如果再次碰到，就返回false
 */
func isHappy(n int) bool {
	computed := map[int]int{}
	for n > 1 {
		if _, ok := computed[n]; ok {
			return false
		}
		res, m := 0, n
		for m > 0 {
			d := m % 10
			res += d * d
			m /= 10
		}
		computed[n] = res
		n = res
	}
	return true
}

// @lc code=end

