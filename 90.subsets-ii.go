/*
 * @lc app=leetcode id=90 lang=golang
 *
 * [90] Subsets II
 */
/**
 * backtacking 关键点，因为有重复元素，要先排序
 * 再使用bits保证backtracking同一层级相同元素只取第一个
 */
// @lc code=start
func subsetsWithDup(nums []int) [][]int {
	sort.Ints(nums)
	res := &[][]int{}
	backtrack(nums, 0, res, []int{}, 0)
	return *res
}

func backtrack(nums []int, start int, res *[][]int, cur []int, bits int) {
	*res = append(*res, append([]int{}, cur...))
	for i := start; i < len(nums); i++ {
		//利用backtracking的特性，同一个层级 i对应的只会出现一次
		//如果i-1的bits为0说明是同一层级
		if i > 0 && nums[i] == nums[i-1] && bits&(1<<(i-1)) == 0 {
			continue
		}
		cur = append(cur, nums[i])
		bits |= 1 << i
		backtrack(nums, i+1, res, cur, bits)
		cur = cur[:len(cur)-1]
		bits ^= 1 << i
	}
}

// @lc code=end

