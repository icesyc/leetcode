/*
 * @lc app=leetcode id=48 lang=golang
 *
 * [48] Rotate Image
 */
/*
 * clockwise rotate
 * first reverse up to down, then swap the symmetry
 * 1 2 3     7 8 9     7 4 1
 * 4 5 6  => 4 5 6  => 8 5 2
 * 7 8 9     1 2 3     9 6 3
 */
// @lc code=start
func rotate(matrix [][]int) {
	vflip(matrix)
	n := len(matrix)
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			matrix[i][j], matrix[j][i] = matrix[j][i], matrix[i][j]
		}
	}
}
func vflip(matrix [][]int) {
	n := len(matrix)
	for i := 0; i < n/2; i++ {
		for j := 0; j < n; j++ {
			matrix[i][j], matrix[n-1-i][j] = matrix[n-1-i][j], matrix[i][j]
		}
	}
}

// [0, 0] => [0, 3], [0, 1] => [1, 3], [0, 2] => [2,3], [0,3] => [3,3]
// [1, 0] => [0, 2], [1, 1] => [1, 2], [1, 2] => [2,2], [1,3] => [3,2]
// [3,1] => [1,0], [3,2] => [2,0]
// [i,j] => [j, n-1-i]
/*
func rotate(matrix [][]int) {
	n := len(matrix)
	circle := 0
	for i := circle; i < n-1-circle; i++ {
		for j := circle; j < n-1-circle; j++ {
			//i,j => j, n-1-i
			//j,n-1-i => n-1-i, n-1-j
			//n-1-i,n-1-j => n-1-j, i
			//n-1-j,i => i,j
			tmp := matrix[i][j]
			matrix[i][j] = matrix[n-1-j][i]
			matrix[n-1-j][i] = matrix[n-1-i][n-1-j]
			matrix[n-1-i][n-1-j] = matrix[j][n-1-i]
			matrix[j][n-1-i] = tmp
		}
		circle++
	}
}
*/

// @lc code=end

