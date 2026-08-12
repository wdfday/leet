/*
 * @lc app=leetcode id=228 lang=golang
 *
 * [228] Summary Ranges
 */

// @lc code=start
func summaryRanges(nums []int) []string {

	res := []string{}
	for i := 0; i < len(nums); i++ {
		start := nums[i]
		for i+1 < len(nums) && nums[i+1] == nums[i]+1 {
			i++
		}

		if start != nums[i] {
			res = append(res, fmt.Sprintf("%d->%d", start, nums[i]))
		} else {
			res = append(res, fmt.Sprintf("%d", start))
		}
	}

	return res
}

// @lc code=end

