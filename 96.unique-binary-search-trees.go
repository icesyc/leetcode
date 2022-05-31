/*
 * @lc app=leetcode id=96 lang=golang
 *
 * [96] Unique Binary Search Trees
 */
/**
 * 同95题，比95题要简单，不需要返回所有路径，只需要返回个数
 * DP[0] = 1
 * DP[1] = 1
 * DP[2] = Node[Val=1] Left in DP[0:1], Right in DP[2:2]
		   + Node[Val=2] Left in DP[1:2], Right in DP[3:2]
 * DP[i] = for j in [1:i] Node[Val=k, Left in DP[j-1], Right in DP[j+1:i]]
*/
// @lc code=start
func numTrees(n int) int {
	dp := make([]int, n+1)
	//计算dp1的时候需要leftNum*rightNum, 设置为0结果不正确
	dp[0] = 1
	for i := 1; i <= n; i++ {
		for j := 1; j <= i; j++ {
			//左树个数*右树个数
			dp[i] += dp[j-1] * dp[i-j]
		}
	}
	return dp[n]
}

// @lc code=end

