/*
 * @lc app=leetcode id=179 lang=golang
 *
 * [179] Largest Number
 */

// @lc code=start
/**
 * 排序问题
 * 一个trick是排序时将a+b与b+a进行比较
 * 由于是字符串排序，他们按ascii的顺序进行比较
 * 和他们表示的数字大小顺序是一致的
 * 对于[10, 2] "102" <"210"
 * 先将数组转换成字符串后再倒序排列，最后连接
 * 对于有0前缀的直接返回0
 */
func largestNumber(nums []int) string {
	strNums := []string{}
	for i := 0; i < len(nums); i++ {
		strNums = append(strNums, strconv.Itoa(nums[i]))
	}
	sort.Slice(strNums, func(a, b int) bool {
		return strNums[a]+strNums[b] > strNums[b]+strNums[a]
	})
	if strNums[0] == "0" {
		return strNums[0]
	}
	return strings.Join(strNums, "")
}

// @lc code=end

