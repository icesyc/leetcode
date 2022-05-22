/*
 * @lc app=leetcode id=509 lang=golang
 *
 * [509] Fibonacci Number
 */

// @lc code=start
func fib(n int) int {
	if n < 2 {
		return n
	}
	prev, cur := 0, 1
	for i := 2; i <= n; i++ {
		sum := prev + cur
		prev = cur
		cur = sum
	}
	return cur
}

// @lc code=end

