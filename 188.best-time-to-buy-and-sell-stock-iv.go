/*
 * @lc app=leetcode id=188 lang=golang
 *
 * [188] Best Time to Buy and Sell Stock IV
 */

// @lc code=start
/**
 * 同123题
 * 定义dp[k][i]为k次交易在第i天所获得的最大利润
 * 有两种情况
 * 1. 第i天不交易  dp[k][i] = dp[k][i-1]
 * 2. 第i天交易 假设第j天买入，那么第i天获取的最大利润为
 * dp[k][i] = max(dp[k][i-1], max(dp[k-1][j-1] + price[i] - price[j])) j=0..i-1
 * 由公式可知 price[i]为固定值，关键就是计算max(dp[k-1][j-1]-price[j])
 */
func maxProfit(k int, prices []int) int {
	if len(prices) == 0 {
		return 0
	}
	dp := make([][]int, k+1)
	dp[0] = make([]int, len(prices))
	for i := 1; i <= k; i++ {
		dp[i] = make([]int, len(prices))
		//初始的curMax由于dp[i-1][j-1]为0，所以就是-price[0]
		curMax := 0 - prices[0]
		for j := 1; j < len(prices); j++ {
			curMax = max(curMax, dp[i-1][j-1]-prices[j])
			dp[i][j] = max(dp[i][j-1], prices[j]+curMax)
		}
	}
	return dp[k][len(prices)-1]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// @lc code=end

