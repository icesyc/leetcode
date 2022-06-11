/*
 * @lc app=leetcode id=121 lang=golang
 *
 * [121] Best Time to Buy and Sell Stock
 */

// @lc code=start
/**
 * 循环prices, 找到当前最小的价格,买入
 * 然后计算每天的价格和minPrice的差价最大值
 */
func maxProfit(prices []int) int {
	minPrice := math.MaxInt
	max := 0
	for i := 0; i < len(prices); i++ {
		if prices[i] < minPrice {
			minPrice = prices[i]
		} else {
			profit := prices[i] - minPrice
			if profit > max {
				max = profit
			}
		}
	}
	return max
}

// @lc code=end

