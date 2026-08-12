/*
 * @lc app=leetcode id=300 lang=golang
 *
 * [300] Longest Increasing Subsequence
 */

// @lc code=start

func lengthOfLIS(nums []int) int {
	tails := []int{}

	for _, x := range nums {
		i := lowerBound(tails, x)

		if i == len(tails) {
			tails = append(tails, x)
		} else {
			tails[i] = x
		}
	}

	return len(tails)
}

func lowerBound(a []int, x int) int {
	l, r := 0, len(a)

	for l < r {
		mid := (l + r) / 2

		if a[mid] >= x {
			r = mid
		} else {
			l = mid + 1
		}
	}

	return l
}

// @lc code=end

func lengthOfLIS(nums []int) int {
	dp := make([]int, len(nums))

	for i, v := range nums {
		m := 1
		for j := 0; j < i; j++ {
			if v > nums[j] {
				m = max(m, dp[j]+1)
			}
		}
		dp[i] = m
	}

	return slices.Max(dp)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}