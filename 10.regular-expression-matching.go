/*
 * @lc app=leetcode id=10 lang=golang
 *
 * [10] Regular Expression Matching
 */

// @lc code=start
/**
 * DP解法
 */
func isMatch(s string, p string) bool {
	dp := make([][]bool, len(s)+1)
	for i, _ := range dp {
		dp[i] = make([]bool, len(p)+1)
	}
	dp[0][0] = true
	for j := 1; j <= len(p); j++ {
		if p[j-1] == '*' {
			dp[0][j] = dp[0][j-2]
		}
	}
	for i := 1; i <= len(s); i++ {
		for j := 1; j <= len(p); j++ {
			if s[i-1] == p[j-1] || p[j-1] == '.' {
				dp[i][j] = dp[i-1][j-1]
			} else if p[j-1] == '*' {
				dp[i][j] = dp[i][j-2]
				if s[i-1] == p[j-2] || p[j-2] == '.' {
					dp[i][j] = dp[i][j] || dp[i-1][j]
				}
			} else {
				dp[i][j] = false
			}
		}
	}
	return dp[len(s)][len(p)]
}

//*/
/**
 * 解法1
 *
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
//*/
// @lc code=end

