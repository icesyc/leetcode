/*
 * @lc app=leetcode id=212 lang=golang
 *
 * [212] Word Search II
 */

// @lc code=start
/**
 * 先使用trie树构建所有word,
 * 再使用dfs搜索board, 检查trie树的当前节点是否为word
 */
type TrieNode struct {
	word     string
	Children [26]*TrieNode
}

func findWords(board [][]byte, words []string) []string {
	res := &[]string{}
	trie := buildTrie(words)
	m, n := len(board), len(board[0])
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			backtrack(board, i, j, trie, res)
		}
	}
	return *res
}

func buildTrie(words []string) *TrieNode {
	root := &TrieNode{}
	for _, word := range words {
		cur := root
		for i := 0; i < len(word); i++ {
			c := word[i] - 'a'
			if cur.Children[c] == nil {
				cur.Children[c] = &TrieNode{}
			}
			cur = cur.Children[c]
		}
		cur.word = word
	}
	return root
}

func backtrack(board [][]byte, i, j int, node *TrieNode, res *[]string) {
	m, n := len(board), len(board[0])
	if i < 0 || j < 0 || i >= m || j >= n {
		return
	}
	c := board[i][j]
	idx := c - 'a'
	if c == '#' || node.Children[idx] == nil {
		return
	}
	node = node.Children[c-'a']
	if node.word != "" {
		*res = append(*res, node.word)
		//防止重复添加
		node.word = ""
	}
	board[i][j] = '#'
	backtrack(board, i+1, j, node, res)
	backtrack(board, i-1, j, node, res)
	backtrack(board, i, j+1, node, res)
	backtrack(board, i, j-1, node, res)
	board[i][j] = c
}

// @lc code=end

