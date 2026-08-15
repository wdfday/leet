/*
 * @lc app=leetcode id=3541 lang=golang
 *
 * [3541] Find Most Frequent Vowel and Consonant
 */

// @lc code=start
func maxFreqSum(s string) int {
	Vowel := []byte{'a', 'e', 'i', 'o', 'u'}
	VowelF := 0
	ConsonantF := 0

	freq := make(map[byte]int)
	for i := 0; i < len(s); i++ {
		freq[s[i]]++
		if slices.Contains(Vowel, s[i]) {
			VowelF = max(VowelF, freq[s[i]])
		} else {
			ConsonantF = max(ConsonantF, freq[s[i]])
		}
	}

	return VowelF + ConsonantF
    
}
// @lc code=end

