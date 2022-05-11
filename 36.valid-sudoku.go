/*
 * @lc app=leetcode id=36 lang=golang
 *
 * [36] Valid Sudoku
 */
/**
 * 使用三个长度为9的数组保存x,y,z(表示子区域)坐标
 * 如mx[0]为第一行, my[2]为第三列, mz[1]为第二个区域
 * 对于每个元素，用1-9bit表示数字1-9是否出现过
 * 遍历表，并将表中的数字对应的bit位置1
 * 如果已经设置过，表示数独表无效
 */
// @lc code=start
func isValidSudoku(board [][]byte) bool {
	mx := [9]int{}
	my := [9]int{}
	mz := [9]int{}
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			n := board[i][j]
			if n == '.' {
				continue
			}
			bit := 1 << (int(n) - 49)
			x, y, z := i, j, i/3*3+j/3
			if mx[x]&bit == bit || my[y]&bit == bit || mz[z]&bit == bit {
				return false
			}
			mx[x] |= bit
			my[y] |= bit
			mz[z] |= bit
		}
	}
	return true
}

// @lc code=end

