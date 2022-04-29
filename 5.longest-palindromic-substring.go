/*
 * @lc app=leetcode id=5 lang=golang
 *
 * [5] Longest Palindromic Substring
 */

// @lc code=start
/**
 * Manacher 算法
 **/
func longestPalindrome(s string) string {
	padStr := padString(s)
	length := len(padStr)
	dp := make([]int, length)
	center, maxRight, begin, maxLength := 0, 0, 0, 1
	for i := 0; i < length; i++ {
		if i < maxRight {
			dp[i] = min(maxRight-i, dp[2*center-i])
		}

		//从i向两端扩展， 已经知道此位置有dp[i]个回文，偏移要减去该数量
		left, right := i-dp[i]-1, i+dp[i]+1
		for left >= 0 && right < length && padStr[left] == padStr[right] {
			dp[i]++
			left--
			right++
		}

		if i+dp[i] > maxRight {
			maxRight = i + dp[i]
			center = i
		}

		if dp[i] > maxLength {
			maxLength = dp[i]
			begin = (i - dp[i]) / 2
		}
	}
	return s[begin : begin+maxLength]
}
func min(a int, b int) int {
	if a < b {
		return a
	}
	return b
}
func padString(s string) string {
	padLen := len(s)*2 + 1
	padStr := make([]rune, padLen)
	j := 0
	for i := 0; j < len(s); i += 2 {
		padStr[i] = '#'
		padStr[i+1] = rune(s[j])
		j++
	}
	padStr[padLen-1] = '#'
	return string(padStr)
}

//*/
/**
 * 动态规划
 **
func longestPalindrome(s string) string {
	length := len(s)
	if length < 2 {
		return s
	}
	dp := make([][]bool, length)
	for i := 0; i < length; i++ {
		dp[i] = make([]bool, length)
		dp[i][i] = true
	}
	var begin, maxLength int = 0, 1
	for j := 1; j < length; j++ {
		for i := 0; i < j; i++ {
			if s[i] != s[j] {
				dp[i][j] = false
			} else if j-i < 3 {
				dp[i][j] = true
			} else {
				dp[i][j] = dp[i+1][j-1]
			}
			if dp[i][j] && j-i+1 > maxLength {
				maxLength = j - i + 1
				begin = i
			}
		}
	}
	return s[begin : begin+maxLength]
}

//*/
/**
 * 中心扩散法
 **
func longestPalindrome(s string) string {
	var from, max int
	for i, _ := range s {
		from, max = extendPalindrome(s, i-1, i+1, from, max)
		from, max = extendPalindrome(s, i, i+1, from, max)
	}
	return s[from : from+max]
}

func extendPalindrome(s string, i, j, from, max int) (int, int) {
	for i >= 0 && j < len(s) && s[i] == s[j] {
		i--
		j++
	}
	length := j - (i + 1)
	if length > max {
		return i + 1, length
	}
	return from, max

}
//*/

// @lc code=end

