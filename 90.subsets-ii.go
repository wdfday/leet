/*
 * @lc app=leetcode id=90 lang=golang
 *
 * [90] Subsets II
 */

// @lc code=start
func subsetsWithDup(nums []int) [][]int {

	res := [][]int{}
	sort.Ints(nums)

	var subset func(cur []int, start int)
	subset = func(cur []int, start int) {
		res = append(res, append([]int{}, cur...))

		seen := make(map[int]bool)
		for i := start; i < len(nums); i++ {
			if seen[nums[i]] {
				continue
			}
			seen[nums[i]] = true

			cur = append(cur, nums[i])
			subset(cur, i+1)
			cur = cur[:len(cur)-1]
		}
	}

	subset([]int{}, 0)
	return res
}

// @lc code=end

