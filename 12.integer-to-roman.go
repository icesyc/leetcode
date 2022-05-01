/*
 * @lc app=leetcode id=12 lang=golang
 *
 * [12] Integer to Roman
 */

// @lc code=start
func intToRoman(num int) string {
	sym := []string{"M", "CM", "D", "CD", "C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"}
	nums := []int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}
	res, i := "", 0
	for num > 0 {
		for nums[i] > num {
			i++
		}
		num -= nums[i]
		res += sym[i]
	}
	return res
}

// @lc code=end

