/**
 * 问题类似68, 给定一组单词序列words和一个最大行宽maxWidth，将单词序列分割成多行并满足以下条件：
 * 1. 每行单词+空格的长度之和为maxWidth
 * 2. 所有行使用的空格个数加起来最小
 * 3. 每行单词为两端对齐，中间使用空格填充，单词之间的空格填充尽量均匀
 * 返回分割后的所有行
 * 示例: maxWidth=10, words=[this, is, a, word, I, like, code]
 * |this____is|
 * |a_word___I|
 * |like__code|
 *
 * 使用DP的思路解题分析如下：
 * 首先假设我们只分二行，我们不知道从哪里分割是方案最优，
 * 我们需要尝试所有的分割方案，然后计算这些方案需要消耗的空格数, 再取它们中的最小值
 * [a|b,c,d] = spaceCost([a]) + spaceCost([b,c,d])
 * [a,b|c,d] = spaceCost([a,b]) + spaceCost([c,d])
 * [a,b,c|d] = spaceCost([a,b,c]) + spaceCost([d])
 * ....
 * 归纳为公式 dp[i] = min(spaceCost([i:j]) + spaceCost([j:]))
 * 如果是分割多行的话，那么spaceCost([j:])就变成了分割words[j:n-1]所需要消耗的最小空格数
 * 可以看到，这是dp[i]的一个子问题，所以我们重新归纳一下：
 * dp[i] = min(spaceCost([i:j]) + dp[j]),  i = [0:n], j = [i+1:n]
 * 由于dp[i]依赖于后续的dp计算结果，我们构建dp数组时需要先计算dp[n-1]
 * 计算顺序为dp[n-1].... dp[0]
 */
package main

import (
	"fmt"
	"math"
	"strings"
)

func main() {
	words := []string{"tushar", "roy", "likes", "to", "code"}
	width := 10
	res := justify(words, width)
	for _, line := range res {
		fmt.Println(line)
	}
}
func justify(words []string, maxWidth int) []string {
	dp, split := getDP(words, maxWidth)
	fmt.Println(dp)
	//fmt.Println(split)
	res := []string{}
	for i := 0; i < len(split); {
		end := split[i]
		newLine := justifyLine(words, maxWidth, i, end)
		res = append(res, newLine)
		i = end
	}
	return res
}
func justifyLine(words []string, width, i, j int) string {
	l := 0
	for k := i; k < j; k++ {
		l += len(words[k])
	}
	totalSpaces := width - l
	var spaceBetween, extraSpaces int
	//最后一行或一行只有一个word要左对齐
	if j-i == 1 || j == len(words) {
		spaceBetween = 0
		extraSpaces = totalSpaces
	} else {
		//j-i-1为单词间隔数
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

//计算words的dp
func getDP(words []string, maxWidth int) ([]int, []int) {
	n := len(words)
	costTable := buildSpaceCostTable(words, maxWidth)
	dp, split := make([]int, n+1), make([]int, n)
	//占位使用
	dp[n] = 0
	for i := n - 1; i >= 0; i-- {
		//先设置一个最大值占位
		dp[i] = math.MaxInt32
		for j := n - 1; j >= i; j-- {
			//这里的costTable的范围是i-j，包括j, 所以子问题的dp取值为j+1
			currentCost := costTable[i][j] + dp[j+1]
			if currentCost < dp[i] {
				dp[i] = currentCost
				//split记录[i:n-1]之间最小cost的分隔位置
				split[i] = j + 1
			}
		}
	}
	return dp, split
}

//先构建所有分割方式需要消耗的空格数量, 这里采用平方做为结果，是为了放大差异
func buildSpaceCostTable(words []string, maxWidth int) [][]int {
	costTable := make([][]int, len(words))
	for i := 0; i < len(words); i++ {
		costTable[i] = make([]int, len(words))
		lineLength := 0
		for j := i; j < len(words); j++ {
			lineLength += len(words[j])
			extraSpaces := maxWidth - lineLength - (j - i)
			if extraSpaces > 0 {
				extraSpaces *= extraSpaces
			} else {
				//extraSpace小于0表示word[i:j+1]不能放到一行
				extraSpaces = math.MaxInt32
			}
			costTable[i][j] = extraSpaces
		}
	}
	return costTable
}
