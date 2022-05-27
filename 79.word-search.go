/*
 * @lc app=leetcode id=79 lang=golang
 *
 * [79] Word Search
 */

// @lc code=start
func exist(board [][]byte, word string) bool {
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[0]); j++ {
			if backtrace(board, word, j, i) {
				return true
			}
		}
	}
	return false
}

func backtrace(board [][]byte, word string, x, y int) bool {
	if x < 0 || y < 0 || x == len(board[0]) || y == len(board) {
		return false
	}
	if board[y][x] != word[0] {
		return false
	}
	if len(word) == 1 {
		return true
	}
	board[y][x] = '#'
	steps := [][]int{[]int{1, 0}, []int{-1, 0}, []int{0, 1}, []int{0, -1}}
	for _, step := range steps {
		if backtrace(board, word[1:], x+step[0], y+step[1]) {
			return true
		}
	}
	board[y][x] = word[0]
	return false
}

// @lc code=end

