/*
 * @lc app=leetcode id=198 lang=golang
 *
 * [198] House Robber
 */

// @lc code=start
/**
 * 定义dp[i]为抢劫第i间屋子时所得到最大钱数
 * 对于第i间屋子，有两种选择，抢或者不抢
 * dp[i] = max(dp[i-1], dp[i-2] + nums[i])
 * 为了便于计算，我们设置dp[0] = 0占位
 * base case dp[1] = nums[0]
 * 求dp[n]
 */
func rob(nums []int) int {
	dp := make([]int, len(nums)+1)
	dp[0], dp[1] = 0, nums[0]
	for i := 2; i <= len(nums); i++ {
		dp[i] = dp[i-2] + nums[i-1]
		if dp[i] < dp[i-1] {
			dp[i] = dp[i-1]
		}
	}
	return dp[len(nums)]
}

// @lc code=end

