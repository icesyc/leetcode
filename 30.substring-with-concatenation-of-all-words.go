/*
 * @lc app=leetcode id=30 lang=golang
 *
 * [30] Substring with Concatenation of All Words
 */

// @lc code=start
/**
 * slide window
 * https://leetcode.wang/leetCode-30-Substring-with-Concatenation-of-All-Words.html
 */
//*/
func findSubstring(s string, words []string) []int {
	res := []int{}
	count := map[string]int{}
	wordLen, wordCount, strLen := len(words[0]), len(words), len(s)
	for _, word := range words {
		count[word]++
	}

	for i := 0; i < wordLen; i++ {
		seenNum, left := 0, i
		seen := map[string]int{}
		for j := i; j <= strLen-wordLen; j += wordLen {
			word := s[j : j+wordLen]
			//不匹配，重置
			if count[word] == 0 {
				seen = map[string]int{}
				left = j + wordLen
				seenNum = 0
				continue
			}
			seen[word]++
			if seen[word] <= count[word] {
				seenNum++
			} else {
				//word多了一个,从左侧删除直到该word数量=count中的数量
				for seen[word] > count[word] {
					leftWord := s[left : left+wordLen]
					seen[leftWord]--
					//除了当前word外，其它删除需要同步减少已匹配的单词数
					if seen[leftWord] < count[leftWord] {
						seenNum--
					}
					left += wordLen
				}
			}
			//找到完整匹配
			if seenNum == wordCount {
				res = append(res, left)
				seen[s[left:left+wordLen]]--
				seenNum--
				left += wordLen
			}
		}
	}
	return res
}

/**
 *
func findSubstring(s string, words []string) []int {
	res := []int{}
	count := map[string]int{}
	wordLen, wordCount := len(words[0]), len(words)
	for _, word := range words {
		count[word]++
	}

	for i := 0; i < len(s)-wordLen*wordCount+1; i++ {
		seen := map[string]int{}
		seenNum := 0
		for seenNum < wordCount {
			start := i + seenNum*wordLen
			word := s[start : start+wordLen]
			if count[word] == 0 || count[word] == seen[word] {
				break
			}
			seen[word]++
			seenNum++
		}
		if seenNum == wordCount {
			res = append(res, i)
		}
	}
	return res

}
//*/
// @lc code=end

