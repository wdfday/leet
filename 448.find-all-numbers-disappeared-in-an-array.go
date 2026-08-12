/*
 * @lc app=leetcode id=448 lang=golang
 *
 * [448] Find All Numbers Disappeared in an Array
 */

// @lc code=start
func findDisappearedNumbers(nums []int) []int {
	cp := make([]bool, len(nums)+1)
	for i := 0; i < len(nums); i++ {
		cp[nums[i]] = true
	}
	res := make([]int, 0)

	for i := 1; i < len(cp); i++ {
		if cp[i] == false {
			res = append(res, i)
		}
	}
	return res
}

// @lc code=end

