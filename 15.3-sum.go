/*
 * @lc app=leetcode id=15 lang=golang
 *
 * [15] 3Sum
 */

// @lc code=start
func threeSum(nums []int) [][]int {
	sort.Ints(nums)
	var res [][]int
	for i := 0; i < len(nums)-2; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		start, end := i+1, len(nums)-1
		for start < end {
			if start > i+1 && nums[start] == nums[start-1] {
				start++
				continue
			}
			if end < len(nums)-1 && nums[end] == nums[end+1] {
				end--
				continue
			}
			sum := nums[i] + nums[start] + nums[end]
			if sum == 0 {
				grp := []int{nums[i], nums[start], nums[end]}
				res = append(res, grp)
				end--
			} else if sum > 0 {
				end--
			} else {
				start++
			}
		}
	}
	return res
}

// @lc code=end

