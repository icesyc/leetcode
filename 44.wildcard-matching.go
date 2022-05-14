/*
 * @lc app=leetcode id=44 lang=golang
 *
 * [44] Wildcard Matching
 */

// @lc code=start
//dp解法，类似第10题
/**
 * 公式 dp[i][j]表示s[0:i]和p[0:j]的匹配关系
 * = dp[i-1][j-1] when s[i-1]==p[j-1] || p[j-1] == '?'
 * = dp[i][j-1] || dp[i-1][j] when p[j-1] == '*'
 * = false when else
 */
func isMatch(s string, p string) bool {
	ls, lp := len(s), len(p)
	dp := make([][]bool, ls+1)
	dp[0] = make([]bool, lp+1)
	dp[0][0] = true
	for j := 1; j <= lp; j++ {
		if p[j-1] == '*' {
			dp[0][j] = dp[0][j-1]
		}
	}
	for i := 1; i <= ls; i++ {
		dp[i] = make([]bool, lp+1)
		for j := 1; j <= lp; j++ {
			if s[i-1] == p[j-1] || p[j-1] == '?' {
				dp[i][j] = dp[i-1][j-1]
			} else if p[j-1] == '*' {
				dp[i][j] = dp[i][j-1] || dp[i-1][j]
			} else {
				dp[i][j] = false
			}
		}
	}
	return dp[ls][lp]
}

// @lc code=end

