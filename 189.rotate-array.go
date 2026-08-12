/*
 * @lc app=leetcode id=189 lang=golang
 *
 * [189] Rotate Array
 */

// @lc code=start
func rotate(nums []int, k int) {
	n := len(nums)
	if n == 0 {
		return
	}

	k %= n
	if k == 0 {
		return
	}

	reverse := func(left, right int) {
		for left < right {
			nums[left], nums[right] = nums[right], nums[left]
			left++
			right--
		}
	}

	reverse(0, n-1)
	reverse(0, k-1)
	reverse(k, n-1)
}

// @lc code=end

