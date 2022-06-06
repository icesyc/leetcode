/*
 * @lc app=leetcode id=118 lang=golang
 *
 * [118] Pascal's Triangle
 */
/**
 * 设置好初始行，后面的每一行根据前一行值两两相加(除第一个和最后一个值)
 */

// @lc code=start
func generate(numRows int) [][]int {
	res := [][]int{[]int{1}}
	for i := 1; i < numRows; i++ {
		lastLevel := res[i-1]
		curLevel := []int{lastLevel[0]}
		for j := 1; j < i; j++ {
			curLevel = append(curLevel, lastLevel[j]+lastLevel[j-1])
		}
		curLevel = append(curLevel, lastLevel[0])
		res = append(res, curLevel)
	}
	return res
}

// @lc code=end

