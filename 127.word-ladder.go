/*
 * @lc app=leetcode id=127 lang=golang
 *
 * [127] Word Ladder
 */

// @lc code=start
/**
 * 与上题相同，返回最短路径的长度即可, 不需要path
 * 一个优化的方案是建立两个map，一个从beginword扩展，一个从endword扩展
 * 每层循环长度小的map，直到两个map元素包含相同元素
 */
func ladderLength(beginWord string, endWord string, wordList []string) int {
	wordMap := map[string]bool{}
	for _, w := range wordList {
		wordMap[w] = true
	}
	if !wordMap[endWord] {
		return 0
	}
	beginMap, endMap := map[string]bool{}, map[string]bool{}
	beginMap[beginWord] = true
	endMap[endWord] = true
	//如果一次就能转换成功，长度是2，所以从2开始
	changes := 2
	for len(beginMap) > 0 {
		tempMap := map[string]bool{}
		for word := range beginMap {
			//扩展
			for i := 0; i < len(word); i++ {
				for j := 'a'; j <= 'z'; j++ {
					w := word[:i] + string(j) + word[i+1:]
					//在endMap中找到，返回路径长度
					if endMap[w] {
						return changes
					}
					//未找到，加入到下一屋循环的列表中,并从词表中删除
					if wordMap[w] {
						tempMap[w] = true
						delete(wordMap, w)
					}
				}
			}
		}
		beginMap = tempMap
		if len(beginMap) > len(endMap) {
			beginMap, endMap = endMap, beginMap
		}
		changes++
	}
	return 0
}

// @lc code=end

