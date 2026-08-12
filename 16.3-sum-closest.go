/*
 * @lc app=leetcode id=16 lang=golang
 *
 * [16] 3Sum Closest
 */

// @lc code=start
func threeSumClosest(nums []int, target int) int {

	n := len(nums)
	slices.Sort(nums)

	res := nums[0] + nums[1] + nums[2]

	if n == 3 {
		return res
	}

	for i := 0; i < n-2; i++ {
		l, r := i+1, n-1

		for l < r {
			sum := nums[i] + nums[l] + nums[r]

			if abs(target-res) > abs(target-sum) {
				res = sum
			} else if sum > target {
				r--
			} else {
				l++
			}
		}
	}

	return res

}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

// @lc code=end

