/*
 * @lc app=leetcode id=3740 lang=golang
 *
 * [3740] Minimum Distance Between Three Equal Elements I
 */

// @lc code=start
func minimumDistance(nums []int) int {
	indices := make(map[int][]int)
	for i, v := range nums {
		indices[v] = append(indices[v], i)
	}

	res := math.MaxInt
	for _, idxs := range indices {
		if len(idxs) < 3 {
			continue
		}
		for i := 2; i < len(idxs); i++ {
			dist := 2 * (idxs[i] - idxs[i-2])
			if dist < res {
				res = dist
			}
		}
	}

	if res == math.MaxInt {
		return -1
	}
	return res
}

// @lc code=end

