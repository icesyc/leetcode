/*
 * @lc app=leetcode id=128 lang=golang
 *
 * [128] Longest Consecutive Sequence
 */

// @lc code=start
/**
 * 最长连续数字序列
 * 先将nums保存到map中，循环map
 * 先找到起点n (n-1不在map中)
 * 不断循环n+1，直到n+1不在map
 */
func longestConsecutive(nums []int) int {
	numsMap := map[int]bool{}
	for _, n := range nums {
		numsMap[n] = true
	}
	maxLen := 0
	for n := range numsMap {
		if !numsMap[n-1] {
			end := n + 1
			for numsMap[end] {
				end++
			}
			if end-n > maxLen {
				maxLen = end - n
			}
		}
	}
	return maxLen
}

// @lc code=end

