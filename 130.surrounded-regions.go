/*
 * @lc app=leetcode id=130 lang=golang
 *
 * [130] Surrounded Regions
 */

// @lc code=start
/**
 * 根据题目可发现，如果O所在的cell的四个方向能连接到边界，那么O不能翻转
 * 其它的O才能翻转
 * 我们遍历4个边界，使用dfs搜索，将所有连接到边界的O先翻转成#
 * 最后遍历整个board,将O变成X, 将#置回O即可
 * dfs的终止条件是到达边界或cell=X|# (#是因为已经访问过了)
 */
func solve(board [][]byte) {
	m, n := len(board), len(board[0])
	//检查第一列和最后一列
	for i := 0; i < m; i++ {
		if board[i][0] == 'O' {
			dfs(board, i, 0)
		}
		if board[i][n-1] == 'O' {
			dfs(board, i, n-1)
		}
	}
	//检查第一行和最后一行
	for j := 0; j < n; j++ {
		if board[0][j] == 'O' {
			dfs(board, 0, j)
		}
		if board[m-1][j] == 'O' {
			dfs(board, m-1, j)
		}
	}
	//翻转所有cell
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if board[i][j] == '#' {
				board[i][j] = 'O'
			} else if board[i][j] == 'O' {
				board[i][j] = 'X'
			}
		}
	}
}

func dfs(board [][]byte, i, j int) {
	m, n := len(board), len(board[0])
	//到达边界
	if i < 0 || i >= m || j < 0 || j >= n || board[i][j] != 'O' {
		return
	}
	board[i][j] = '#'
	dfs(board, i+1, j)
	dfs(board, i, j+1)
	dfs(board, i-1, j)
	dfs(board, i, j-1)
}

// @lc code=end

