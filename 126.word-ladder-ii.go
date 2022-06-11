/*
 * @lc app=leetcode id=126 lang=golang
 *
 * [126] Word Ladder II
 */

// @lc code=start
/**
 * bfs搜索过程
 * 根据beginWord扩展所有可能的下一步word
 * 这里使用的技巧是对word的每个字符使用a-z进行替换，这样循环的次数会比循环wordList次数少
 * 循环所有扩展的word, 如果找到endWord,则加入结果
 * 没找到则加入到队列中，进入下一层搜索
 * queue = [[hit]]
 * path = [hit], expand=[hot], newPath = [hit,hot]
 * queue = [[hot]]
 * path = [hit,hot], expand=[dot,lot],newPath=[hit,hot,dot],[hit,hot,lot]
 * queue = [[hit,hot,dot],[hit,hot,lot]]
 * path = [hit,hot,dot], expand=[dog], newPath=[hit,hot,dot,dog]
 * path = [hit,hot,lot], expand=[log], newPath=[hit,hot,dot,log]
 * queue = [[hit,hot,dot,dog],[hit,hot,lot,log]]
 * path = [hit,hot,dot,dog], expand=[cog] newPath=[hit,hot,dot,dog,cog]
 * path = [hit,hot,lot,log], expand=[cog] newPath=[hit,hot,lot,log,cog]
 * queue = []
 */
func findLadders(beginWord string, endWord string, wordList []string) [][]string {
	res := [][]string{}
	wordMap := map[string]bool{}
	for _, word := range wordList {
		wordMap[word] = true
	}
	curPath := []string{beginWord}
	queue := [][]string{curPath}
	visited := []string{beginWord}
	for len(queue) > 0 {
		layerSize := len(queue)
		for i := 0; i < layerSize; i++ {
			curPath = queue[i]
			curWord := curPath[len(curPath)-1]
			expandList := expand(curWord, wordMap)
			for i := 0; i < len(expandList); i++ {
				visited = append(visited, expandList[i])
				//一次要注意newPath必须要复制，否则添加到queue里的newPath可能被覆盖
				newPath := append([]string{}, curPath...)
				newPath = append(newPath, expandList[i])
				if expandList[i] == endWord {
					res = append(res, newPath)
				} else {
					queue = append(queue, newPath)
				}
			}
		}
		queue = queue[layerSize:]
		//从map中移除本层已经访问过的word,下层搜索时不应该包含它们
		for i := 0; i < len(visited); i++ {
			delete(wordMap, visited[i])
		}
	}
	return res
}

//对beginWord变换一个字符，
func expand(beginWord string, wordMap map[string]bool) []string {
	expandList := []string{}
	for i := 0; i < len(beginWord); i++ {
		for j := 'a'; j <= 'z'; j++ {
			if beginWord[i] == byte(j) {
				continue
			}
			w := beginWord[:i] + string(j) + beginWord[i+1:]
			if wordMap[w] {
				expandList = append(expandList, w)
			}
		}
	}
	return expandList
}

// @lc code=end

