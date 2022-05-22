/*
 * @lc app=leetcode id=64 lang=golang
 *
 * [64] Minimum Path Sum
 */

/**
 * 与前一题类似，不过这次由于是求最小和
 * 我们需要dp[i][j]为止的最小和
 * dp[i][j] = min(dp[i-1][j], dp[i][j-1]) + grid[i][j]
 */
// @lc code=start
func minPathSum(grid [][]int) int {
	m, n := len(grid), len(grid[0])
	dp := make([]int, n)
	//先初始化第一行
	dp[0] = grid[0][0]
	for j := 1; j < n; j++ {
		dp[j] = dp[j-1] + grid[0][j]
	}
	for i := 1; i < m; i++ {
		//初始化第一列
		dp[0] += grid[i][0]
		for j := 1; j < n; j++ {
			dp[j] = min(dp[j], dp[j-1]) + grid[i][j]
		}
	}
	return dp[n-1]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// @lc code=end

