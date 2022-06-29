/*
 * @lc app=leetcode id=204 lang=golang
 *
 * [204] Count Primes
 */

// @lc code=start
/**
 * 计算小于n的所有素数(不能被1和自己以外的任何数字除开)
 * 使用数组保存每个数字是否为素数
 * 每次将数字i与j相乘(i*j<n)，将结果记为非素数
 */
func countPrimes(n int) int {
	notPrime := make([]bool, n)
	cnt := 0
	for i := 2; i < n; i++ {
		if !notPrime[i] {
			cnt++
			for j := 2; i*j < n; j++ {
				notPrime[i*j] = true
			}
		}
	}
	return cnt
}

// @lc code=end

