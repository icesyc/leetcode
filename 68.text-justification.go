/*
 * @lc app=leetcode id=68 lang=golang
 *
 * [68] Text Justification
 */

// @lc code=start
func fullJustify(words []string, maxWidth int) []string {
	i, j, lineLength := 0, 0, 0
	res := []string{}
	for i < len(words) {
		lineLength = len(words[i])
		j = i + 1
		//lineLength保存words[i:j]的所有word长度之和
		//words[i:j+1]之间的空格长度(j-i)
		for j < len(words) && lineLength+len(words[j])+(j-i) <= maxWidth {
			lineLength += len(words[j])
			j++
		}
		totalSpaces := maxWidth - lineLength
		res = append(res, justify(words, totalSpaces, i, j))
		i = j
	}
	return res
}

//对words[i:j]进行排版
func justify(words []string, totalSpaces, i, j int) string {
	//j-i为words[i:j]之间隔断数
	var spaceBetween, extraSpaces int
	//最后一行或一行只有一个word要左对齐
	if j-i == 1 || j == len(words) {
		spaceBetween = 0
		extraSpaces = totalSpaces
	} else {
		spaceBetween = totalSpaces / (j - i - 1)
		extraSpaces = totalSpaces % (j - i - 1)
	}
	var b strings.Builder
	b.WriteString(words[i])
	for k := i + 1; k < j; k++ {
		spaces := spaceBetween
		if extraSpaces > 0 {
			spaces += 1
			extraSpaces--
		}
		b.WriteString(strings.Repeat(" ", spaces))
		b.WriteString(words[k])
	}
	//对于左对齐的，填充剩余的空格
	if extraSpaces > 0 {
		b.WriteString(strings.Repeat(" ", extraSpaces))
	}
	return b.String()
}

// @lc code=end

