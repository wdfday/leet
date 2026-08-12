/*
 * @lc app=leetcode id=953 lang=golang
 *
 * [953] Verifying an Alien Dictionary
 */

// @lc code=start
func isAlienSorted(words []string, order string) bool {
	var rank [26]int
	for i := 0; i < len(order); i++ {
		rank[order[i]-'a'] = i
	}

	inOrder := func(w1, w2 string) bool {
		for i := 0; i < len(w1) && i < len(w2); i++ {
			if w1[i] != w2[i] {
				return rank[w1[i]-'a'] < rank[w2[i]-'a']
			}
		}
		return len(w1) <= len(w2)
	}

	for i := 0; i < len(words)-1; i++ {
		if !inOrder(words[i], words[i+1]) {
			return false
		}
	}
	return true
}

// @lc code=end

