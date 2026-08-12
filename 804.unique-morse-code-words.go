/*
 * @lc app=leetcode id=804 lang=golang
 *
 * [804] Unique Morse Code Words
 */

// @lc code=start
func uniqueMorseRepresentations(words []string) int {

	codes := []string {".-","-...","-.-.","-..",".","..-.","--.","....","..",".---","-.-",".-..","--","-.","---",".--.","--.-",".-.","...","-","..-","...-",".--","-..-","-.--","--.."}
	m := make(map[string]struct{})

	for _, word := range words {
		v := ""
		for _, c := range word {
			v += codes[int(c - 'a')]
		}
		m[v] = struct{}{}
	}
	return len(m)
    
}
// @lc code=end

