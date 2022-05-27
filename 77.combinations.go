/*
 * @lc app=leetcode id=77 lang=golang
 *
 * [77] Combinations
 */

// @lc code=start
/*
 * backtrace
 */
func combine(n int, k int) [][]int {
	res := &[][]int{}
	helper(n, k, 1, res, []int{})
	return *res
}

func helper(n int, k int, start int, res *[][]int, cur []int) {
	if len(cur) == k {
		tmp := make([]int, k)
		copy(tmp, cur)
		*res = append(*res, tmp)
		return
	}
	for i := start; i <= n; i++ {
		cur = append(cur, i)
		helper(n, k, i+1, res, cur)
		cur = cur[0 : len(cur)-1]
	}
}

//*/

// @lc code=end

