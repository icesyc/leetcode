/*
 * @lc app=leetcode id=78 lang=golang
 *
 * [78] Subsets
 */

// @lc code=start
func subsets(nums []int) [][]int {
	res := &[][]int{}
	backtrace(nums, 0, res, []int{})
	return *res

}
func backtrace(nums []int, start int, res *[][]int, cur []int) {
	tmp := make([]int, len(cur))
	copy(tmp, cur)
	*res = append(*res, tmp)
	for i := start; i < len(nums); i++ {
		cur = append(cur, nums[i])
		backtrace(nums, i+1, res, cur)
		cur = cur[0 : len(cur)-1]
	}
}

// @lc code=end

