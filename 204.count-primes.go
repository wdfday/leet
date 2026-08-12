/*
 * @lc app=leetcode id=204 lang=golang
 *
 * [204] Count Primes
 */

// @lc code=start
func countPrimes(n int) int {
	if n < 3 {
		return 0
	}
	isNotPrime := make([]bool, n)
	res := 0
	for i := 2; i < n; i++ {
		if isNotPrime[i] {
			continue
		}
		res++
		for j := i * i; j < n; j += i {
			isNotPrime[j] = true
		}
	}
	return res
}

// @lc code=end

