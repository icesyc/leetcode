/*
 * @lc app=leetcode id=91 lang=golang
 *
 * [91] Decode Ways
 */
/*
 * dp[0] = 0, dp[1] = 1
 * dp[i] = dp[i-1] if s[i-1] in[1,9]
 *       += dp[i-2] if s[i-2:i] in [10,26]
 * return dp[n]
 */
// @lc code=start
func numDecodings(s string) int {
	dp := make([]int, len(s)+1)
	dp[0] = 1
	if s[0] != '0' {
		dp[1] = 1
	}
	for i := 2; i <= len(s); i++ {
		second, _ := strconv.Atoi(s[i-2 : i])
		if s[i-1] != '0' {
			dp[i] += dp[i-1]
		}
		if second >= 10 && second <= 26 {
			dp[i] += dp[i-2]
		}
	}
	return dp[len(dp)-1]
}

// @lc code=end

