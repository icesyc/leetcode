/*
 * @lc app=leetcode id=60 lang=golang
 *
 * [60] Permutation Sequence
 */
/**
 * 假设给定数字为n=4, k = 9
 * 所有排列可分为4组，每组6个
 * 1+排列{2,3,4}
 * 2+排列{1,3,4}
 * 3+排列{1,2,4}
 * 4+排列{1,2,3}
 * 可知第9个排列在第2组, 可以确定第一个数字为2,即[1,2,3,4]中第二个元素
 * 由于索引从0开始，我们假定k=8
 * 公式为 index=k/(n-1)! n为4
 * 继续在第二组{1,3,4}里查找k的位置，因为第二组已经重新计数，
 * 前面有6个需要减去，公式为k -= index * (n-1)!
 * 再重复前面的步骤, 重复4次即可获得结果
 */
// @lc code=start
func getPermutation(n int, k int) string {
	nums := make([]byte, n)
	//保存n的阶乘
	fact := 1
	for i := 1; i <= n; i++ {
		fact *= i
		nums[i-1] = byte(i) + '0'
	}
	res := make([]byte, n)
	//以k-1计算
	k--
	for i := 0; i < n; i++ {
		//计算(n-1)! = n!/n
		fact /= (n - i)
		index := k / fact
		k -= index * fact
		res[i] = nums[index]
		nums = append(nums[0:index], nums[index+1:]...)
	}
	return string(res)
}

/*
func getPermutation(n int, k int) string {
	nums := make([]byte, n)
	for i := 0; i < n; i++ {
		nums[i] = byte(i+1) + '0'
	}
	return backtrace(nums, "", &count, k)
}

func backtrace(nums []byte, cur string, count *int, k int) string {
	if len(cur) == len(nums) {
		*count++
		if *count == k {
			return cur
		}
		return ""
	}
	for i := 0; i < len(nums); i++ {
		if strings.IndexByte(cur, nums[i]) != -1 {
			continue
		}
		cur += string(nums[i])
		str := backtrace(nums, cur, count, k)
		if str != "" {
			return str
		}
		cur = cur[0 : len(cur)-1]
	}
	return ""
}
*/

// @lc code=end

