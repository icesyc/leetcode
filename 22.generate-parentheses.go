/*
 * @lc app=leetcode id=22 lang=golang
 *
 * [22] Generate Parentheses
 */
// @lc code=start
func generateParenthesis(n int) []string {
	var res []string
	backTrace(n, n, "", &res)
	return res
}

func backTrace(left, right int, cur string, res *[]string) {
	if left == 0 && right == 0 {
		*res = append(*res, cur)
		return
	}
	if left > 0 {
		backTrace(left-1, right, cur+"(", res)
	}
	if right > left {
		backTrace(left, right-1, cur+")", res)
	}
}

// @lc code=end

