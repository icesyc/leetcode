/*
 * @lc app=leetcode id=17 lang=golang
 *
 * [17] Letter Combinations of a Phone Number
 */

// @lc code=start
/**
 * 解法2
 */
func letterCombinations(digits string) []string {
	letterMap := []string{
		"", "", "abc", "def", "ghi", "jkl",
		"mno", "pqrs", "tuv", "wxyz",
	}
	if len(digits) == 0 {
		return []string{}
	}
	result := []string{""}
	for i := 0; i < len(digits); i++ {
		n := digits[i]
		letters := letterMap[n-'0']
		tmp := make([]string, len(letters)*len(result))
		m := 0
		for j := 0; j < len(letters); j++ {
			for k := 0; k < len(result); k++ {
				tmp[m] = result[k] + string(letters[j])
				m++
			}
		}
		result = tmp
	}
	return result
}

//*/
/**
 * 解法1
 *
func letterCombinations(digits string) []string {
	var res []string
	if len(digits) == 0 {
		return []string{}
	}
	n := digits[0]
	start := 97 + (n-'2')*3
	if n > '7' {
		start += 1
	}
	charNum := 3
	if n == '7' || n == '9' {
		charNum = 4
	}
	for i := 0; i < charNum; i++ {
		char := string(start + byte(i))
		if len(digits) == 1 {
			res = append(res, char)
			continue
		}
		strs := letterCombinations(digits[1:])
		for j := 0; j < len(strs); j++ {
			res = append(res, char+strs[j])
		}
	}
	return res
}
//*/
// @lc code=end

