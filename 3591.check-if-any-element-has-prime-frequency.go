/*
 * @lc app=leetcode id=3591 lang=golang
 *
 * [3591] Check if Any Element Has Prime Frequency
 */

// @lc code=start
func checkPrimeFrequency(nums []int) bool {
	freq := make(map[int]int)
	for _, n := range nums {
		freq[n]++
	}

	isPrime := func(n int) bool {
		if n < 2 {
			return false
		}
		for i := 2; i*i <= n; i++ {
			if n%i == 0 {
				return false
			}
		}
		return true
	}

	for _, count := range freq {
		if isPrime(count) {
			return true
		}
	}
	return false
}
// @lc code=end

