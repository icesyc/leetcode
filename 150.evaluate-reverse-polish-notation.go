/*
 * @lc app=leetcode id=150 lang=golang
 *
 * [150] Evaluate Reverse Polish Notation
 */

// @lc code=start
/**
 * 逆波兰表达式求值，使用stack保存操作数，遇到操作符弹出最后2个操作数，操作完成后再push到stack
 * 小技巧：使用map保存对应的操作函数，使代码更为整洁
 */
func evalRPN(tokens []string) int {
	stack := []int{}
	opMap := map[string]func(int, int) int{
		"+": func(a, b int) int { return a + b },
		"-": func(a, b int) int { return a - b },
		"*": func(a, b int) int { return a * b },
		"/": func(a, b int) int { return a / b },
	}
	for _, token := range tokens {
		if opMap[token] != nil {
			op1 := stack[len(stack)-2]
			op2 := stack[len(stack)-1]
			opResult := opMap[token](op1, op2)
			stack = stack[:len(stack)-2]
			stack = append(stack, opResult)
		} else {
			tk, _ := strconv.Atoi(token)
			stack = append(stack, tk)
		}
	}
	return stack[0]
}

// @lc code=end

