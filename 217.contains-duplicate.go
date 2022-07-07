/*
 * @lc app=leetcode id=217 lang=golang
 *
 * [217] Contains Duplicate
 */

// @lc code=start
/**
 * 就是用hashmap保存一下所有出现过的数字，比较简单
 */
func containsDuplicate(nums []int) bool {
	hash := map[int]bool{}
	for i := 0; i < len(nums); i++ {
		if _, ok := hash[nums[i]]; ok {
			return true
		}
		hash[nums[i]] = true
	}
	return false
}

// @lc code=end

