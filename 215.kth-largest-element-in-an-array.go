/*
 * @lc app=leetcode id=215 lang=golang
 *
 * [215] Kth Largest Element in an Array
 */

// @lc code=start
/**
 * 求第k大，排序问题
 */
/*
func findKthLargest(nums []int, k int) int {
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] > nums[j]
	})
	return nums[k-1]
}
*/

/**
 * 使用快排的partition算法
 * 找到pivot,使pivot左侧的元素都大于pivot,右侧都小于pivot
 * 如果pivot的位置正好是k,那么就是第k大
 * 否则我们根据poviot的位置来决定是继续排左边还是排右边
 */
func findKthLargest(nums []int, k int) int {
	left, right := 0, len(nums)-1
	k = len(nums) - k
	for left < right {
		pos := partition(nums, left, right)
		if pos < k {
			left = pos + 1
		} else if pos > k-1 {
			right = pos - 1
		} else {
			break
		}
	}
	return nums[k]
}

func partition(nums []int, left, right int) int {
	pivot := nums[left]
	for left < right {
		for left < right && nums[right] >= pivot {
			right--
		}
		nums[left] = nums[right]
		for left < right && nums[left] <= pivot {
			left++
		}
		nums[right] = nums[left]
	}
	nums[left] = pivot
	return left
}

// @lc code=end

