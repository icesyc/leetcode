/*
 * @lc app=leetcode id=38 lang=golang
 *
 * [38] Count and Say
 */

// @lc code=start
func countAndSay(n int) string {
	if n == 1 {
		return "1"
	}
	str := countAndSay(n - 1)

	count := 1
	res := ""
	for i := 1; i < len(str); i++ {
		if str[i] != str[i-1] {
			res += fmt.Sprintf("%d%c", count, str[i-1])
			count = 0
		}
		count++
	}
	res += fmt.Sprintf("%d%c", count, str[len(str)-1])
	return res
}

// @lc code=end

