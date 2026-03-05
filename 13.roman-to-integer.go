/*
 * @lc app=leetcode id=13 lang=golang
 *
 * [13] Roman to Integer
 */

// @lc code=start
func romanToInt(s string) int {
	romanMap := map[byte]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}

	total := 0
	for i := 0; i < len(s); i++ {
		value := romanMap[s[i]]
		if i+1 < len(s) && value < romanMap[s[i+1]] {
			total -= value
		} else {
			total += value
		}
	}

	return total

}

// @lc code=end

