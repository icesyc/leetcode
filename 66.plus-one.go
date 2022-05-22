/*
 * @lc app=leetcode id=66 lang=golang
 *
 * [66] Plus One
 */

// @lc code=start
func plusOne(digits []int) []int {
	dl := len(digits)
	added := 1
	for i := dl - 1; i >= 0; i-- {
		sum := digits[i] + added
		added = sum / 10
		digits[i] = sum % 10
		if added == 0 {
			break
		}
	}
	res := digits
	if added > 0 {
		res = append([]int{added}, digits...)
	}
	return res
}

// @lc code=end

