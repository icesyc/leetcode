/*
 * @lc app=leetcode id=208 lang=golang
 *
 * [208] Implement Trie (Prefix Tree)
 */

// @lc code=start
/**
 * 就是trie树的实现，记住TrieNode的结构
 * 是使用数组来保存children, 数组的索引为字母的ascii编码
 */
type TrieNode struct {
	Val      bool
	Children [256]*TrieNode
}
type Trie struct {
	Root *TrieNode
}

func Constructor() Trie {
	return Trie{Root: &TrieNode{}}
}

func (this *Trie) Insert(word string) {
	cur := this.Root
	for i := 0; i < len(word); i++ {
		c := int(word[i])
		if cur.Children[c] == nil {
			cur.Children[c] = &TrieNode{}
		}
		cur = cur.Children[c]
	}
	cur.Val = true
}

func (this *Trie) SearchNode(word string) *TrieNode {
	cur := this.Root
	for i := 0; i < len(word); i++ {
		c := int(word[i])
		if cur.Children[c] == nil {
			return nil
		}
		cur = cur.Children[c]
	}
	return cur
}
func (this *Trie) Search(word string) bool {
	node := this.SearchNode(word)
	return node != nil && node.Val
}

func (this *Trie) StartsWith(prefix string) bool {
	node := this.SearchNode(prefix)
	return node != nil
}

/**
 * Your Trie object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Insert(word);
 * param_2 := obj.Search(word);
 * param_3 := obj.StartsWith(prefix);
 */
// @lc code=end

