/*
 * @lc app=leetcode id=925 lang=golang
 *
 * [925] Long Pressed Name
 */

// @lc code=start
func isLongPressedName(name string, typed string) bool {
	i, j := 0, 0

	for i < len(name) && j < len(typed) {
		if name[i] != typed[j] {
			return false
		}

		cn := 1
		for i+1 < len(name) && name[i+1] == name[i] {
			i++
			cn++
		}

		ct := 1
		for j+1 < len(typed) && typed[j+1] == typed[j] {
			j++
			ct++
		}

		if ct < cn {
			return false
		}

		i++
		j++
	}

	return i == len(name) && j == len(typed)
}

// @lc code=end

