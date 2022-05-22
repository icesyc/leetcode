/*
 * @lc app=leetcode id=59 lang=golang
 *
 * [59] Spiral Matrix II
 */
/**
 * 使用x,y来表示当前填充的坐标,使用dx,dy来保存当前的方向步数
 * 需要rotate 90度的条件
 * 1 到达边界
 * 2 到达之前已经填充过的坐标
 * 旋转90度的方法 dx,dy = -dy,dx
 */
// @lc code=start
func generateMatrix(n int) [][]int {
	res := make([][]int, n)
	for i := 0; i < n; i++ {
		res[i] = make([]int, n)
	}
	x, y, dx, dy := 0, 0, 1, 0
	for i := 0; i < n*n; i++ {
		res[y][x] = i + 1
		if x+dx < 0 || x+dx >= n || y+dy < 0 || y+dy >= n || res[y+dy][x+dx] != 0 {
			dx, dy = -dy, dx
		}
		x, y = x+dx, y+dy
	}
	return res
}

// @lc code=end

