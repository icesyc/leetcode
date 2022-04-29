/*
 * @lc app=leetcode id=9 lang=golang
 *
 * [9] Palindrome Number
 */

// @lc code=start
/**
 * 解法2
 */

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}
	var rev int
	y := x
	for y > 0 {
		rev = rev*10 + y%10
		y /= 10
	}
	return rev == x
}

//*/
/**
 * 解法1
 *
func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}
	var n []byte
	for x > 0 {
		n = append(n, byte(x%10))
		x = x / 10
	}
	for i, j := 0, len(n)-1; i <= j; i, j = i+1, j-1 {
		if n[i] != n[j] {
			return false
		}
	}
	return true
}
//*/
// @lc code=end

