/*
 * @lc app=leetcode id=62 lang=golang
 *
 * [62] Unique Paths
 */
/**
 * 动态规划
 * 假设start为终点，dp[i][j]表示从i,j的位置到0,0所有的路线
 * 可知道 dp[0][j] 和dp[[i][0]全部为1
 * dp[1][1] = dp[0][1] + dp[1][0]
 * dp[i][j] = dp[i-1][j] + dp[i][j-1]
 * 由于只能往上和往左走，dp[i][j]取决于dp[i][j-1]的路径数
 * 和dp[i-1][j]的路径数之和
 * 最终求dp[m-1][n-1]
 *
 */
// @lc code=start
func uniquePaths(m int, n int) int {
	dp := make([][]int, m)
	for i := 0; i < m; i++ {
		dp[i] = make([]int, n)
		for j := 0; j < n; j++ {
			if i == 0 || j == 0 {
				dp[i][j] = 1
			} else {
				dp[i][j] = dp[i-1][j] + dp[i][j-1]
			}
		}
	}
	return dp[m-1][n-1]
}

// @lc code=end

