/*
 * @lc app=leetcode id=46 lang=golang
 *
 * [46] Permutations
 */

// @lc code=start
func permute(nums []int) [][]int {
	return backtrace(nums, []int{})
}

func backtrace(nums []int, perm []int) [][]int {
	res := [][]int{}
	if len(perm) == len(nums) {
		tmp := make([]int, len(perm))
		copy(tmp, perm)
		res = append(res, tmp)
		return res
	}
	for i := 0; i < len(nums); i++ {
		if indexOf(perm, nums[i]) >= 0 {
			continue
		}
		perm = append(perm, nums[i])
		nextRes := backtrace(nums, perm)
		res = append(res, nextRes...)
		perm = perm[0 : len(perm)-1]
	}
	return res
}

func indexOf(arr []int, value int) int {
	for i := 0; i < len(arr); i++ {
		if arr[i] == value {
			return i
		}
	}
	return -1
}

// @lc code=end

