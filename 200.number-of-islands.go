/*
 * @lc app=leetcode id=200 lang=golang
 *
 * [200] Number of Islands
 */

// @lc code=start
/**
 * .dfs搜索，将所有1标识为#，每次遇到新的1,即为新的岛屿
 */
func numIslands(grid [][]byte) int {
	res := 0
	m, n := len(grid), len(grid[0])
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == '1' {
				res++
				dfs(grid, i, j)
			}
		}
	}
	return res
}

func dfs(grid [][]byte, i, j int) {
	m, n := len(grid), len(grid[0])
	if i < 0 || i >= m || j < 0 || j >= n {
		return
	}
	if grid[i][j] == '0' || grid[i][j] == '#' {
		return
	}
	grid[i][j] = '#'
	dfs(grid, i+1, j)
	dfs(grid, i-1, j)
	dfs(grid, i, j+1)
	dfs(grid, i, j-1)
}

// @lc code=end

