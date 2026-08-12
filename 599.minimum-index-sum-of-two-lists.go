/*
 * @lc app=leetcode id=599 lang=golang
 *
 * [599] Minimum Index Sum of Two Lists
 */

// @lc code=start
func findRestaurant(list1 []string, list2 []string) []string {
	indexMap := make(map[string]int)
	for i, w := range list1 {
		indexMap[w] = i
	}

	res := make([]string, 0)
	minSum := math.MaxInt32

	for j, w := range list2 {
		if i, ok := indexMap[w]; ok {
			sum := i + j
			if sum < minSum {
				minSum = sum
				res = []string{w}
			} else if sum == minSum {
				res = append(res, w)
			}
		}
	}

	return res
}

// @lc code=end

