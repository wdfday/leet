/*
 * @lc app=leetcode id=30 lang=golang
 *
 * [30] Substring with Concatenation of All Words
 */

// @lc code=start
func findSubstring(s string, words []string) []int {
	if s == "" || len(words) == 0 {
		return []int{}
	}

	wordCount := make(map[string]int)
	for _, word := range words {
		wordCount[word]++
	}

	wordLength := len(words[0])
	substringLength := wordLength * len(words)
	result := []int{}

	for i := 0; i <= len(s)-substringLength; i++ {
		seenWords := make(map[string]int)
		j := 0
		for j < len(words) {
			word := s[i+j*wordLength : i+(j+1)*wordLength]
			if count, exists := wordCount[word]; !exists || seenWords[word] >= count {
				break
			}
			seenWords[word]++
			j++
		}
		if j == len(words) {
			result = append(result, i)
		}
	}

	return result
}

// @lc code=end

