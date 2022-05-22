/*
 * @lc app=leetcode id=47 lang=golang
 *
 * [47] Permutations II
 */

//利用bits保存tempList已经用过的元素索引
// @lc code=start
func permuteUnique(nums []int) [][]int {
	sort.Ints(nums)
	return backtrace(nums, []int{}, 0)
}

func backtrace(nums []int, tempList []int, bits int) [][]int {
	res := [][]int{}

	if len(tempList) == len(nums) {
		tmp := make([]int, len(tempList))
		copy(tmp, tempList)
		res = append(res, tmp)
		return res
	}

	for i := 0; i < len(nums); i++ {
		if bits&(1<<i) > 0 {
			continue
		}
		//相同元素只取第一个, 在同级循环中，i在bits中只会出现一次
		if i > 0 && nums[i] == nums[i-1] && bits&(1<<(i-1)) == 0 {
			continue
		}
		bits |= 1 << i
		tempList = append(tempList, nums[i])
		res = append(res, backtrace(nums, tempList, bits)...)
		bits ^= 1 << i
		tempList = tempList[0 : len(tempList)-1]
	}
	return res
}

// @lc code=end

