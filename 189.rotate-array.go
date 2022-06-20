/*
 * @lc app=leetcode id=189 lang=golang
 *
 * [189] Rotate Array
 */

// @lc code=start
/**
 * 建立临时数组，将nums尾部的k个数字保存
 * 从后往前将nums[k..n]值向右移动k个位置
 *
 */
/*
func rotate(nums []int, k int) {
	k = k % len(nums)
	if k == 0 {
		return
	}
	tmpArr := make([]int, k)
	for i := 0; i < k; i++ {
		tmpArr[i] = nums[len(nums)-k+i]
	}
	for i := len(nums) - 1; i >= k; i-- {
		nums[i] = nums[i-k]
	}
	for i := 0; i < k; i++ {
		nums[i] = tmpArr[i]
	}
}
*/
/**
 * 不需要临时数组
 * 需要reserse 3次
 * nums = "----->-->"; k =3
 * result = "-->----->";
 * reverse "----->-->" we can get "<--<-----"
 * reverse "<--" we can get "--><-----"
 * reverse "<-----" we can get "-->----->"
 */
func rotate(nums []int, k int) {
	k = k % len(nums)
	reverse(nums, 0, len(nums)-1)
	reverse(nums, 0, k-1)
	reverse(nums, k, len(nums)-1)
}

func reverse(nums []int, start, end int) {
	for start < end {
		nums[start], nums[end] = nums[end], nums[start]
		start++
		end--
	}
}

// @lc code=end

