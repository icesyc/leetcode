/*
 * @lc app=leetcode id=14 lang=golang
 *
 * [14] Longest Common Prefix
 */

// @lc code=start
/**
 * 解法2
 */
func longestCommonPrefix(strs []string) string {
	sort.Strings(strs)
	first, last, i := strs[0], strs[len(strs)-1], 0
	for i = 0; i < len(first); i++ {
		if first[i] != last[i] {
			break
		}
	}
	return first[0:i]
}

//*/

/**
 * 解法1
 *
func longestCommonPrefix(strs []string) string {
	res := ""
	for i := 0; i < len(strs[0]); i++ {
		for j := 1; j < len(strs); j++ {
			if i >= len(strs[j]) || strs[0][i] != strs[j][i] {
				return res
			}
		}
		res = strs[0][:i+1]
	}
	return res
}
//*/

// @lc code=end

