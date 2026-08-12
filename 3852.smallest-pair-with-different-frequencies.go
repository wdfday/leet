/*
 * @lc app=leetcode id=3852 lang=golang
 *
 * [3852] Smallest Pair With Different Frequencies
 */

// @lc code=start
func minDistinctFreqPair(nums []int) []int {
	freqMap := make(map[int]int)
	for _, n := range nums {
		freqMap[n]++
	}

	freqs := make([]int, 0, len(freqMap))
	for num, _ := range freqMap {
		freqs = append(freqs, num)
	}
	if len(freqs) < 2 {
		return []int{-1, -1}
	}
	sort.Ints(freqs)

	for j := 1; j < len(freqs); j++ {
		if freqMap[freqs[j]] != freqMap[freqs[0]] {
			return []int{freqs[0], freqs[j]}
		}
	}

	return []int{-1, -1}
}

// @lc code=end

