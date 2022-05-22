/*
 * @lc app=leetcode id=58 lang=golang
 *
 * [58] Length of Last Word
 */

// @lc code=start
func lengthOfLastWord(s string) int {
	ls := len(s)
	begin, end := 0, -1
	for i := ls - 1; i >= 0; i-- {
		if s[i] != ' ' && end == -1 {
			end = i
		}
		if s[i] == ' ' && end != -1 {
			begin = i + 1
			break
		}
	}
	return end - begin + 1
}

// @lc code=end

