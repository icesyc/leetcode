/*
 * @lc app=leetcode id=42 lang=golang
 *
 * [42] Trapping Rain Water
 */
/**
 * 设置左右两个指针往中间扫描, 同时保存左右两侧的最高位置
 * 每次扫描移动高度小的一边，发现高度小于最高位置，就累计当前水的容量
 * 假设输入为[4,2,1,3]
 * left=4, right=3，此时我们需要移动right
 * right-1的值为1，那么无论左侧是什么，水的容量至少是可以有2个单元
 * 如果我们移动left，left-1的值为2，我们无法保证左侧最大值4-2为当前水的容量
 * 见图(https://raw.githubusercontent.com/icesyc/diagram/main/trapping_rain_water.drawio.svg)
 *
 */
// @lc code=start
func trap(height []int) int {
	left, right := 0, len(height)-1
	max := 0
	maxLeft, maxRight := 0, 0
	for left < right {
		if height[left] < height[right] {
			if height[left] > maxLeft {
				maxLeft = height[left]
			} else {
				max += maxLeft - height[left]
			}
			left++
		} else {
			if height[right] > maxRight {
				maxRight = height[right]
			} else {
				max += maxRight - height[right]
			}
			right--
		}
	}
	return max
}

// @lc code=end

