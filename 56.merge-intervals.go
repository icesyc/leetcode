/*
 * @lc app=leetcode id=56 lang=golang
 *
 * [56] Merge Intervals
 */

// @lc code=start
func merge(intervals [][]int) [][]int {
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})
	res := [][]int{intervals[0]}
	j := 0
	for i := 1; i < len(intervals); i++ {
		//排好序的 res[j][0]是当前区间的最小值,扩展区间右侧
		if intervals[i][0] <= res[j][1] {
			if intervals[i][1] > res[j][1] {
				res[j][1] = intervals[i][1]
			}
			continue
		}
		res = append(res, intervals[i])
		j++
	}
	return res
}

// @lc code=end

