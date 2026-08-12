/*
 * @lc app=leetcode id=485 lang=golang
 *
 * [485] Max Consecutive Ones
 */

// @lc code=start
func findMaxConsecutiveOnes(nums []int) int {
	maxLength := 0
	consec := 0
	for _, v := range nums {
		if v == 1 {
			consec++
			if consec > maxLength {
				maxLength = consec
			}
		} else {
			consec = 0
		}
	}

	return maxLength

}

// @lc code=end

