/*
 * @lc app=leetcode id=50 lang=golang
 *
 * [50] Pow(x, n)
 */

//解法同29题，将指数不断*2，并累积结果
//如果超过，则保存当前结果，将指数剩余部分继续做上面的运算，再将结果*当前的运算结果
// @lc code=start
func myPow(x float64, n int) float64 {
	if n == 0 {
		return 1
	}
	m := n
	if n < 0 {
		m = -n
	}
	var res float64 = 1
	for m > 0 {
		temp, tm := x, 1
		for (tm << 1) < m {
			temp *= temp
			tm = tm << 1
		}
		res *= temp
		m -= tm
	}
	if n < 0 {
		return 1 / res
	}
	return res
}

// @lc code=end

