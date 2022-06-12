/*
 * @lc app=leetcode id=131 lang=golang
 *
 * [131] Palindrome Partitioning
 */

// @lc code=start
/**
 * 字符串的回文分割组合
 * 思路，dp + backtracking
 * 根据第5题，可知计算字符串i..j是否是回文的方法
 * dp[i][j] = s[i] == s[j] && (j-i<3 || dp[i+1][j-1])
 * 使用backtracking的输入参数
 * s 字符串, begin, end 当前要判断回文的子串,
 * cur 当前处理过的子串回文集合, res 最终结果
 * 我们从s[0]开始检查
 * 发现一个回文，把该回文添加到cur中，并递归处理s[i+1..end]区间
 * 处理完成后，将cur回退一个长度，继续循环
 * 如果begin >= end 表示处理完成，我们把当前结果添加到最终结果并返回
 * tracking的顺序如下 以字符串abab为例
 * a,b,a,b -> a,b,ab ->a,ba,b ->a,bab
 * ab->a,b,-> ab,ab ->aba,b ->abab
 * 可见dp[i+1][j-1]会先处理,所以可以直接在backtracking中计算dp
 */
func partition(s string) [][]string {
	res := &[][]string{}
	dp := getDP(s)
	backtracking(s, dp, 0, len(s), []string{}, res)
	return *res
}

func backtracking(s string, dp [][]bool, start, end int, cur []string, res *[][]string) {
	if start >= end {
		*res = append(*res, append([]string{}, cur...))
		return
	}
	for i := start; i < end; i++ {
		if s[start] != s[i] {
			dp[start][i] = false
		} else if i-start < 3 {
			dp[start][i] = true
		} else {
			dp[start][i] = dp[start+1][i-1]
		}
		if dp[start][i] {
			cur = append(cur, s[start:i+1])
			backtracking(s, dp, i+1, end, cur, res)
			cur = cur[:len(cur)-1]
		}
	}
}

func getDP(s string) [][]bool {
	m := len(s)
	dp := make([][]bool, m)
	for i := 0; i < m; i++ {
		dp[i] = make([]bool, m)
		dp[i][i] = true
	}
	return dp
}

// @lc code=end

