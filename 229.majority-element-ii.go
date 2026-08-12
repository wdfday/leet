/*
 * @lc app=leetcode id=229 lang=golang
 *
 * [229] Majority Element II
 */

// @lc code=start
func majorityElement(nums []int) []int {
	n := len(nums)
	if n == 0 {
		return []int{}
	}

	var cand1, cand2 int
	count1, count2 := 0, 0

	for _, x := range nums {
		if count1 > 0 && x == cand1 {
			count1++
		} else if count2 > 0 && x == cand2 {
			count2++
		} else if count1 == 0 {
			cand1 = x
			count1 = 1
		} else if count2 == 0 {
			cand2 = x
			count2 = 1
		} else {
			count1--
			count2--
		}
	}

	// verify
	count1, count2 = 0, 0
	for _, x := range nums {
		if x == cand1 {
			count1++
		} else if x == cand2 {
			count2++
		}
	}

	res := make([]int, 0, 2)
	if count1 > n/3 {
		res = append(res, cand1)
	}
	if cand2 != cand1 && count2 > n/3 {
		res = append(res, cand2)
	}

	return res

}

// @lc code=end

