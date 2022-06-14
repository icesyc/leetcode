/*
 * @lc app=leetcode id=139 lang=golang
 *
 * [139] Word Break
 */

// @lc code=start
/**
 * dp解法
 * 定义dp[i]为字符地串s[0..i-1]是否可以拆分
 * base case dp[0]=true
 * 如果存在j=[0..i-1],dp[j]=true 并且s[j:i]在dict中, 那么dp[i]=true
 *
 * 我们从长度1开始，一直到n进行遍历
 * 内层查找j=[0..i-1]，看是否存在dp[j]=true && s[j:i]在dict中
 * 以例子leetcode [leet, code]为例
 * 很容易发现当i=4时dp[4]=true
 * 当i=8时,我们不断查找
 * ^|leetcode,l|eetcode,lee|tcode,leet|code
 * 发现dp[4]=true 并且s[4:8] = code 在dict中，所以dp[8]=true
 */
func wordBreak(s string, wordDict []string) bool {
	wordMap := map[string]bool{}
	for _, word := range wordDict {
		wordMap[word] = true
	}
	n := len(s)
	dp := make([]bool, n+1)
	dp[0] = true

	for i := 1; i <= n; i++ {
		for j := 0; j < i; j++ {
			word := s[j:i]
			if dp[j] && wordMap[word] {
				dp[i] = true
			}
		}
	}
	return dp[n]
}

// @lc code=end

