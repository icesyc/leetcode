/*
 * @lc app=leetcode id=52 lang=golang
 *
 * [52] N-Queens II
 */

/**
 * 解法同上一题，还要简单一些，这次换一种方法
 * flagCol保存哪些列被占用 有n条
 * flag45保存45度上被占用的线 有2n-1条
 * flag135保存15度上被占用的线 有2n-1条
 * 45度的线公式col+row
 * 135度的线公式n-1+col-row
 */
// @lc code=start
type BitMap struct {
	flagCol []bool
	flag45  []bool
	flag135 []bool
}

func (bitMap *BitMap) Add(row, col int) bool {
	n := len(bitMap.flagCol)
	if bitMap.flagCol[col] || bitMap.flag45[row+col] || bitMap.flag135[n-1+col-row] {
		return false
	}
	bitMap.flagCol[col] = true
	bitMap.flag45[row+col] = true
	bitMap.flag135[n-1+col-row] = true
	return true
}
func (bitMap *BitMap) Remove(row, col int) {
	n := len(bitMap.flagCol)
	bitMap.flagCol[col] = false
	bitMap.flag45[row+col] = false
	bitMap.flag135[n-1+col-row] = false
}
func totalNQueens(n int) int {
	bitMap := &BitMap{
		flagCol: make([]bool, n),
		flag45:  make([]bool, 2*n-1),
		flag135: make([]bool, 2*n-1),
	}
	return dfs(bitMap, 0)
}

func dfs(bitMap *BitMap, row int) int {
	n := len(bitMap.flagCol)
	res := 0
	if row == n {
		res++
		return res
	}
	for col := 0; col < n; col++ {
		if bitMap.Add(row, col) {
			res += dfs(bitMap, row+1)
			bitMap.Remove(row, col)
		}
	}
	return res
}

// @lc code=end

