/*
 * @lc app=leetcode id=10 lang=golang
 *
 * [10] Regular Expression Matching
 */

// @lc code=start
func isMatch(s string, p string) bool {
	i, j, lens, lenp := 0, 0, len(s), len(p)
	for j = 0; j < lenp; i, j = i+1, j+1 {
		if j+1 < lenp && p[j+1] == '*' {
			matchStart := i
			for i < lens && (p[j] == '.' || s[i] == p[j]) {
				i++
			}
			j += 2
			//模式扫描结束(a*$), 检查字符串是否扫描结束
			if j == lenp {
				return i == lens
			}
			//采用回溯法继续匹配后面的
			s2 := ""
			for i >= matchStart {
				if i < lens {
					s2 = s[i:lens]
				}
				if isMatch(s2, p[j:lenp]) {
					return true
				}
				i--
			}
			return false
		}
		if i == lens || p[j] != '.' && s[i] != p[j] {
			return false
		}
	}
	return i == lens
}

// @lc code=end

