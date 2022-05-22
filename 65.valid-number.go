/*
 * @lc app=leetcode id=65 lang=golang
 *
 * [65] Valid Number
 */

// @lc code=start
/**
 * 使用NFA, 要先定义好各状态的转换关系
 */
func isNumber(s string) bool {
	state := []map[string]int{
		map[string]int{},
		//q1
		map[string]int{"sign": 2, "digit": 3, ".": 4},
		//q2
		map[string]int{"digit": 3, ".": 4},
		//q3
		map[string]int{"digit": 3, ".": 5, "e": 6},
		//q4
		map[string]int{"digit": 5},
		//q5
		map[string]int{"digit": 5, "e": 6},
		//q6
		map[string]int{"sign": 7, "digit": 8},
		//q7
		map[string]int{"digit": 8},
		//q8
		map[string]int{"digit": 8},
	}
	curState := 1
	for i := 0; i < len(s); i++ {
		c := string(s[i])
		if s[i] >= '0' && s[i] <= '9' {
			c = "digit"
		} else if s[i] == '+' || s[i] == '-' {
			c = "sign"
		} else if s[i] == 'E' || s[i] == 'e' {
			c = "e"
		}
		ok := false
		if curState, ok = state[curState][c]; !ok {
			return false
		}
	}
	return curState == 3 || curState == 5 || curState == 8
}

/*
func isNumber(s string) bool {
	numberSeen, dotSeen, eSeen := false, false, false
	for i := 0; i < len(s); i++ {
		if s[i] >= '0' && s[i] <= '9' {
			numberSeen = true
		} else if s[i] == '.' {
			if dotSeen || eSeen {
				return false
			}
			dotSeen = true
		} else if s[i] == 'e' || s[i] == 'E' {
			if eSeen || !numberSeen {
				return false
			}
			numberSeen = false
			eSeen = true
		} else if s[i] == '+' || s[i] == '-' {
			if i != 0 && s[i-1] != 'e' && s[i-1] != 'E' {
				return false
			}
			numberSeen = false
		} else {
			return false
		}
	}
	return numberSeen
}
*/
// @lc code=end

