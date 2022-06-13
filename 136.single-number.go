/*
 * @lc app=leetcode id=136 lang=golang
 *
 * [136] Single Number
 */

// @lc code=start
/**
 * 利用map保存出现过的数字，再次出现时从map中删除
 * 最后返回map中唯一的值
 */
/*
func singleNumber(nums []int) int {
	set := map[int]bool{}
	for _, i := range nums {
		if !set[i] {
			set[i] = true
		} else {
			delete(set, i)
		}
	}
	res := -1
	for i := range set {
		res = i
	}

	return res
}
*/
/**
 * 利用xor的特性 n ^ n = 0
 * 那么将所有数字xor，最后返回的肯定是出现一次的数字
 */
func singleNumber(nums []int) int {
	res := 0
	for _, i := range nums {
		res ^= i
	}
	return res
}

// @lc code=end

