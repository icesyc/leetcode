/*
 * @lc app=leetcode id=122 lang=golang
 *
 * [122] Best Time to Buy and Sell Stock II
 */

// @lc code=start
/**
 * 可以卖任意多次，那么只要第i天价格比i-1天高，就可以卖掉
 * 利润就是sum(profit[i])
 */
func maxProfit(prices []int) int {
	totalProfit := 0
	for i := 1; i < len(prices); i++ {
		if prices[i] > prices[i-1] {
			totalProfit += prices[i] - prices[i-1]
		}
	}
	return totalProfit
}

// @lc code=end

