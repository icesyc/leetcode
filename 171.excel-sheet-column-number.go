/*
 * @lc app=leetcode id=171 lang=golang
 *
 * [171] Excel Sheet Column Number
 */

// @lc code=start
/**
 * 第168题的反向操作
 * 还是注意一下转成10进制时需要+1的操作
 */
func titleToNumber(columnTitle string) int {
	n := 0
	for _, c := range columnTitle {
		n = n*26 + int(c-'A'+1)
	}
	return n
}

// @lc code=end

