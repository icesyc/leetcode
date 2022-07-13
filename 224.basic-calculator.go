/*
 * @lc app=leetcode id=224 lang=golang
 *
 * [224] Basic Calculator
 */

// @lc code=start
/**
 * 循环处理字符串,考虑以下几种情况
 * 1. 遇到数字时，将解析的数字保存到number中
 * 2. 遇到+时，表示后续的操作是加法，将当前符号位计为1, 之前的结果进行累加,重置number
 * 3. 遇到-时，表示后续的作品是减法, 将当前符号位计为-1, 同样累加结果, 重置number
 * 4. 遇到(时，将结果push到stack上，重置number,result,sign,开始新一轮计算
 * 5. 遇到)时，将stack上的结果pop出现，与当前结果合并
 */
func calculate(s string) int {
	stack := []int{}
	//最开始的sign为1
	number, result, sign := 0, 0, 1
	for i := 0; i < len(s); i++ {
		if s[i] >= '0' && s[i] <= '9' {
			number = number*10 + int(s[i]-'0')
		} else if s[i] == '+' {
			result += sign * number
			number, sign = 0, 1
		} else if s[i] == '-' {
			result += sign * number
			number, sign = 0, -1
		} else if s[i] == '(' {
			stack = append(stack, result, sign)
			result, number, sign = 0, 0, 1
		} else if s[i] == ')' {
			result += sign * number
			stackNum, stackSign := stack[len(stack)-2], stack[len(stack)-1]
			stack = stack[:len(stack)-2]
			//合并栈上的结果
			result *= stackSign
			result += stackNum
			number = 0
		}
	}
	//处理最后的数字
	if number > 0 {
		result += sign * number
	}
	return result
}

// @lc code=end

