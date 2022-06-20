/*
 * @lc app=leetcode id=191 lang=golang
 *
 * [191] Number of 1 Bits
 */

// @lc code=start
/**
 * n&n-1 可消除最右边的1
 * 循环计数
 */
func hammingWeight(num uint32) int {
	i := 0
	for num > 0 {
		num &= num - 1
		i++
	}
	return i
}

// @lc code=end

