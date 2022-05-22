/*
 * @lc app=leetcode id=54 lang=golang
 *
 * [54] Spiral Matrix
 */

// @lc code=start
func spiralOrder(matrix [][]int) []int {
	left, right, top, bottom := 0, len(matrix[0])-1, 0, len(matrix)-1
	res := make([]int, 0)
	//direction 1, 2, 3, 4 => right, bottom, left, top
	dir := 1
	for left <= right && top <= bottom {
		if dir == 1 {
			for i := left; i <= right; i++ {
				res = append(res, matrix[top][i])
			}
			top++
			dir = 2
		} else if dir == 2 {
			for i := top; i <= bottom; i++ {
				res = append(res, matrix[i][right])
			}
			right--
			dir = 3
		} else if dir == 3 {
			for i := right; i >= left; i-- {
				res = append(res, matrix[bottom][i])
			}
			bottom--
			dir = 4
		} else if dir == 4 {
			for i := bottom; i >= top; i-- {
				res = append(res, matrix[i][left])
			}
			left++
			dir = 1
		}
	}
	return res
}

// @lc code=end

