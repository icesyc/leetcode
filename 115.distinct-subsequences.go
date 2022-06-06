/*
 * @lc app=leetcode id=115 lang=golang
 *
 * [115] Distinct Subsequences
 */

/**
 *  DP解题
 * 对于字符串s和子序列t, 定义dp[i][j]为s[0:i]包含t[0:j]
 * 子序列的个数
 * 当s[i-1] != t[j-1] 时 dp[i][j] = dp[i-1][j]
 * 当s[i-1] == t[j-1]时 dp[i][j] = dp[i-1][j] + dp[i-1][j-1]
 * base case:
 * dp[i][0] = 1, dp[0][j] = 0
 * 举例说明
 * 假设s=baagd, t=bag，i, j指向如^所示
 *          ^      ^
 * d!=g 那么s包含t的个数就看s[baag]包含了几个t
 * 假设s=baagg, t=bag，i, j指向如^所示
 *          ^      ^
 * g=g, 那么s包含t的个数至少是s[baag]包含的个数t
 * 又因为s[i]是t的最后一个字符，那我们就看s[baag]包含了多少个
 * t的前缀[ba], 有多少个前缀再加上g就等于多了多少个t
 * 所以结论dp[i][j] = dp[i-1][j] + dp[i-1][j-1]
 */

// @lc code=start
func numDistinct(s string, t string) int {
	m, n := len(s), len(t)
	dp := make([][]int, m+1)
	dp[0] = make([]int, n+1)
	for i := 0; i <= m; i++ {
		dp[i] = make([]int, n+1)
		dp[i][0] = 1
	}
	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if s[i-1] == t[j-1] {
				dp[i][j] = dp[i-1][j] + dp[i-1][j-1]
			} else {
				dp[i][j] = dp[i-1][j]
			}
		}
	}
	return dp[m][n]
}

// @lc code=end

