/*
 * @lc app=leetcode id=72 lang=golang
 *
 * [72] Edit Distance
 */
/**
 * dp[i][j]为从word1[0:i+1] -> word2[0:j+1]的编辑距离
 * 考虑几种情况
 * word1[i] == word2[j]，当前字符相等，那编辑距离就是dp[i-1][j-1]
 * word1[i] != word[j], 就要取min(dp[i-1][j-1],dp[i][j-1],dp[i-1][j])
 * 举例说明 abcd -> azc, i=3, j=2
 * 我们看从abcd->azc的编辑距离，当前d!=c,那么可以分为以下三种
 * 已知 abcd->az (dp[i][j-1])  需要3步，那么从abcd->azc就需要添加一个c,为4步[insert]
 * 已知 abc->azc (dp[i-1][j]) 需要1步, 那么abcd->azc就需要删除一个d,为2步 [delete]
 * 已知 abc->az (dp[i-1][j-1]) 需要2步，那么abcd->azc就需要将d替换为c, 为3步 [replace]
 * 由上可以，我们从abcd->azc最小距离为2步
 */
// @lc code=start
func minDistance(word1 string, word2 string) int {
	len1, len2 := len(word1), len(word2)
	dp := make([][]int, len1+1)
	for i := 0; i <= len1; i++ {
		dp[i] = make([]int, len2+1)
		dp[i][0] = i
	}
	for j := 0; j <= len2; j++ {
		dp[0][j] = j
	}
	for i := 1; i <= len1; i++ {
		for j := 1; j <= len2; j++ {
			if word1[i-1] == word2[j-1] {
				dp[i][j] = dp[i-1][j-1]
			} else {
				dp[i][j] = min(dp[i-1][j-1], dp[i][j-1], dp[i-1][j]) + 1
			}
		}
	}
	return dp[len1][len2]
}

func min(args ...int) int {
	m := args[0]
	for i := 1; i < len(args); i++ {
		if args[i] < m {
			m = args[i]
		}
	}
	return m
}

// @lc code=end

