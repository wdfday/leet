/*
 * @lc app=leetcode id=977 lang=golang
 *
 * [977] Squares of a Sorted Array
 */

// @lc code=start
func sortedSquares(nums []int) []int {
	n := len(nums)
	res := make([]int, n)
	left, right := 0, n-1
	pos := n - 1 // điền từ cuối res về đầu

	for left <= right {
		l2 := nums[left] * nums[left]
		r2 := nums[right] * nums[right]
		if l2 > r2 {
			res[pos] = l2
			left++
		} else {
			res[pos] = r2
			right--
		}
		pos--
	}

	return res
}

// @lc code=end

