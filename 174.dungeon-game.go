/*
 * @lc app=leetcode id=174 lang=golang
 *
 * [174] Dungeon Game
 */

// @lc code=start
/**
 * 考虑dp[i][j]为到达i,j点所需要的最小血量
 * dp[m-1][n-1]为公主所在的房间，所需要的最小血量为 1-dungeon[m-1][n-1]
 * 由于只能往右或往下走，所以我们在i,j房间所需要的最小血量取决于下一个房间
 * 再减去当前房间的消耗
 * 即dp[i][j] = min(dp[i+1][j], dp[i][j+1]) - dungeon[i][j]
 * 如果最小血量小于0，说明补血充足，但我们也需要至少1点，骑士的血量不能小于1
 * 最后算出dp[0][0]即为初始状态所需的最小血量
 * 在计算时，先计算出最右和最下方的房间
 *
 * 这里需要注意的是，我们不能从左上开始计算， 如果从左上开始计算
 * 每次取左或上受到的最小伤害
 * 例如[-2,-5,10,30,-5]，因为中途补的血量超过了受到的伤害值，所以计算是错误
 * 我们至少需要8点血量才能到达10的位置，计算会更加复杂
 */
func calculateMinimumHP(dungeon [][]int) int {
	m, n := len(dungeon), len(dungeon[0])
	dp := make([][]int, m)
	for i := 0; i < m; i++ {
		dp[i] = make([]int, n)
	}
	dp[m-1][n-1] = max(1, 1-dungeon[m-1][n-1])
	for j := n - 2; j >= 0; j-- {
		dp[m-1][j] = max(1, dp[m-1][j+1]-dungeon[m-1][j])
	}
	for i := m - 2; i >= 0; i-- {
		dp[i][n-1] = max(1, dp[i+1][n-1]-dungeon[i][n-1])
	}
	for i := m - 2; i >= 0; i-- {
		for j := n - 2; j >= 0; j-- {
			dp[i][j] = max(1, min(dp[i+1][j], dp[i][j+1])-dungeon[i][j])
		}
	}
	return dp[0][0]
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

