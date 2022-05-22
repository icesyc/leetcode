/*
 * @lc app=leetcode id=51 lang=golang
 *
 * [51] N-Queens
 */

// @lc code=start
func solveNQueens(n int) [][]string {
	bitMap := make([]byte, n*n)
	for i := 0; i < len(bitMap); i++ {
		bitMap[i] = '.'
	}
	return solve(bitMap, 0, n, 0)
}
func solve(bitMap []byte, placedNum, n, start int) [][]string {
	res := [][]string{}
	if placedNum == n {
		res = append(res, copyMap(bitMap, n))
		return res
	}
	for i := start; i < start+n; i++ {
		if canPlace(bitMap, i, n) {
			bitMap[i] = 'Q'
			curRes := solve(bitMap, placedNum+1, n, start+n)
			res = append(res, curRes...)
			bitMap[i] = '.'
		}
	}
	return res
}

func copyMap(bitMap []byte, n int) []string {
	resStr := []string{}
	for i := 0; i < n; i++ {
		resStr = append(resStr, string(bitMap[i*n:i*n+n]))
	}
	return resStr
}

//填充map，返回是否填充成功
func canPlace(bitMap []byte, position int, n int) bool {
	x, y := position/n, position%n
	//检查每列
	for i := 0; i < n; i++ {
		if bitMap[i*n+y] == 'Q' {
			return false
		}
	}
	//检查45度
	for i, j := x, y; i >= 0 && j < n; i, j = i-1, j+1 {
		if bitMap[i*n+j] == 'Q' {
			return false
		}
	}
	//检查135度
	for i, j := x, y; i >= 0 && j >= 0; i, j = i-1, j-1 {
		if bitMap[i*n+j] == 'Q' {
			return false
		}
	}
	return true
}

// @lc code=end

