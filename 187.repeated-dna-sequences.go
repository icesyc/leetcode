/*
 * @lc app=leetcode id=187 lang=golang
 *
 * [187] Repeated DNA Sequences
 */

// @lc code=start
func findRepeatedDnaSequences(s string) []string {
	seqMap := map[string]int{}
	res := []string{}
	for i := 0; i <= len(s)-10; i++ {
		seq := s[i : i+10]
		seqMap[seq]++
	}
	for seq, n := range seqMap {
		if n > 1 {
			res = append(res, seq)
		}
	}
	return res
}

// @lc code=end

