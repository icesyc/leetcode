/*
 * @lc app=leetcode id=93 lang=golang
 *
 * [93] Restore IP Addresses
 */
/**
 * 典型的backtracking问题
 * ip长度为1..3, 那么我们从字符串开始不断尝试
 * 先取第一段s[:1], 再递归从字符串剩余部分取第二断
 * 使用cur保存当前的段数
 * 如果cur到达4，结束递归
 * 如果s的长度为0了，那么说明正好取完，保存cur到res中
 * 循环的部分需要判断不能出现01这样的段
 * 不能出现>255的数字
 * 同时注意i不能超过len(s), 防止越界
 */
// @lc code=start
func restoreIpAddresses(s string) []string {
	res := &[]string{}
	backtrack(s, res, []string{})
	return *res
}

func backtrack(s string, res *[]string, cur []string) {
	if len(cur) == 4 {
		if len(s) == 0 {
			*res = append(*res, strings.Join(cur, "."))
		}
		return
	}
	for i := 1; i <= len(s) && i <= 3; i++ {
		if s[0] == '0' && i > 1 {
			break
		}
		if n, _ := strconv.Atoi(s[:i]); n > 255 {
			break
		}
		cur = append(cur, s[:i])
		backtrack(s[i:], res, cur)
		cur = cur[:len(cur)-1]
	}
}

// @lc code=end

