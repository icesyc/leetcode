/*
 * @lc app=leetcode id=165 lang=golang
 *
 * [165] Compare Version Numbers
 */

// @lc code=start
/**
 * 双层循环，内层循环每次从version1,version2中分别取出一个数字
 * 每次遇到.后停止，计算ver1和ver2
 */
func compareVersion(version1 string, version2 string) int {
	i, j := 0, 0
	ver1, ver2 := 0, 0
	for i < len(version1) || j < len(version2) {
		for i < len(version1) && version1[i] != '.' {
			ver1 = ver1*10 + int(version1[i]-'0')
			i++
		}
		for j < len(version2) && version2[j] != '.' {
			ver2 = ver2*10 + int(version2[j]-'0')
			j++
		}
		if ver1 < ver2 {
			return -1
		} else if ver1 > ver2 {
			return 1
		}
		ver1, ver2 = 0, 0
		i++
		j++
	}
	return 0
}

// @lc code=end

