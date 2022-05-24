/*
 * @lc app=leetcode id=73 lang=golang
 *
 * [73] Set Matrix Zeroes
 */

// @lc code=start
func setZeroes(matrix [][]int) {
	rowFlag := make([]int, len(matrix))
	colFlag := make([]int, len(matrix[0]))
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			if matrix[i][j] == 0 {
				rowFlag[i] = 1
				colFlag[j] = 1
			}
		}
	}
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			if rowFlag[i] == 1 || colFlag[j] == 1 {
				matrix[i][j] = 0
			}
		}
	}
}

// @lc code=end

