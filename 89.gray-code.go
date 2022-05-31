/*
 * @lc app=leetcode id=89 lang=golang
 *
 * [89] Gray Code
 */
/**
 * n=1 00 01
 * n=2 00   01  11  10
 * n=3 000 001 011 010 110 111 101 100
 * f(n) = f(n-1) + f(n-1)|2^(n-1)
 */
// @lc code=start
func grayCode(n int) []int {
	res := []int{0}
	for i := 1; i <= n; i++ {
		base := 1 << (i - 1)
		for j := len(res) - 1; j >= 0; j-- {
			res = append(res, base+res[j])
		}
	}
	return res
}

// @lc code=end

