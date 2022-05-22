/*
 * @lc app=leetcode id=63 lang=golang
 *
 * [63] Unique Paths II
 */

/**
 * 同前一题，做了一下优化，dp[j]保存当前行的路径数
 * 根据上一题 dp[i][j] = 左侧的路径数+上面的路径数
 * dp[j] 只需要把之前的dp[j]+左侧的路径数
 * dp[j] += dp[j-1]
 * 同时需要判断一下障碍物，如果为，dp[j]=0
 *
 */
// @lc code=start
func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	m, n := len(obstacleGrid), len(obstacleGrid[0])
	dp := make([]int, n)
	dp[0] = 1
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if obstacleGrid[i][j] == 1 {
				dp[j] = 0
			} else if j > 0 {
				dp[j] += dp[j-1]
			}
		}
	}
	return dp[n-1]
}

// @lc code=end

