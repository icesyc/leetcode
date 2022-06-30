/*
 * @lc app=leetcode id=205 lang=golang
 *
 * [205] Isomorphic Strings
 */

// @lc code=start
/**
 * 字符串s是否可以映射到t, 要求每个字符只能映射一次
 * 使用两个map使用s->t,t->s的映射，循环判断即可
 */
func isIsomorphic(s string, t string) bool {
	sMap := map[byte]byte{}
	tMap := map[byte]byte{}
	for i := 0; i < len(s); i++ {
		if c, ok := sMap[s[i]]; ok && c != t[i] {
			return false
		}
		if c, ok := tMap[t[i]]; ok && c != s[i] {
			return false
		}
		sMap[s[i]] = t[i]
		tMap[t[i]] = s[i]
	}
	return true
}

// @lc code=end

