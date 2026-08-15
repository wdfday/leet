/*
 * @lc app=leetcode id=3550 lang=golang
 *
 * [3550] Smallest Index With Digit Sum Equal to Index
 */

// @lc code=start
func smallestIndex(nums []int) int {
	for i, v := range nums {
		t := 0
		for v > 0 {
			t += v % 10
			v /= 10
		}
		if t == i {
			return i
		}
	}
	return -1
}
// @lc code=end

