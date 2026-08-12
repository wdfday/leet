/*
 * @lc app=leetcode id=762 lang=golang
 *
 * [762] Prime Number of Set Bits in Binary Representation
 */

// @lc code=start
func countPrimeSetBits(left int, right int) int {
	primes := map[int]bool{
		2: true, 3: true, 5: true, 7: true,
		11: true, 13: true, 17: true, 19: true,
	}

	count := 0
	for n := left; n <= right; n++ {
		if primes[bits.OnesCount(uint(n))] {
			count++
		}
	}
	return count
}

// @lc code=end

