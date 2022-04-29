/*
 * @lc app=leetcode id=6 lang=golang
 *
 * [6] Zigzag Conversion
 */

/**
 * 计算法
 * i
 * 0, i + 2n-2
 * 1, i + 2n-2-2, i + 2n-2
 * 2, i + 2n-2-4, i + 2n-2
 * 3, i + 2n-2-6, i + 2n-2
 */
func convert(s string, numRows int) string {
	if numRows == 1 {
		return s
	}
	strlen := len(s)
	loopNum := 2*numRows - 2
	var res string
	for i := 0; i < numRows; i++ {
		for j := i; j < strlen; j += loopNum {
			res += string(s[j])
			nextPos := j + 2*(numRows-i) - 2
			if nextPos%(numRows-1) > 0 && nextPos < strlen {
				res += string(s[nextPos])
			}
		}
	}
	return res
}

//*/
/**
 * 填充法
 *
func convert(s string, numRows int) string {
	strLen := len(s)
	rows := make([][]byte, numRows)

	var j, step int = 0, 1
	for i := 0; i < strLen; i++ {
		rows[j] = append(rows[j], s[i])
		if numRows == 1 {
			step = 0
		} else if j == 0 {
			step = 1
		} else if j == numRows-1 {
			step = -1
		}
		j += step
	}
	var res string
	for _, row := range rows {
		res += string(row)
	}
	return res
}

//*/
// @lc code=end