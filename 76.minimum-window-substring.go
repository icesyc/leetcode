/*
 * @lc app=leetcode id=76 lang=golang
 *
 * [76] Minimum Window Substring
 */
/**
 * 字符串搜索问题，双指针算法
 * 使用map保存t中所有的字符及个数
 * start,end来标识一个滑动窗口
 * 遍历s, 将窗口向右移动，如果出现t中的字符串，则计数
 * 当计数符合条件时，说明该窗口有效，
 * 此时移动左边界，直到发现s[start]为t中的字符
 * start-end即为符合条件的窗口，记录minStart, minLen
 * 继续该过程，发现有效窗口，记录minStart,minLen 直到结束
 */
// @lc code=start
func minWindow(s string, t string) string {
	start, end, minStart, minLen := 0, 0, 0, math.MaxInt
	counter := len(t)
	tmap := map[byte]int{}
	for i := 0; i < len(t); i++ {
		tmap[t[i]]++
	}

	for ; end < len(s); end++ {
		if tmap[s[end]] > 0 {
			counter--
		}
		tmap[s[end]]--
		//找到有效窗口, 移动start指针，直到发现t中的字符
		for counter == 0 {
			//计算窗口大小
			if end-start+1 < minLen {
				minStart = start
				minLen = end - start + 1
			}
			//恢复所有字符的计数
			tmap[s[start]]++
			//由当前条件可知 tmap中所有t的字符计数都为0,
			//其它为负数, 如果+1后>0，就说明该字符在t中
			//有一种情况，如窗口为[A B A C], 由于A出现两次
			//tmap[A] = -1，tmap[B]=0，start指向B时
			//才满足条件, 取的才是最小窗口
			if tmap[s[start]] > 0 {
				//窗口排除了一个t中的字符, 需要修改counter计数
				counter++
			}
			start++
		}
	}

	if minLen != math.MaxInt {
		return string(s[minStart : minStart+minLen])
	}
	return ""
}

// @lc code=end

