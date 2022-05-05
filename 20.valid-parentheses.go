/*
 * @lc app=leetcode id=20 lang=golang
 *
 * [20] Valid Parentheses
 */

// @lc code=start
/**
 * 解法2
 */
func isValid(s string) bool {
	var stack []rune
	for _, c := range s {
		if c == '(' {
			stack = append(stack, ')')
		} else if c == '{' {
			stack = append(stack, '}')
		} else if c == '[' {
			stack = append(stack, ']')
		} else if len(stack) > 0 && stack[len(stack)-1] == c {
			stack = stack[0 : len(stack)-1]
		} else {
			return false
		}
	}
	return len(stack) == 0
}

//*/
/**
 * 解法1
 *
func isValid(s string) bool {
	var stack []rune
	pairs := map[rune]rune{
		')': '(', '}': '{', ']': '[',
	}
	for _, c := range s {
		if c == '(' || c == '{' || c == '[' {
			stack = append(stack, c)
		} else if p, ok := pairs[c]; ok {
			if len(stack) == 0 || stack[len(stack)-1] != p {
				return false
			}
			stack = stack[0 : len(stack)-1]
		}
	}
	return len(stack) == 0
}

//*/

// @lc code=end

