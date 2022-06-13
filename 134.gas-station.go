/*
 * @lc app=leetcode id=134 lang=golang
 *
 * [134] Gas Station
 */

// @lc code=start
/**
 * 如果gas总量大于cost的总量，必然没有解
 * tank保存经过每一站剩余的气量
 * 如果tank <0,说明无法从0到达i站, 那起始点只能是i+1..n
 * 由于题目说明只有一个解，那循环结束后start必然是解
 */
func canCompleteCircuit(gas []int, cost []int) int {
	m := len(gas)
	start, tank, totalGas := 0, 0, 0
	for i := 0; i < m; i++ {
		tank += gas[i] - cost[i]
		totalGas += gas[i] - cost[i]
		if tank < 0 {
			start = i + 1
			tank = 0
		}
	}
	if totalGas < 0 {
		return -1
	}
	return start
}

// @lc code=end

