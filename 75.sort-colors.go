/*
 * @lc app=leetcode id=75 lang=golang
 *
 * [75] Sort Colors
 */
/**
 * three-way partition问题
 * 给定一个数组，将数组按targetValue分为三部分
 * [0:i] < target, [i:k] = target, [k:n] > target
 * 设置三个指针i, j, k，循环数组
 * nums[j] < target => 放到[0:i]区
 * nums[j] = target => 放到[i:k]区
 * nums[j] > target => 放到[k:n]区
 */
// @lc code=start
func sortColors(nums []int) {
	i, j, k := 0, 0, len(nums)-1
	for j <= k {
		if nums[j] < 1 {
			nums[j], nums[i] = nums[i], nums[j]
			i++
			j++
		} else if nums[j] > 1 {
			nums[j], nums[k] = nums[k], nums[j]
			//置换后还要继续比较，所以不需要向前移动j
			k--
		} else {
			j++
		}
	}
}

// @lc code=end

