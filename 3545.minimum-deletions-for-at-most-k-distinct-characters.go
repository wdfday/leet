/*
 * @lc app=leetcode id=3545 lang=golang
 *
 * [3545] Minimum Deletions for At Most K Distinct Characters
 */

// @lc code=start
func minDeletion(s string, k int) int {
	freq := make([]int, 26)

	for _, c := range s {
		freq[c-'a']++
	}

	counts := []int{}
	for _, f := range freq {
		if f > 0 {
			counts = append(counts, f)
		}
	}
	sort.Ints(counts)

	deletions := 0
	remove := len(counts) - k

	for i := 0; i < remove; i++ {
		deletions += counts[i]
	}

	return deletions
}
// @lc code=end

