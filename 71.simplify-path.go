/*
 * @lc app=leetcode id=71 lang=golang
 *
 * [71] Simplify Path
 */

// @lc code=start
func simplifyPath(path string) string {

	tokens := strings.Split(path, "/")
	res := "/"

	var stack []string

	for _, token := range tokens {
		if token == "" || token == "." {
			continue
		} else if token == ".." {
			if len(stack) > 0 {
				stack = stack[:len(stack)-1]
			}
		} else {
			stack = append(stack, token)
		}
	}

	res += strings.Join(stack, "/")
	return res

}

// @lc code=end

