/*
 * @lc app=leetcode id=45 lang=golang
 *
 * [45] Jump Game II
 */

// @lc code=start
/**
 * O(n)算法的解法
 * 从0开始，保存当前能跳的最远距离，计为end
 * 循环1-end之间查找第2跳的最远距离
 * 循环结束时(i==end)，设置最远距离为end
 * 继续循环 直到结束
 * 3 2 4 2 1 5 4
 * -------
 * 如上图第一跳end的索引为3, 遍历后4跳的最远
 * 3 2 4 2 1 5 4
 * -------
 *     |--------
 */
func jump(nums []int) int {
	jump, end, farthest := 0, 0, 0
	for i := 0; i < len(nums)-1; i++ {
		if nums[i]+i > farthest {
			farthest = nums[i] + i
		}
		if i == end {
			jump++
			end = farthest
		}
	}
	return jump
}

/*
func jump(nums []int) int {
	jumpIndex := -1
	for i := len(nums) - 1; i >= 0; i-- {
		if nums[i]+i >= len(nums)-1 {
			jumpIndex = i
		}
	}
	if len(nums) == 1 {
		return 0
	}
	if jumpIndex == 0 {
		return 1
	}
	return 1 + jump(nums[0:jumpIndex+1])
}
*/
// @lc code=end

