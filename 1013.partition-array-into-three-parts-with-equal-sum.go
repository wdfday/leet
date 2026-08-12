/*
 * @lc app=leetcode id=1013 lang=golang
 *
 * [1013] Partition Array Into Three Parts With Equal Sum
 */

// @lc code=start
func canThreePartsEqualSum(arr []int) bool {
	total := 0
	for _, v := range arr {
		total += v
	}
	if total%3 != 0 {
		return false
	}
	target := total / 3

	count, cur := 0, 0
	for i, v := range arr {
		cur += v
		if cur == target {
			count++
			cur = 0
			if count == 2 && i != len(arr)-1 {
				return true
			}
		}
	}
	return false
}

// @lc code=end

