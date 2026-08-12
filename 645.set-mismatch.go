/*
 * @lc app=leetcode id=645 lang=golang
 *
 * [645] Set Mismatch
 */

// @lc code=start
func findErrorNums(nums []int) []int {
	count := make([]int, len(nums)+1)
	for _, v := range nums {
		count[v]++
	}

	dup, missing := -1, -1
	for i := 1; i <= len(nums); i++ {
		if count[i] == 2 {
			dup = i
		} else if count[i] == 0 {
			missing = i
		}
	}
	return []int{dup, missing}
}

// @lc code=end

