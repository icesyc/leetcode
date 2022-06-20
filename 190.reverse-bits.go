/*
 * @lc app=leetcode id=190 lang=golang
 *
 * [190] Reverse Bits
 */

// @lc code=start
/**
 * 将num不断右移, 将结果不断左移再加上num的最右一位
 */
func reverseBits(num uint32) uint32 {
	var res uint32
	for i := 0; i < 32; i++ {
		res = (res << 1) | num&1
		num >>= 1
	}
	return res
}

// @lc code=end

