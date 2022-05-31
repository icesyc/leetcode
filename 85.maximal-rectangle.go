/*
 * @lc app=leetcode id=85 lang=golang
 *
 * [85] Maximal Rectangle
 */
/**
 * 解题思路是以84题为基础
 * 我们按行扫描， 可以得到该行每列的最大高度
 * 每个位置发现是0，就重置该列的高度为0, 否则就高度+1
 * 得到高度的数组中，就得前一题一样了，
 * 已知一个由不同高度组成的直方图，那么计算直方图的最大矩形面积
 * 计算方式见前一个题的解题说明
 * 公式为area := height[i] * (right - left - 1)
 * 关键点在于left,right的确定
 */

// @lc code=start
func maximalRectangle(matrix [][]byte) int {
	rows := len(matrix)
	cols := len(matrix[0])
	heights := make([]int, cols)
	maxArea := 0
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if matrix[i][j] == '1' {
				heights[j]++
			} else {
				heights[j] = 0
			}
		}
		//计算左右边界
		lessFromLeft, lessFromRight := calcLess(heights)
		for j := 0; j < cols; j++ {
			left, right := lessFromLeft[j], lessFromRight[j]
			area := heights[j] * (right - left - 1)
			if area > maxArea {
				maxArea = area
			}
		}
	}
	return maxArea
}

func calcLess(heights []int) ([]int, []int) {
	lessFromLeft := make([]int, len(heights))
	lessFromRight := make([]int, len(heights))
	lessFromLeft[0] = -1
	lessFromRight[len(heights)-1] = len(heights)
	for i := 1; i < len(heights); i++ {
		p := i - 1
		for p >= 0 && heights[p] >= heights[i] {
			p = lessFromLeft[p]
		}
		lessFromLeft[i] = p
	}
	for i := len(heights) - 2; i >= 0; i-- {
		p := i + 1
		for p < len(heights) && heights[p] >= heights[i] {
			p = lessFromRight[p]
		}
		lessFromRight[i] = p
	}
	return lessFromLeft, lessFromRight
}

// @lc code=end

