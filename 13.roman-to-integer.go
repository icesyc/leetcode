/*
 * @lc app=leetcode id=13 lang=golang
 *
 * [13] Roman to Integer
 */

// @lc code=start
func romanToInt(s string) int {
	symMap := map[string]int{
		"M": 1000, "D": 500,
		"C": 100, "L": 50,
		"X": 10, "V": 5, "I": 1,
	}
	res, i := 0, 0
	for i = 0; i < len(s); i++ {
		n := symMap[s[i:i+1]]
		if i+2 <= len(s) && n < symMap[s[i+1:i+2]] {
			res -= n
		} else {
			res += n
		}
	}
	return res
}

// @lc code=end

