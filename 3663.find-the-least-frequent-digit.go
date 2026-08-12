/*
 * @lc app=leetcode id=3663 lang=golang
 *
 * [3663] Find The Least Frequent Digit
 */

// @lc code=start
func getLeastFrequentDigit(n int) int {
	freq := make(map[int]int)

	for n > 0 {
		k := n % 10
		n /= 10
		freq[k]++
	}

	res := 10
	f := 100
	for k, v := range freq {
		if v < f || (v == f && res > k) {
			res = k
			f = v
		}
	}
	return res
}

// @lc code=end

