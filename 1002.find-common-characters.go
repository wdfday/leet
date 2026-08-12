/*
 * @lc app=leetcode id=1002 lang=golang
 *
 * [1002] Find Common Characters
 */

// @lc code=start
func commonChars(words []string) []string {

	m := make(map[rune]int)

	for _, b := range words[0] {
		m[b]++
	}

	for i := 1; i < len(words); i++ {
		tmp := make(map[rune]int)
		for _, b := range words[i] {
			if v, ok := m[b]; ok {
				tmp[b]++
				if tmp[b] > v {
					tmp[b] = v
				}
			}
		}
		m = tmp
	}

	res := []string{}
	for k, v := range m {
		for range v {		
			res = append(res, string(k))
}
	}
	return res
}
// @lc code=end

