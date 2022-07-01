/*
 * @lc app=leetcode id=213 lang=golang
 *
 * [213] House Robber II
 */

// @lc code=start
/**
 * 198题的变种
 * 对于一间房子，有2种选择，抢或不抢
 * 不抢 dp[i] = dp[i-1]
 * 抢 dp[i] = dp[i-2] + nums[i]
 * dp[i] = max(dp[i-1], dp[i-2] + nums[i])
 * 由于有环，问题变成需要处理2种情况
 * - 不抢1，抢n-1
 * - 抢1, 不抢n-1
 * 我们分别计算nums[1..n-1] nums[0..n-2]和dp结果，取最大值
 */
func rob(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}
	return max(helper(nums[1:]), helper(nums[:len(nums)-1]))
}

func helper(nums []int) int {
	n := len(nums)
		dp := make([]int, n+1)
		dp[0], dp[1] = 0, nums[0]
		for i := 2; i <= n; i++ {
			dp[i] = max(dp[i-1], dp[i-2]+nums[i-1])
		}
		return dp[n]
	}
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// @lc code=end

