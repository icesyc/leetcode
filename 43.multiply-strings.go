/*
 * @lc app=leetcode id=43 lang=golang
 *
 * [43] Multiply Strings
 */

// @lc code=start
func multiply(num1 string, num2 string) string {
	pos := make([]int, len(num1)+len(num2))
	for i := len(num1) - 1; i >= 0; i-- {
		for j := len(num2) - 1; j >= 0; j-- {
			p1, p2 := i+j, i+j+1
			mul := int(num1[i]-'0') * int(num2[j]-'0')
			sum := mul + pos[p2]
			pos[p1] += sum / 10
			pos[p2] = sum % 10
		}
	}

	res := []byte{}
	for i := 0; i < len(pos); i++ {
		if len(res) == 0 && pos[i] == 0 {
			continue
		}
		res = append(res, byte(pos[i])+'0')
	}
	if len(res) == 0 {
		return "0"
	}
	return string(res)
}

// @lc code=end

