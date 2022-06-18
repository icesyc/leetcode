/*
 * @lc app=leetcode id=169 lang=golang
 *
 * [169] Majority Element
 */

// @lc code=start
/**
 * hashmap计数法
 */
/*
func majorityElement(nums []int) int {
	dict := map[int]int{}
	n := len(nums)
	for _, i := range nums {
		dict[i]++
		if dict[i] > n/2 {
			return i
		}
	}
	return -1
}
*/
/**
 * moore voting(摩尔投票)
 * 每次从数组中删除两个不一样的数字
 * 最后剩下的数字就是majority
 */
func majorityElement(nums []int) int {
	majority, count := 0, 0
	for _, i := range nums {
		if count == 0 {
			majority = i
		}
		if i == majority {
			count++
		} else {
			count--
		}
	}
	return majority
}

// @lc code=end

