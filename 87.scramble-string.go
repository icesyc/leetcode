/*
 * @lc app=leetcode id=87 lang=golang
 *
 * [87] Scramble String
 */
/**
 * 递归+memorize
 * 基本思路是将s1在每个可能的位置切割
 * 假设切割的位置为i, 切割后判断左右串是否为scramble，
 * 再对调s1的左右两个串，与s2再进行判断, 判断逻辑如下：
 * 1. isScramble(s1[:i],s2[:i]) is true ?
 *    isScramble(1[i:],s2[i:]) is true ?
 * 2. isScramble(s1[:i],s2[l-i:]) is true ?
 *    isScramble(s1[i:], s2[:l-i]) is true ?
 * 如果两种情况有一种满足, 说明s1和s2互为扰乱的字符串
 * 递归的base case:  长度为1 只需要判断s1==s2
 * 同时使用s1+s2做为key生成cache,
 * 每次计算完成保存到cache中，防止重复计算
 */
// @lc code=start
func isScramble(s1 string, s2 string) bool {
	cache := map[string]bool{}
	return helper(s1, s2, cache)
}

func helper(s1, s2 string, cache map[string]bool) bool {
	l := len(s1)
	if l == 1 {
		res := s1 == s2
		cache[s1+s2] = res
		return res
	}
	if res, ok := cache[s1+s2]; ok {
		return res
	}
	res := false
	for i := 1; i < l; i++ {
		res = res || helper(s1[:i], s2[:i], cache) && helper(s1[i:], s2[i:], cache)
		res = res || helper(s1[:i], s2[l-i:], cache) && helper(s1[i:], s2[:l-i], cache)
	}
	cache[s1+s2] = res
	return res

}

// @lc code=end

