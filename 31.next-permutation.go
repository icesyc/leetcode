/*
 * @lc app=leetcode id=31 lang=golang
 *
 * [31] Next Permutation
 */

// @lc code=start
/**
 * 下一个排列是在所有排列中 排序仅大于当前值的一组值
 * 1 3 2 4 => 1 3 4 2, 1 4 2 3 => 1 4 3 2
 * 也就是说要尽可能的从末尾调整，来保证调整后的值刚好大于当前值
 * 这里有一个特征，假设要调整的位置为i, 那么一定存在
 * nums[i] < nums[i+1] 并且nums[i+1:n] 为降序排列
 * 基于以上规则，可以从右侧开始扫描，找到nums[i]< nums[i+1]的值
 * 然后在nums[i+1:n]里去找到刚比好nums[i]大的值，并交换
 * 为了保证调整后的值最小，我们需要将nums[i+1:n]做升序排列
 * 由于交换后nums[i+1:n]仍然保持降序，我们只需要反转即可
 * 将nums[i+1:n]的头尾两两交换
 */
func nextPermutation(nums []int) {
	l := len(nums)
	idx := -1
	for i := l - 2; i >= 0; i-- {
		if nums[i] < nums[i+1] {
			idx = i
			break
		}
	}
	low, high := idx+1, l-1
	//nums[idx+1:n] is descending order, use binary search
	//for example 5 [9 6 5 4 3]
	if idx >= 0 {
		for low <= high {
			mid := (low + high) / 2
			if nums[mid] > nums[idx] {
				low = mid + 1
			} else {
				high = mid - 1
			}
		}
		//num[high] is the most right value great than nums[idx]
		nums[idx], nums[high] = nums[high], nums[idx]
	}
	//reverse the right part
	for i, j := idx+1, l-1; i < j && nums[i] > nums[j]; i, j = i+1, j-1 {
		nums[i], nums[j] = nums[j], nums[i]
	}
}

// @lc code=end

