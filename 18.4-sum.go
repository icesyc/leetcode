/*
 * @lc app=leetcode id=18 lang=golang
 *
 * [18] 4Sum
 */

// @lc code=start
/**
 * 解法2 kSum
 */
func fourSum(nums []int, target int) [][]int {
	sort.Ints(nums)
	return kSum(nums, target, 4)
}

func kSum(nums []int, target int, k int) [][]int {
	var res [][]int
	if k == 2 {
		start, end := 0, len(nums)-1
		for start < end {
			if start > 0 && nums[start] == nums[start-1] {
				start++
				continue
			}
			if end < len(nums)-1 && nums[end] == nums[end+1] {
				end--
				continue
			}
			sum := nums[start] + nums[end]
			if sum == target {
				res = append(res, []int{nums[start], nums[end]})
				start++
				end--
			} else if sum < target {
				start++
			} else {
				end--
			}
		}
	} else {
		for i := 0; i < len(nums)-k+1; i++ {
			//target小于最小元素的k倍或大于最大元素的k倍，表明不存在这样的组合
			if target < nums[i]*k || target > nums[len(nums)-1]*k {
				break
			}
			if i > 0 && nums[i] == nums[i-1] {
				continue
			}
			temp := kSum(nums[i+1:], target-nums[i], k-1)
			for _, group := range temp {
				res = append(res, append([]int{nums[i]}, group...))
			}
		}
	}
	return res
}

/**
 * 解法1
 *
func fourSum(nums []int, target int) [][]int {
	sort.Ints(nums)
	var res [][]int
	for i := 0; i < len(nums)-3; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		for j := i + 1; j < len(nums)-2; j++ {
			if j > i+1 && nums[j] == nums[j-1] {
				continue
			}
			start, end := j+1, len(nums)-1
			for start < end {
				if start > j+1 && nums[start] == nums[start-1] {
					start++
					continue
				}
				if end < len(nums)-1 && nums[end] == nums[end+1] {
					end--
					continue
				}
				sum := nums[i] + nums[j] + nums[start] + nums[end]
				if sum == target {
					res = append(res, []int{nums[i], nums[j], nums[start], nums[end]})
					start++
				} else if sum < target {
					start++
				} else {
					end--
				}
			}
		}
	}
	return res
}
//*/
// @lc code=end

