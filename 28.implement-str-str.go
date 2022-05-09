/*
 * @lc app=leetcode id=28 lang=golang
 *
 * [28] Implement strStr()
 */

// @lc code=start
/**
 * kmp
 */
func strStr(haystack string, needle string) int {
	next := buildNext(needle)
	i, j := 0, 0
	for i < len(haystack) {
		if haystack[i] == needle[j] {
			i++
			j++
		} else if j > 0 {
			j = next[j-1]
		} else {
			i++
		}
		if j == len(needle) {
			return i - j
		}
	}
	return -1
}

func buildNext(pattern string) []int {
	pl := len(pattern)
	next := make([]int, pl)
	i, j := 1, 0
	for i < pl {
		if pattern[i] == pattern[j] {
			next[i] = j + 1
			i++
			j++
		} else if j > 0 {
			j = next[j-1]
		} else {
			next[i] = 0
			i++
		}
	}
	return next
}

//*/
/*
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
*/
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

