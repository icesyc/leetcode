/*
 * @lc app=leetcode id=168 lang=golang
 *
 * [168] Excel Sheet Column Title
 */

// @lc code=start
/**
 * 就是10进制转26进制，但是由于A代表的是1，
 * 而我们的A是从0开始的
 * 所以处理的时候需要取(columnNumber-1)
 * 进制转换的公式是 x = y%26^n... + y%26^1 + y%26^0
 */
func convertToTitle(columnNumber int) string {
	res := ""
	for columnNumber > 0 {
		res = string((columnNumber-1)%26+'A') + res
		columnNumber = (columnNumber - 1) / 26
	}
	return res
}

// @lc code=end

