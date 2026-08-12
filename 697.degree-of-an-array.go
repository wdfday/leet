/*
 * @lc app=leetcode id=697 lang=golang
 *
 * [697] Degree of an Array
 */

// @lc code=start
func findShortestSubArray(nums []int) int {
	m := make(map[int][]int)
	for i := 0; i < len(nums); i++ {
		m[nums[i]] = append(m[nums[i]], i)
	}
	res := len(nums)
	degree := 0
	for _, v := range m {
		if len(v) > degree {
			degree = len(v)
			res = v[len(v)-1] - v[0] + 1
		} else if len(v) == degree {
			res = min(res, v[len(v)-1]-v[0]+1)
		}
	}
	return res

}

// @lc code=end

