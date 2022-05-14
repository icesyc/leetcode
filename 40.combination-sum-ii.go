/*
 * @lc app=leetcode id=40 lang=golang
 *
 * [40] Combination Sum II
 */
/**
 * 同上一题，关键点在于每个元素只能取一次
 * 递归时的指针改为i+1可保证每个元素取一次
 * 为了保证结果集不重复，要先排序, 循环时只取重复的第一个元素
 */
// @lc code=start
func combinationSum2(candidates []int, target int) [][]int {
	sort.Ints(candidates)
	return combine(candidates, target, 0, []int{})
}
func combine(candidates []int, target, startIdx int, cur []int) [][]int {
	res := [][]int{}
	if target == 0 {
		tmp := make([]int, len(cur))
		copy(tmp, cur)
		res = append(res, tmp)
		return res
	}
	if target > 0 {
		for i := startIdx; i < len(candidates); i++ {
			//确保重复元素只取第一个
			if i > startIdx && candidates[i] == candidates[i-1] {
				continue
			}
			cur = append(cur, candidates[i])
			curRes := combine(candidates, target-candidates[i], i+1, cur)
			res = append(res, curRes...)
			cur = cur[0 : len(cur)-1]
		}
	}
	return res
}

// @lc code=end

