/*
 * @lc app=leetcode id=152 lang=golang
 *
 * [152] Maximum Product Subarray
 */

// @lc code=start
func maxProduct(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	res := nums[0]
	currMax := nums[0]
	currMin := nums[0]

	max2 := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}
	min2 := func(a, b int) int {
		if a < b {
			return a
		}
		return b
	}

	for i := 1; i < len(nums); i++ {
		x := nums[i]
		if x < 0 {
			currMax, currMin = currMin, currMax
		}
		currMax = max2(x, currMax*x)
		currMin = min2(x, currMin*x)
		if currMax > res {
			res = currMax
		}
	}

	return res
}

// @lc code=end

