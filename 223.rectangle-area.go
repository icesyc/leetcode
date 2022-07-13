/*
 * @lc app=leetcode id=223 lang=golang
 *
 * [223] Rectangle Area
 */

// @lc code=start
/**
 * 计算面积，先考虑不相交的情况，如果不相交直接返回a1+a2
 * 相交的情况，计算出相交时的x1,y1,x2,y2
 * 左下角取a,b中大的坐标，右上角取a,b中小的坐标即为相交的部分
 */
func computeArea(ax1 int, ay1 int, ax2 int, ay2 int, bx1 int, by1 int, bx2 int, by2 int) int {
	a1 := (ax2 - ax1) * (ay2 - ay1)
	a2 := (bx2 - bx1) * (by2 - by1)
	//没有相交
	if bx1 > ax2 || bx2 < ax1 || by1 > ay2 || by2 < ay1 {
		return a1 + a2
	}
	ix1 := max(ax1, bx1)
	iy1 := max(ay1, by1)
	ix2 := min(ax2, bx2)
	iy2 := min(ay2, by2)
	inArea := (ix2 - ix1) * (iy2 - iy1)
	return a1 + a2 - inArea
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// @lc code=end

