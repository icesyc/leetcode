/*
 * @lc app=leetcode id=123 lang=golang
 *
 * [123] Best Time to Buy and Sell Stock III
 */

// @lc code=start
/**
 * 买卖股票问题，最多买卖2次，用动态规划解决最多买卖K次的最大利润
 * DP[k][i]表示K次交易在在第i天所获取的最大利润
 * 那么有两种情况
 * 1.第i天不交易，那么DP[k][i] = DP[k][i-1]
 * 2.第i天卖出，那么DP[k][i] = max(DP[k][i-1], price[i]-price[j] + DP[k-1][j-1]) j=0..i-1
 * 总结公式
 * DP[k][i] = max(DP[k][i-1], price[i] + max(DP[k-1][j-1]-price[j])) j=0..i-1
 * 如果第i天卖出，我们买入的可能为第j天(j=0..i-1)
 * 第K次的交易利润就是price[i]-price[j]
 * 再加上之前的最大利润DP[k-1][j-1]就是能获取的最大利润
 * 因为是在第j天买入，那么之前的交易截止为第j-1天
 * 根据公式，关键的计算在于max(DP[k-1][j-1]-price[j])
 * 在循环价格时，当前max的值可以保存到变量里，只需要和当前循环的值比较即可
 * 防止重复计算
 */
func maxProfit(prices []int) int {
	return maxProfitK(prices, 2)
}

func maxProfitK(prices []int, maxK int) int {
	dp := make([][]int, maxK+1)
	dp[0] = make([]int, len(prices))
	for k := 1; k <= maxK; k++ {
		dp[k] = make([]int, len(prices))
		//初始的curMax由于dp[k-1][j-1]为0，所以就是-price[0]
		curMax := 0 - prices[0]
		for i := 1; i < len(prices); i++ {
			curMax = max(curMax, dp[k-1][i-1]-prices[i])
			dp[k][i] = max(dp[k][i-1], prices[i]+curMax)
		}
	}
	return dp[maxK][len(prices)-1]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// @lc code=end

