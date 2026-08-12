/*
 * @lc app=leetcode id=239 lang=golang
 *
 * [239] Sliding Window Maximum
 */

// @lc code=start
func maxSlidingWindow(nums []int, k int) []int {
	deque := []int{}
	res := []int{}

	for r := 0; r < len(nums); r++ {
		// pop tail nếu value nhỏ hơn nums[r]
		for len(deque) > 0 && nums[deque[len(deque)-1]] < nums[r] {
			deque = deque[:len(deque)-1]
		}
		deque = append(deque, r)

		// pop front nếu out of window
		if deque[0] < r-k+1 {
			deque = deque[1:]
		}

		// ghi nhận max khi window đủ k
		if r >= k-1 {
			res = append(res, nums[deque[0]])
		}
	}
	return res
}

// @lc code=end

