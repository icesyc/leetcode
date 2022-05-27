/*
 * @lc app=leetcode id=80 lang=golang
 *
 * [80] Remove Duplicates from Sorted Array II
 */

// @lc code=start
/**
 * 如果保证每个元素最多重复一次的话
 * [1,1,2,3,3,4,6]
 * i-2的元素一定和i不相同
 * 使用i指针保存最左侧的位置，发现不满足条件就移动元素到指针处
 * 同时向前移动指针
 */
func removeDuplicates(nums []int) int {
	i := 0
	for j := 0; j < len(nums); j++ {
		if i < 2 || nums[j] > nums[i-2] {
			nums[i] = nums[j]
			i++
		}
	}
	return i
}

/**
 * 设置两个指针
 * i 指向当前要替换的位置
 * j 指向要比较的元素
 * cnt 保存当前元素的重复个数
 * 发现元素相同，进行计数+1, 不同重置计数为1
 * 如果计数<=2，向前移动位置，将元素移到i的位置
 */
/*
func removeDuplicates(nums []int) int {
	i, cnt := 0, 1
	for j := 1; j < len(nums); j++ {
		if nums[i] == nums[j] {
			cnt++
		} else {
			cnt = 1
		}
		if cnt <= 2 {
			i++
			nums[i] = nums[j]
		}
	}
	return i + 1
}
*/
// @lc code=end

