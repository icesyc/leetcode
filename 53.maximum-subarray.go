/*
 * @lc app=leetcode id=53 lang=golang
 *
 * [53] Maximum Subarray
 */

// @lc code=start
func maxSubArray(nums []int) int {
	return divideAndConquer(nums)
}

func divideAndConquer(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}
	mid := len(nums) / 2
	ml, leftSum := -100000, 0
	for i := mid - 1; i >= 0; i-- {
		leftSum += nums[i]
		if leftSum > ml {
			ml = leftSum
		}
	}
	mr, rightSum := -100000, 0
	for i := mid; i < len(nums); i++ {
		rightSum += nums[i]
		if rightSum > mr {
			mr = rightSum
		}
	}
	maxLeft := divideAndConquer(nums[:mid])
	maxRight := divideAndConquer(nums[mid:])
	return max(max(maxLeft, maxRight), ml+mr)
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

/*
func maxSubArray(nums []int) int {
	curSum, maxSum := 0, -100000
	for i := 0; i < len(nums); i++ {
		if curSum > 0 {
			curSum = curSum + nums[i]
		} else {
			curSum = nums[i]
		}
		if curSum > maxSum {
			maxSum = curSum
		}
	}
	return maxSum
}
*/
/*
func maxSubArray(nums []int) int {
	dp := make([]int, len(nums))
	dp[0] = nums[0]
	max := dp[0]
	for i := 1; i < len(nums); i++ {
		if dp[i-1] > 0 {
			dp[i] = nums[i] + dp[i-1]
		} else {
			dp[i] = nums[i]
		}
		if dp[i] > max {
			max = dp[i]
		}
	}
	return max
}
*/
// @lc code=end

