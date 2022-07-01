/*
 * @lc app=leetcode id=211 lang=golang
 *
 * [211] Design Add and Search Words Data Structure
 */

// @lc code=start
/**
 * 使用trie树+backtracking解决
 * 查询word时，如果发现是"."，则遍历所有Children检查word[1:]
 * 不是".", 看节点存在并且是一个单词的结束
 * backtracking结束条件
 * node为nil, 或者len(word)==0
 */
type TrieNode struct {
	Val      bool
	Children [26]*TrieNode
}
type WordDictionary struct {
	Root *TrieNode
}

func Constructor() WordDictionary {
	return WordDictionary{
		Root: &TrieNode{},
	}
}

func (this *WordDictionary) AddWord(word string) {
	cur := this.Root
	for i := 0; i < len(word); i++ {
		c := int(word[i] - 'a')
		if cur.Children[c] == nil {
			cur.Children[c] = &TrieNode{}
		}
		cur = cur.Children[c]
	}
	cur.Val = true
}

func (this *WordDictionary) Search(word string) bool {
	return this.backtrack(word, this.Root)
}

func (this *WordDictionary) backtrack(word string, node *TrieNode) bool {
	if node == nil {
		return false
	}
	if len(word) == 0 {
		return node.Val
	}
	c := int(word[0])
	if c != '.' {
		c = int(word[0] - 'a')
		return this.backtrack(word[1:], node.Children[c])
	}
	for i := 0; i < len(node.Children); i++ {
		if this.backtrack(word[1:], node.Children[i]) {
			return true
		}
	}
	return false
}

/**
 * Your WordDictionary object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AddWord(word);
 * param_2 := obj.Search(word);
 */
// @lc code=end

