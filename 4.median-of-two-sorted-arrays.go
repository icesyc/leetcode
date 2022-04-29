/*
 * @lc app=leetcode id=4 lang=golang
 *
 * [4] Median of Two Sorted Arrays
 */

// @lc code=start
//ref https://www.youtube.com/watch?v=LPFhl65R7ww&t=1014s
func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	m := len(nums1)
	n := len(nums2)
	if m > n {
		return findMedianSortedArrays(nums2, nums1)
	}
	low, high := 0, m
	for low <= high {
		posX := (low + high) / 2
		posY := (m+n+1)/2 - posX
		var maxXLeft, maxYLeft, minXRight, minYRight int
		if posX == 0 {
			maxXLeft = math.MinInt
		} else {
			maxXLeft = nums1[posX-1]
		}
		if posX == m {
			minXRight = math.MaxInt
		} else {
			minXRight = nums1[posX]
		}
		if posY == 0 {
			maxYLeft = math.MinInt
		} else {
			maxYLeft = nums2[posY-1]
		}
		if posY == n {
			minYRight = math.MaxInt
		} else {
			minYRight = nums2[posY]
		}
		if maxXLeft <= minYRight && maxYLeft <= minXRight {
			if (m+n)%2 == 0 {
				left := max(maxXLeft, maxYLeft)
				right := min(minXRight, minYRight)
				return float64(left+right) / 2
			}
			return float64(max(maxXLeft, maxYLeft))
		} else if maxXLeft > minYRight {
			high = posX - 1
		} else {
			low = posX + 1
		}
	}
	return 0
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

