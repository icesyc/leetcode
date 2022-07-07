/*
 * @lc app=leetcode id=216 lang=golang
 *
 * [216] Combination Sum III
 */

// @lc code=start
/**
 * 回溯法
 * 由于只能取1-9，所以回溯的循环为start..9
 * 结束条件为k==0 || n == 0
 * 如果k和n都为0, 那么找到一个解，添加到结果中
 * 否则循环,添加i到当前数组中,继续递归求解n-i, k-1
 */
func combinationSum3(k int, n int) [][]int {
	return backtrack(k, n, 1, []int{})
}

func backtrack(k int, n int, start int, cur []int) [][]int {
	res := [][]int{}
	if k == 0 && n == 0 {
		tmp := make([]int, len(cur))
		copy(tmp, cur)
		res = append(res, tmp)
		return res
	} else if k == 0 || n == 0 {
		return res
	}
	for i := start; i <= 9 && i <= n; i++ {
		cur = append(cur, i)
		nextRes := backtrack(k-1, n-i, i+1, cur)
		res = append(res, nextRes...)
		cur = cur[:len(cur)-1]
	}
	return res
}

// @lc code=end

