/*
 * @lc app=leetcode id=166 lang=golang
 *
 * [166] Fraction to Recurring Decimal
 */

// @lc code=start
/**
 * 关键点 分数的计算方法
 * 计算x/y的小数，令x/y 得到整数部分
 * 小数计算: 令x=x%y
 * 让x*10/y 得到的整数部分即一位小数
 * 再取余x=x*10%y，不断循环直到x==0
 * 为了发现重复，建一个map保存当前分子的值和当前的循环计数
 * 当再次发现分子的值在map中时，之前保存的计数即循环的起始位置
 * 同时要注意处理负数
 */
func fractionToDecimal(numerator int, denominator int) string {
	//处理整数
	intPart := numerator / denominator
	numerator = numerator % denominator
	res := strconv.Itoa(intPart)
	if numerator == 0 {
		return res
	}
	if intPart == 0 && numerator*denominator < 0 {
		res = "-" + res
	}
	res += "."
	//处理小数部分
	nmap := map[int]int{}
	i := 0
	decPart := ""
	for numerator != 0 {
		//是否发现重复
		if fromIdx, ok := nmap[numerator]; ok {
			decPart = decPart[:fromIdx] + "(" + decPart[fromIdx:] + ")"
			break
		}
		nmap[numerator] = i
		quot := numerator * 10 / denominator
		decPart += string(abs(quot) + '0')
		numerator = (numerator * 10) % denominator
		i++
	}
	return res + decPart
}
func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

// @lc code=end

