/*
 * @lc app=leetcode id=119 lang=golang
 *
 * [119] Pascal's Triangle II
 */
/**
 * 只需要一维数组即可
 * 每次使用该数组生成下一行的数据
 */
// @lc code=start
func getRow(rowIndex int) []int {
	row := make([]int, rowIndex+1)
	for i := 0; i <= rowIndex; i++ {
		prev := row[0]
		for j := 1; j < i; j++ {
			tmp := row[j]
			row[j] += prev
			prev = tmp
		}
		row[i] = 1
	}
	return row
}

// @lc code=end

