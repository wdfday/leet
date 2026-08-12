/*
 * @lc app=leetcode id=1 lang=golang
 *
 * [1] Two Sum
 */

// @lc code=start
func twoSum(nums []int, target int) []int {

	m := make(map[int]int)

	for i := 0; i < len(nums); i++ {
		if k, ok := m[target-nums[i]]; ok {
			return []int{k, i}
		}
		m[nums[i]] = i
	}

	return nil

}

// @lc code=end

