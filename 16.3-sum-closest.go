/*
 * @lc app=leetcode id=16 lang=golang
 *
 * [16] 3Sum Closest
 */

// @lc code=start
func threeSumClosest(nums []int, target int) int {
	sort.Ints(nums)
	closestSum, diff := 0, math.MaxInt32
	for i := 0; i < len(nums)-2; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		start, end := i+1, len(nums)-1
		for start < end {
			sum := nums[i] + nums[start] + nums[end]
			if sum < target {
				start++
			} else {
				end--
			}
			newDiff := int(math.Abs(float64(sum - target)))
			if newDiff < diff {
				closestSum, diff = sum, newDiff
			}
		}
	}
	return closestSum
}

// @lc code=end

