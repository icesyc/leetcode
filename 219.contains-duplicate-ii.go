/*
 * @lc app=leetcode id=219 lang=golang
 *
 * [219] Contains Duplicate II
 */

// @lc code=start
/**
 * 使用hash-table保存数字的index
 * 如果hash-table中存在，取出index并与当前index进行比较
 */
func containsNearbyDuplicate(nums []int, k int) bool {
	idxMap := map[int]int{}
	for i, v := range nums {
		if idx, ok := idxMap[v]; ok && i-idx <= k {
			return true
		}
		idxMap[v] = i
	}
	return false
}

// @lc code=end

