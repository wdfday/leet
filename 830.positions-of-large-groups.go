/*
 * @lc app=leetcode id=830 lang=golang
 *
 * [830] Positions of Large Groups
 */

// @lc code=start
func largeGroupPositions(s string) [][]int {

	res := make([][]int, 0)

	index := 0
	tex := s[0]
	for i := 1; i < len(s); i++ {
		if s[i] != tex {
			if i-index >= 3 {
				res = append(res, []int{index, i - 1})
			}
			index = i
			tex = s[i]

		}
	}

	if len(s)-index >= 3 {
		res = append(res, []int{index, len(s) - 1})
	}

	return res
}

// @lc code=end

// func largeGroupPositions(s string) [][]int {
// 	res := make([][]int, 0)
// 	n := len(s)

// 	i := 0
// 	for i < n {
// 		j := i
// 		for j < n && s[j] == s[i] {
// 			j++
// 		}
// 		if j-i >= 3 {
// 			res = append(res, []int{i, j - 1})
// 		}
// 		i = j
// 	}

// 	return res
// }