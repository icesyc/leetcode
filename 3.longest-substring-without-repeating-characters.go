/*
 * @lc app=leetcode id=3 lang=golang
 *
 * [3] Longest Substring Without Repeating Characters
 */

// @lc code=start
func lengthOfLongestSubstring(s string) int {
	idxMap := make(map[rune]int)
	var start, end, max int
	for i, c := range s {
		if idx, ok := idxMap[c]; ok && idx >= start {
			start = idx + 1
		}
		if end-start+1 > max {
			max = end - start + 1
		}
		idxMap[c] = i
		end++
	}
	return max

}

// @lc code=end

