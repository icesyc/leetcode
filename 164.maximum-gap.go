/*
 * @lc app=leetcode id=164 lang=golang
 *
 * [164] Maximum Gap
 */

// @lc code=start
/**
 * 看题目是需要排序，但要求O(n)的时间和空间复杂度
 * 考虑桶排序, 如何划分桶是关键
 * 先找出nums的最小值和最大值min,max
 * 我们知道长度为n的数组，有n-1个gap
 * 这些gap的平均值为ceil((max-min) /(n-1))
 * 我们知道最大的gap一定是>=这个平均值的
 * 那么我们可以将平均值作为桶的size，将所有元素放到对应的桶中
 * [min..min+size-1]...[min+k*size..min+(k+1)size-1]
 * 因为所有小于这个平均值的都在1个桶内，所以maxGap一定是在两个桶之间
 * maxGap=bucket[i].min -bucket[i-1].max
 * 我们只要遍历所有的桶，找出桶之间最大的gap即可
 * 还要考虑如果有空桶，需要跳过这些桶
 *
 */
func maximumGap(nums []int) int {
	minVal, maxVal := nums[0], nums[0]
	n := len(nums)
	for _, i := range nums {
		minVal = min(minVal, i)
		maxVal = max(maxVal, i)
	}
	if minVal == maxVal {
		return 0
	}
	//保存桶的最大值和最小值
	bucketMin, bucketMax := make([]int, n), make([]int, n)
	for i := 0; i < n; i++ {
		bucketMax[i] = math.MinInt
		bucketMin[i] = math.MaxInt
	}
	//计算出gap的平均值,并以平均值的大小作为桶的尺寸
	bucketSize := math.Ceil(float64(maxVal-minVal) / float64(n-1))
	for _, i := range nums {
		//桶的索引
		idx := (i - minVal) / int(bucketSize)
		bucketMax[idx] = max(bucketMax[idx], i)
		bucketMin[idx] = min(bucketMin[idx], i)
	}
	//循环所有桶，找出最大gap
	maxGap := 0
	prev := bucketMax[0]
	for i := 1; i < n; i++ {
		//skip emtpy bucket
		if bucketMin[i] == math.MaxInt {
			continue
		}
		maxGap = max(maxGap, bucketMin[i]-prev)
		prev = bucketMax[i]
	}
	return maxGap
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// @lc code=end

