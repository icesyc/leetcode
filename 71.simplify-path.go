/*
 * @lc app=leetcode id=71 lang=golang
 *
 * [71] Simplify Path
 */

// @lc code=start
func simplifyPath(path string) string {
	//path以/开头，第一个为空串
	pathSeg := []string{}
	for i := 1; i < len(path); {
		j := nextSep(path, i)
		f := string(path[i:j])
		if f == ".." {
			if len(pathSeg) > 0 {
				pathSeg = pathSeg[:len(pathSeg)-1]
			}
		} else if f != "." && f != "" {
			pathSeg = append(pathSeg, f)
		}
		i = j + 1
	}
	return "/" + strings.Join(pathSeg, "/")
}

func nextSep(path string, startIndex int) int {
	for i := startIndex; i < len(path); i++ {
		if path[i] == '/' {
			return i
		}
	}
	return len(path)
}

// @lc code=end

