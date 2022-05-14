/*
 * @lc app=leetcode id=39 lang=golang
 *
 * [39] Combination Sum
 */

//就是递归+回溯回解
// @lc code=start
func combinationSum(candidates []int, target int) [][]int {
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
			cur = append(cur, candidates[i])
			curRes := combine(candidates, target-candidates[i], i, cur)
			res = append(res, curRes...)
			cur = cur[0 : len(cur)-1]
		}
	}
	return res
}

// @lc code=end

