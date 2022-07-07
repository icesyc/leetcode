/*
 * @lc app=leetcode id=220 lang=golang
 *
 * [220] Contains Duplicate III
 */

// @lc code=start
/**
 * 类似于桶排序的原理
 * 题目的要求有2个，1是差值<=t, 2是间距在k以内
 * 我们将桶的大小设置为t+1, 桶内最大值-最小值的结果<=t
 * 为什么是t+1呢，而不是t? 假设t为3，那么桶宽度为4,
 * 0,1,2,3都在一个桶内,如果宽度为3,3就会落到下一个桶内
 * 这样对于满足nums[i]-nums[j] <=t 的元素只有3种情况
 * 都在一个桶k内，在k-1和k内，在k和k+1内
 * 分在2个桶的时候，需要判断nums[i]-nums[j] <= t
 *
 * 对于第2个条件，我们使用滑动窗口，使map中仅保存最近k个桶
 * 超过的桶全都删除
 *
 * 另外由于nums[i]可能是负数，为了保证所有桶都是正数
 * 需要做一个remap, 将所有数字都-math.MinInt32
 */
func containsNearbyAlmostDuplicate(nums []int, k int, t int) bool {
	if k == 0 {
		return false
	}
	buckets := map[int]int{}
	for i, v := range nums {
		remapNum := v - math.MinInt32
		bucket := remapNum / (t + 1)
		if _, ok := buckets[bucket]; ok {
			return true
		}
		if n, ok := buckets[bucket-1]; ok && remapNum-n <= t {
			return true
		}
		if n, ok := buckets[bucket+1]; ok && n-remapNum <= t {
			return true
		}
		if len(buckets) >= k {
			lastBucket := (nums[i-k] - math.MinInt32) / (t + 1)
			delete(buckets, lastBucket)
		}
		buckets[bucket] = remapNum
	}
	return false
}

// @lc code=end

