/*
 * @lc app=leetcode id=8 lang=golang
 *
 * [8] String to Integer (atoi)
 */

// @lc code=start
func myAtoi(s string) int {

	isLeading := true
	isNegative := false
	var res int
	for _, c := range s {
		if c == ' ' && isLeading {
			continue
		}
		if isLeading && c == '-' {
			isNegative = true
		} else if isLeading && c == '+' {
		} else if c > 47 && c < 58 {
			res = res*10 + int(c) - 48
		} else {
			break
		}
		isLeading = false
		if !isNegative && res > 1<<31-1 {
			return 1<<31 - 1
		} else if isNegative && res > 1<<31 {
			return -(1 << 31)
		}
	}
	if isNegative {
		res = -res
	}
	return res

}

// @lc code=end

