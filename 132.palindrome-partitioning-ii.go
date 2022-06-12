/*
 * @lc app=leetcode id=132 lang=golang
 *
 * [132] Palindrome Partitioning II
 */

// @lc code=start
/**
 * 字符串的回文分割最小分割次数
 * 定义cut[j]为字符串0..j的最小分割次数
 * 如果字符串s[i..j]是回文，那么
 * cut[j] = min(cut[0..i-1]) + 1
 * 例如
 * a    b    a    |    c    c
 *                     i    j
 *          i-1   |    s[i,j] 是回文
 *      cut[i-1]  +   1
 */
func minCut(s string) int {
	m := len(s)
	dp := make([][]bool, m)
	cut := make([]int, m)
	for i := 0; i < m; i++ {
		dp[i] = make([]bool, m)
	}
	for j := 0; j < m; j++ {
		cut[j] = j
		for i := 0; i <= j; i++ {
			if s[i] == s[j] && (j-i < 3 || dp[i+1][j-1]) {
				dp[i][j] = true
				if i == 0 {
					cut[j] = 0
				} else {
					cut[j] = min(cut[j], cut[i-1]+1)
				}
			}
		}
	}
	return cut[m-1]
}
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// @lc code=end

