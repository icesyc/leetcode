/*
 * @lc app=leetcode id=201 lang=golang
 *
 * [201] Bitwise AND of Numbers Range
 */

// @lc code=start
/**
 * 求[left..right]范围内所有数字的bitAnd
 * 我们知道只有两个位为1，结果才为1
 * 假设m为10xxx n为11xxx
 * 那么m..n之间一定存在11000，由于后3位全是0
 * 所以m..n的bitAnd结果 bitAnd(m, n) == bitAnd(m, 11000)
 * 那么bitAnd(m, n) = 10000
 * 换句话说，本题求的是m,n的二进制最长公共前缀
 * 使用n&n-1可以消除最右侧的1，我们不断缩小n，直到n<=m
 * 这时的n即为最长公共前缀
 */
func rangeBitwiseAnd(left int, right int) int {
	for right > left {
		right &= right - 1
	}
	return right
}

// @lc code=end

