/*
 * @lc app=leetcode id=214 lang=golang
 *
 * [214] Shortest Palindrome
 */
// @lc code=start
/**
 * 使用kmp的Partical Match Table来求解
 *
 * PMT的定义
 * pmt[i]为字符串s[0..i]的最长公共前后缀的长度
 * 例如 a b d a b c d 对应的pmt为
 *     0 0 0 1 2 0 0
 * 知道定义后，关键问题在于我们如何快速求取pmt
 * 我们采用2个指针，一个指向首位置，一个指向第2个字符， 将i不断向后移动
 * 如果 s[i] == s[j] 那么说明存在公共前后缀，长度为1, 否则为0
 *
 * j     i
 * a b c a b a b
 * 0 0 0 1
 *
 * 这时我们向前移动j和i继续比较，发现s[i] == s[j] 那么 pmt[i] = 2 即 j+1
 *
 *   j     i
 * a b c a b a b
 * 0 0 0 1 2
 *
 * 如果s[i] != s[j](如下图)，我们需要减小j, 在保持A区间的前缀=B区间的后缀的前提下，尽可能让j最大
 * 由于A区间=B区间, 即查找A区间的最长公共前后缀，这个值已经计算过了，就是pmt[j-1]
 * 如下图pmt[j-1]为2，表示A区间和B区间的最长公共前后缀为2,那我们移到s[2]的位置，只要判断s[2]与s[i]是否相等
 * 然后继续判断s[i]==s[j]，直到j==0时pmt[i]即为0
 *
 *           j             i
 * a b c a b d d a b c a b c
 * 0 0 0 1 2
 * ----A----     -----B---
 *
 *     j                   i
 * a b c a b d d a b c a b c
 * 0 0 0 1 2
 * ----A----     -----B---
 *
 * 输入字符串s, 在前面增加最小的长度，使s变为回文，
 * 如果我们知道s[0..i]是回文，那么只需要将字符串s[i+1..n]反转添加到s的前面就得到答案了
 * 所以问题变为求字符串s从0开始的最长回文
 * 回文的特性是反转后的字符串与原串相等 s=reverse(s)
 * 令t = s + # + reverse(s)
 * 任务变成在t中查找最长公共前后缀(在s中的前缀=在reverse(s)中的后缀)，查找的结果就是最长回文前缀
 */
/*
func shortestPalindrome(s string) string {
	rs := reverse(s)
	t := s + "#" + rs
	pmt := buildPmt(t)
	maxPalindromePrefix := pmt[len(pmt)-1]
	return rs[:len(rs)-maxPalindromePrefix] + s
}

func buildPmt(s string) []int {
	pmt := make([]int, len(s))
	i, j := 1, 0
	for i < len(s) {
		if s[i] == s[j] {
			pmt[i] = j + 1
			i++
			j++
		} else if j > 0 {
			j = pmt[j-1]
		} else {
			pmt[i] = 0
			i++
		}
	}
	return pmt
}
*/
/**
 * 将字符串反转，判断rs[i..n]是否与s[0..n-i]相同
 * 即求字符串s的最长回文前缀
 */
func shortestPalindrome(s string) string {
	rs := reverse(s)
	for i := 0; i < len(rs); i++ {
		if rs[i:] == s[:len(s)-i] {
			return rs[:i] + s
		}
	}
	return rs + s
}
func reverse(s string) string {
	rs := []byte(s)
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		rs[i], rs[j] = rs[j], rs[i]
	}
	return string(rs)
}

// @lc code=end

