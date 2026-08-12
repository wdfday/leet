/*
 * @lc app=leetcode id=442 lang=golang
 *
 * [442] Find All Duplicates in an Array
 */

// @lc code=start
func findDuplicates(nums []int) []int {
	for i := 0; i < len(nums); {
		correct := nums[i] - 1

		if nums[i] != nums[correct] {
			nums[i], nums[correct] = nums[correct], nums[i]
		} else {
			i++
		}
	}

	res := []int{}
	for i := range nums {
		if nums[i] != i+1 {
			res = append(res, nums[i])
		}
	}

	return res
}

// @lc code=end

func findDuplicates(nums []int) []int {
	res := []int{}

	for _, x := range nums {
		idx := abs(x) - 1

		if nums[idx] < 0 {
			res = append(res, abs(x))
		} else {
			nums[idx] = -nums[idx]
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