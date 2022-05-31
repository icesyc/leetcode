/*
 * @lc app=leetcode id=84 lang=golang
 *
 * [84] Largest Rectangle in Histogram
 */
/**
 * 考虑任意位置的bar，计算该bar对应的矩形面积为
 * area := height[i] * (right - left - 1)
 * 其中left为指针从bar往左侧移动时遇到的第一个高度小于height[i]的位置
 * right为指针从bar往右侧移动时遇到的第一个高度小于height[i]的位置
 * 使用升序stack保存当前的位置i, 循环heights, 当heights[i] < height[stack.top]时， 表明我们找到right
 * 我们需要计算stack中每个位置的最大面积
 * 由于stack是升序，所以stack中所有位置的right都为i, stack中每个p的left都为p-1(有相等元素除外，但不影响计算结果)
 * 我们需要不断pop stack,直到height[stack.top] < height[i]或者stack为空, 对于stack中的每个位置，分别计算对应的面积
 * 高度(H)为height[stack.top], right为i, left为stack.top-1
 * maxArea = max(H * (right-left-1))
 * 例如：heights = [2,1,5,6,2,3]
 * 为了方便计算最后一个元素的right, 给heights增加一个0, [2,1,5,6,2,3,0]
 * heights[0]的lessLeft = -1, heights[5]的lessRight = 6
 * i=0, heights[i]=2, stack=[-1] => stack =>[-1,0]
 * i=1, heights[i]=1, stack=[-1,0], H=2, right=1,left=-1, area=2 => stack=[-1,1]
 * i=2, heights[i]=5, stack=[-1,1] => stack=[-1,1,2]
 * i=3, heights[i]=6, stack=[-1,1,2] => stack=[-1,1,2,3]
 * i=4, heights[i]=2, stack=[-1,1,2,3], H=6,right=4,left=2, area=6 => stack=[-1,1,2]
 * .....................................H=5,right=4,left=1,area=10 => stack=[-1,1]
 * .....................................计算结束 stack=[-1,1,4]
 * i=5, heights[i]=3, stack=[-1,1,4] => stack=[-1,1,4,5]
 * i=6, heights[i]=0, stack=[-1,1,4,5], H=3, right=6, left=4, area=3 => stack => [-1,1,4]
 * .....................................H=2, right=6, left=1,area=8 => stack=>[-1,1]
 * .....................................H=1, right=6, left=-1, area=6 => stack=>[-1]
 *
 */
// @lc code=start
func largestRectangleArea(heights []int) int {
	stack := []int{-1}
	heights = append(heights, 0)
	ans := 0
	for i := 0; i < len(heights); i++ {
		top := stack[len(stack)-1]
		for len(stack) > 1 && heights[top] >= heights[i] {
			h := heights[top]
			right := i
			left := stack[len(stack)-2]
			area := h * (right - left - 1)
			if area > ans {
				ans = area
			}
			top = left
			stack = stack[0 : len(stack)-1]
		}
		stack = append(stack, i)
	}
	return ans
}

// @lc code=end

/**
 * 考虑任意位置的bar，计算该bar对应的矩形面积为
 * area := height[i] * (right - left - 1)
 * 其中left为指针从bar往左侧移动时遇到的第一个高度小于height[i]的位置
 * right为指针从bar往右侧移动时遇到的第一个高度小于height[i]的位置
 * 我们需要计算每个bar对应的lessFromLeft和lessFromRight
 * 让p初始为i-1, 并不断向左移动，直到找到小于height[i]的位置
 * 关键是我们可以利用之前计算的结果，不需要p--
 * 让p = lessFromLeft[p]，为什么这样可以？
 * 首先如果p=i-1时 height[p] < height[i]的话，那么
 * height[i] = p 就是我们要计算的结果
 * 如果height[p]>=height[i], 我们需要往左侧移动
 * 已知我们已经计算过lessFromLeft[i-1]
 * lessFromLeft[i-1]保存的是i-1位置的lessLeft
 * 假设为q，那么可知height[q+1:i-2]都是要>=height[i-1]的
 * 也就是说height[q+1:i-1]都是>=height[i]的，
 * 所以我们将p = lessFromLeft[p]跳到位置q上，继续检查
 *
 * lessFromRight计算同上
 */
/*
func largestRectangleArea(heights []int) int {
	max := 0
	lessFromLeft, lessFromRight := calcLess(heights)
	for i := 0; i < len(heights); i++ {
		left, right := lessFromLeft[i], lessFromRight[i]
		area := heights[i] * (right - left - 1)
		if area > max {
			max = area
		}
	}
	return max
}

func calcLess(heights []int) ([]int, []int) {
	lessFromLeft := make([]int, len(heights))
	lessFromRight := make([]int, len(heights))
	lessFromLeft[0] = -1
	lessFromRight[len(heights)-1] = len(heights)
	for i := 1; i < len(heights); i++ {
		p := i - 1
		for p >= 0 && heights[p] >= heights[i] {
			p = lessFromLeft[p]
		}
		lessFromLeft[i] = p
	}
	for i := len(heights) - 2; i >= 0; i-- {
		p := i + 1
		for p < len(heights) && heights[p] >= heights[i] {
			p = lessFromRight[p]
		}
		lessFromRight[i] = p
	}
	return lessFromLeft, lessFromRight
}
*/
