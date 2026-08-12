/*
 * @lc app=leetcode id=747 lang=golang
 *
 * [747] Largest Number At Least Twice of Others
 */

// @lc code=start
func dominantIndex(nums []int) int {
	res := -1
	max1 := -math.MaxInt32
	max2 := -math.MaxInt32

	for i := 0; i < len(nums); i++ {
		if nums[i] > max1 {
			max2 = max1
			max1 = nums[i]
			if max1 >= 2*max2 {
				res = i
			} else {
				res = -1
			}
		} else if nums[i] > max2 {
			max2 = nums[i]
			if max1 < 2*max2 {
				res = -1
			}
		}
	}

	return res

}

// @lc code=end

