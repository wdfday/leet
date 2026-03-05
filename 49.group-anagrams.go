/*
 * @lc app=leetcode id=49 lang=golang
 *
 * [49] Group Anagrams
 */

// @lc code=start
func groupAnagrams(strs []string) [][]string {
	anagramMap := make(map[[26]uint16][]string, len(strs))
	for _, s := range strs {
		var key [26]uint16
		for i := 0; i < len(s); i++ {
			key[s[i]-'a']++
		}
		anagramMap[key] = append(anagramMap[key], s)
	}
	res := make([][]string, 0, len(anagramMap))
	for _, v := range anagramMap {
		res = append(res, v)
	}
	return res
}

// @lc code=end

// func groupAnagrams(strs []string) [][]string {

// 	m := make(map[string][]string)
// 	for _, s := range strs {
// 		key := sortString(s)
// 		m[key] = append(m[key], s)
// 	}
// 	var res [][]string
// 	for _, v := range m {
// 		res = append(res, v)
// 	}
// 	return res
// }

// func sortString(s string) string {
// 	runes := []rune(s)
// 	sort.Slice(runes, func(i, j int) bool {
// 		return runes[i] < runes[j]
// 	})
// 	return string(runes)
// }
