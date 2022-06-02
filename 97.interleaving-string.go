/*
 * @lc app=leetcode id=97 lang=golang
 *
 * [97] Interleaving String
 */
/*
 * dp[i][j]表示字符串s1[0:i],s2[0:j]和s3[0:i+j]是交错关系
 * dp[i][j] = F(s1[0:i],s2[0:j],s3[0:i+j])
 * 那么有以下结论
 * dp[i][j] = s1[i] == s3[i+j] && dp[i-1][j]
 *         || s2[j] == s3[i+j] && dp[i][j-1]
 * 假设s1=abc s2=adf, s3 = abadcf
 * dp[3][3]是否成立 需要看s3的最后一个字符
 * f=s2[2]，我们移除s2的一个字符, 看dp[3][2]是否成立
 * 以此类推，如果s3[i+j]和s1[i],s2[j]都相等，那么
 * dp[i][j] = dp[i-1][j] || dp[i][j-1]
 */

// @lc code=start
func isInterleave(s1 string, s2 string, s3 string) bool {
	m, n := len(s1), len(s2)
	if m+n != len(s3) {
		return false
	}
	dp := make([][]bool, m+1)
	dp[0] = make([]bool, n+1)
	dp[0][0] = true
	for i := 1; i < m+1; i++ {
		dp[i] = make([]bool, n+1)
		dp[i][0] = dp[i-1][0] && s1[i-1] == s3[i-1]
	}
	for j := 1; j < n+1; j++ {
		dp[0][j] = dp[0][j-1] && s2[j-1] == s3[j-1]
	}
	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if s1[i-1] == s3[i+j-1] {
				dp[i][j] = dp[i-1][j]
			}
			if s2[j-1] == s3[i+j-1] {
				dp[i][j] = dp[i][j] || dp[i][j-1]
			}
		}
	}
	return dp[m][n]
}

// @lc code=end

