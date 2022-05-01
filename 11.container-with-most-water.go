/*
 * @lc app=leetcode id=11 lang=golang
 *
 * [11] Container With Most Water
 */

// @lc code=start
func maxArea(height []int) int {
	lens := len(height)
	maxArea, i, j := 0, 0, lens-1
	for i < j {
		h := min(height[i], height[j])
		area := h * (j - i)
		if area > maxArea {
			maxArea = area
		}
		if height[i] < height[j] {
			i++
		} else {
			j--
		}
	}
	return maxArea
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// @lc code=end

