/*
 * @lc app=leetcode id=149 lang=golang
 *
 * [149] Max Points on a Line
 */
// @lc code=start
/**
 * 根据斜率公式 a=(y2-y1)/(x2-x1)
 * 斜率相同的点在一条线上, 我们使用dx/dy的基数作为key
 * 需要计算dx,dy的最大公约数
 * 每次遍历初始化一个map, 把对应斜率的点+1
 * 最后别忘了把对比的点i也要加进去, 返回的max要+1
 */
func maxPoints(points [][]int) int {
	max := 0
	for i := 0; i < len(points); i++ {
		pointMap := map[string]int{}
		for j := i + 1; j < len(points); j++ {
			dx := points[j][0] - points[i][0]
			dy := points[j][1] - points[i][1]
			g := gcd(dx, dy)
			dx /= g
			dy /= g
			pointMap[key]++
			if pointMap[key] > max {
				max = pointMap[key]
			}
		}
	}
	return max + 1
}

func gcd(a, b int) int {
	if b == 0 {
		return a
	}
	return gcd(b, a%b)
}

// @lc code=end

