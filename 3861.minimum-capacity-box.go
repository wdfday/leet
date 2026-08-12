/*
 * @lc app=leetcode id=3861 lang=golang
 *
 * [3861] Minimum Capacity Box
 */

// @lc code=start
func minimumIndex(capacity []int, itemSize int) int {
	res := -1
	minsize := math.MaxInt32
	for i := 0; i < len(capacity); i++ {
		if capacity[i] >= itemSize && capacity[i] < minsize {
			res = i
			minsize = capacity[i]
		}
	}
	return res
}

// @lc code=end

