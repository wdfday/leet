/*
 * @lc app=leetcode id=888 lang=golang
 *
 * [888] Fair Candy Swap
 */

// @lc code=start
func fairCandySwap(aliceSizes []int, bobSizes []int) []int {
	sumA, sumB := 0, 0
	setA := make(map[int]bool)

	for _, v := range aliceSizes {
		sumA += v
		setA[v] = true
	}

	for _, v := range bobSizes {
		sumB += v
	}

	diff := (sumA - sumB) / 2

	for _, y := range bobSizes {
		x := y + diff
		if setA[x] {
			return []int{x, y}
		}
	}

	return nil
}

// @lc code=end

