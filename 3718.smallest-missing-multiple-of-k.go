/*
 * @lc app=leetcode id=3718 lang=golang
 *
 * [3718] Smallest Missing Multiple of K
 */

// @lc code=start
func missingMultiple(nums []int, k int) int {
	m := make([]bool, len(nums) + 1)
	for _, v := range nums {
		if v % k == 0 {
			if v/k < len(m) {
				m[v/k]= true
			}
		}
	}

	for i := 1; i < len(m); i++ {
		if !m[i] {
			return i * k
		}
	}
    return len(m) * k
}
// @lc code=end

