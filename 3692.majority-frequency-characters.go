/*
 * @lc app=leetcode id=3692 lang=golang
 *
 * [3692] Majority Frequency Characters
 */

// @lc code=start
func majorityFrequencyGroup(s string) string {
	freq := make(map[byte]int)

	for i := 0; i < len(s); i++ {
		freq[s[i]]++
	}

	groups := make(map[int][]byte)

	for ch, f := range freq {
		groups[f] = append(groups[f], ch)
	}

	maxSize := 0
	res := ""
	maxFreq := 0

	for f, chars := range groups {
		if len(chars) > maxSize {
			maxSize = len(chars)
			res = string(chars)
			maxFreq = f
		} else if len(chars) == maxSize {
			if f > maxFreq {
				res = string(chars)
				maxFreq = f
			}
		}
	}

	return res

}
// @lc code=end

