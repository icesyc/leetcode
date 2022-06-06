/*
 * @lc app=leetcode id=120 lang=golang
 *
 * [120] Triangle
 */

// @lc code=start
/**
 * dp问题
 * 三角形的每一个点都可能与i,i+1相加
 * 我们使用bottom-up方式，从三角形底部开始遍历
 * dp[n][i]就是min(triangle[n][0..n])
 * 由于dp[n-1][i]只能与dp[n][i],dp[n][i+1]相加
 * 那么最小值的计算就是
 * dp[n-1][i] = min(dp[n][i], dp[n][i+1]) + triangle[n-1][i]
 * 最终求dp[0][0]
 * 由于每一层的dp只使用一次, 也可以将二维dp简化层一维
 * dp[i] = min(dp[i], dp[i+1]) + triangle[n-1][i]
 */
func minimumTotal(triangle [][]int) int {
	m := len(triangle)
	dp := append([]int{}, triangle[m-1]...)
	for i := m - 2; i >= 0; i-- {
		for j := 0; j < len(triangle[i]); j++ {
			dp[j] = min(dp[j], dp[j+1]) + triangle[i][j]
		}
	}
	return dp[0]
}
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// @lc code=end

