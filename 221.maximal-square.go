/*
 * @lc app=leetcode id=221 lang=golang
 *
 * [221] Maximal Square
 */

// @lc code=start
/**
 * 最小的方块面积，如果矩阵中没有连续的1，那么边长只有1
 * 如果4个方块中都有1，边长为2
 * 定义dp[i][j]为matrix位置[i,j]的最大矩形边长
 * 可以看到dp[i][j]取决于3个位置，上，左和左上
 * 以示例绿色部分为例
 * 如果3个位置都是1，那么该位置的边长为2
 * 继续扩展到9个方块，如果如果左、上、左上都为2，边长为3
 * 也就是dp[i][j]取上、左、左上中的最小值+1
 * 得出公式为
 * if matrix[i][j] == '0' dp[i][j] = 0
 * if matrix[i][j] == '1' dp[i][j] = min(dp[i-1][j-1], dp[i-1][j], dp[i][j-1]) + 1
 * 对于第一行和第一列的位置，dp[i][j]只能是matrix[i][j]
 */
func maximalSquare(matrix [][]byte) int {
	m, n := len(matrix), len(matrix[0])
	dp := make([][]int, m)
	max := 0
	for i := 0; i < m; i++ {
		dp[i] = make([]int, n)
		for j := 0; j < n; j++ {
			if i == 0 || j == 0 {
				dp[i][j] = int(matrix[i][j] - '0')
			} else if matrix[i][j] == '1' {
				dp[i][j] = min(dp[i-1][j-1], min(dp[i-1][j], dp[i][j-1])) + 1
			}
			if dp[i][j] > max {
				max = dp[i][j]
			}
		}
	}
	return max * max
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// @lc code=end

