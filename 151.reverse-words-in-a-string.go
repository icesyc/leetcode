/*
 * @lc app=leetcode id=151 lang=golang
 *
 * [151] Reverse Words in a String
 */

// @lc code=start
func reverseWords(s string) string {
	words := []string{}
	for i := 0; i < len(s); {
		j := i
		for j < len(s) && s[j] != ' ' {
			j++
		}
		if j > i {
			words = append([]string{s[i:j]}, words...)
		}
		i = j + 1
	}
	return strings.Join(words, " ")
}

// @lc code=end

