/*
 * @lc app=leetcode id=67 lang=golang
 *
 * [67] Add Binary
 */

// @lc code=start
func addBinary(a string, b string) string {
	la, lb := len(a), len(b)
	carry := byte(0)
	res := ""
	for i, j := la-1, lb-1; j >= 0 || i >= 0; i, j = i-1, j-1 {
		if i >= 0 {
			carry += a[i] - '0'
		}
		if j >= 0 {
			carry += b[j] - '0'
		}
		res = string(carry%2+'0') + res
		carry = carry / 2
	}
	if carry > 0 {
		res = "1" + res
	}
	return res
}

// @lc code=end

