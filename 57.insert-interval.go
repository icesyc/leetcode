/*
 * @lc app=leetcode id=57 lang=golang
 *
 * [57] Insert Interval
 */

// @lc code=start
func insert(intervals [][]int, newInterval []int) [][]int {
	res := [][]int{}
	i := 0
	start, end := newInterval[0], newInterval[1]
	for ; i < len(intervals) && intervals[i][1] < start; i++ {
		res = append(res, intervals[i])
	}
	for ; i < len(intervals) && intervals[i][0] <= end; i++ {
		start = min(start, intervals[i][0])
		end = max(end, intervals[i][1])
	}
	res = append(res, []int{start, end})
	if i < len(intervals) {
		res = append(res, intervals[i:]...)
	}
	return res
}

/*
func insert(intervals [][]int, newInterval []int) [][]int {
	res := [][]int{}
	start, end := newInterval[0], newInterval[1]
	inserted := false
	for i := 0; i < len(intervals); i++ {
		if start <= intervals[i][1] {
			start = min(start, intervals[i][0])
		}
		if end >= intervals[i][0] {
			end = max(end, intervals[i][1])
		}
		if intervals[i][1] < start {
			res = append(res, intervals[i])
		} else if intervals[i][0] > end {
			inserted = true
			res = append(res, []int{start, end})
			res = append(res, intervals[i:]...)
			break
		}
	}
	if !inserted {
		res = append(res, []int{start, end})
	}
	return res
}
*/

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// @lc code=end

