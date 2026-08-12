/*
 * @lc app=leetcode id=860 lang=golang
 *
 * [860] Lemonade Change
 */

// @lc code=start
func lemonadeChange(bills []int) bool {

	five, ten := 0, 0
	for _, b := range bills {
		switch b {
		case 5:
			five++
		case 10:
			ten++
			five--
			if five < 0 {
				return false
			}
		case 20:
			if ten > 0 && five > 0 {
				ten--
				five--
			} else {
				five -= 3
			}

			if five < 0 {
				return false
			}
		}
	}
	return true
}

// @lc code=end

