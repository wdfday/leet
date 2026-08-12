/*
 * @lc app=leetcode id=238 lang=golang
 *
 * [238] Product of Array Except Self
 */

// @lc code=start
func productExceptSelf(nums []int) []int {

	zeroCount := 0
	zeroIndex := -1

	prod := 1
	for i, value := range nums {
		if value == 0 {
			zeroCount++
			zeroIndex = i
			if zeroCount > 1 {
				return make([]int, len(nums), len(nums))
			}
		} else {
			prod *= value
		}
	}
	if zeroCount == 1 {
		for i := 0; i < len(nums); i++ {
			if i != zeroIndex {
				nums[i] = 0
			} else {
				nums[i] = prod
			}
		}
	} else {
		for i := 0; i < len(nums); i++ {
			nums[i] = prod / nums[i]
		}
	}

	return nums

}

// @lc code=end

