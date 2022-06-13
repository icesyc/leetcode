/*
 * @lc app=leetcode id=137 lang=golang
 *
 * [137] Single Number II
 */

// @lc code=start
/**
 * 同上一题，只不过每个元素出现的次数为3
 * 如果使用map，可保存每个元素的次数
 * 再循环set，返回次数为1的即可
 * 但题目要求O(1)的space，需要换一种思路
 * 如果某个数字出现3次，那么该数字所有的bit都会出现3次
 * 那么所有出现3次的数字的bit只能是3的倍数,3n
 * 只出现一次的数字，对应的bit就为3n+1
 * 如[2,2,3,2]
 * 0 1 0
 * 0 1 0
 * 0 1 1
 * 0 1 0
 * ----------
 * 0 4 1
 * 结果为4%3 <<1 + 1%3 << 0
 */
func singleNumber(nums []int) int {
	//注意这里要使用int32，否则会因为64位而变成超出32位正整数,
	//会出现bad case
	var res int32
	for i := 0; i < 32; i++ {
		bitCount := 0
		for _, n := range nums {
			if (n>>i)&1 == 1 {
				bitCount++
			}
		}
		if bitCount%3 == 1 {
			res += 1 << i
		}
	}
	return int(res)
}

// @lc code=end

