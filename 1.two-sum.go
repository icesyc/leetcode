/*
 * @lc app=leetcode id=1 lang=golang
 *
 * [1] Two Sum
 */

// @lc code=start
func twoSum(nums []int, target int) []int {
	m := make(map[int]int)
	for k, a := range nums {
		b := target - a
		if j, ok := m[b]; ok {
			return []int{j, k}
		}
		m[a] = k
	}
	return nil
}

// @lc code=end

