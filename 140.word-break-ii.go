/*
 * @lc app=leetcode id=140 lang=golang
 *
 * [140] Word Break II
 */

// @lc code=start
/**
 * 回溯法
 * 循环s[0..n] 如果发现一个词s[0:i]，加入到组合中
 * 再递归剩余的s[i:]部分，递归的结束条件为剩余的子串长度为0
 * 此时将之前的临时保存的组合加入到最终结果中
 *
 * 可以优化的地方是使用map缓存每次遍历的结果，后续如果有相同的case可直接返回结果
 * 如[catsandlog] 第1个结果[cat,sand,log]，当生成第二个结果[cats, and, log]时
 * log的结果已经被缓存，可以直接使用，减少后面的循环
 */
func wordBreak(s string, wordDict []string) []string {
	wordMap := map[string]bool{}
	for _, word := range wordDict {
		wordMap[word] = true
	}
	res := &[]string{}
	backtrack(s, wordMap, []string{}, res)
	return *res
}

func backtrack(s string, wordMap map[string]bool, cur []string, res *[]string) {
	if len(s) == 0 {
		*res = append(*res, strings.Join(cur, " "))
		return
	}
	for i := 1; i <= len(s); i++ {
		word := s[:i]
		if wordMap[word] {
			cur = append(cur, word)
			backtrack(s[i:], wordMap, cur, res)
			cur = cur[:len(cur)-1]
		}
	}
}

// @lc code=end

