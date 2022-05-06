/*
 * @lc app=leetcode id=28 lang=golang
 *
 * [28] Implement strStr()
 */

// @lc code=start
func strStr(haystack string, needle string) int {
	n1, n2 := len(haystack), len(needle)
	for i := 0; ; i++ {
		for j := 0; ; j++ {
			if j == n2 {
				return i
			}
			if i+j == n1 {
				return -1
			}
			if haystack[i+j] != needle[j] {
				break
			}
		}
	}
}

/*
func strStr(haystack string, needle string) int {
	i, j, n1, n2 := 0, 0, len(haystack), len(needle)

	if n2 == 0 {
		return 0
	}
	begin := -1
	for i = 0; i < n1; i++ {
		if haystack[i] == needle[j] {
			begin = i
			for i < n1 && j < n2 && haystack[i] == needle[j] {
				i++
				j++
			}
			if j == n2 {
				return begin
			}
			if i == n1 {
				return -1
			}
			i = begin
			j = 0
		}
	}
	return -1
}
*/
// @lc code=end

