/*
 * @lc app=leetcode id=74 lang=golang
 *
 * [74] Search a 2D Matrix
 */

// @lc code=start
func searchMatrix(matrix [][]int, target int) bool {
	rows, cols := len(matrix), len(matrix[0])
	left := 0
	right := rows*cols - 1

	for left <= right {
		mid := (left + right) / 2
		row := mid / cols
		col := mid % cols
		if matrix[row][col] == target {
			return true
		} else if matrix[row][col] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return false
}

// @lc code=end

