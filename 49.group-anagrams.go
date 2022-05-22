/*
 * @lc app=leetcode id=49 lang=golang
 *
 * [49] Group Anagrams
 */

//思路就是归一化，所有字母相同的映射到一组
//用freqArr记录每个字母出现的次数, 结果做为map的key
// @lc code=start
func groupAnagrams(strs []string) [][]string {
	strMap := map[string][]string{}
	for i := 0; i < len(strs); i++ {
		freqArr := make([]byte, 26)
		for j := 0; j < len(strs[i]); j++ {
			char := int(strs[i][j] - 'a')
			freqArr[char]++
		}
		key := string(freqArr)
		strMap[key] = append(strMap[key], strs[i])
	}
	res := [][]string{}
	for _, v := range strMap {
		res = append(res, v)
	}
	return res
}

// @lc code=end

