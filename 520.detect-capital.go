/*
 * @lc app=leetcode id=520 lang=golang
 *
 * [520] Detect Capital
 */

// @lc code=start
func detectCapitalUse(word string) bool {
	if len(word) < 2 {
		return true
	}
	var mark bool
	if word[0] < 'a' {
		mark = word[1] >= 'a'
	} else if word[1] < 'a' {
		return false
	} else {
		mark = true
	}

	for i := 2; i < len(word); i++ {
		if word[i] <= 'Z' && mark {
			return false
		} else if word[i] >= 'a' && !mark {
			return false
		}
	}

	return true

}

// @lc code=end

