/*
 * @lc app=leetcode id=324 lang=golang
 *
 * [324] Wiggle Sort II
 */

// @lc code=start
func wiggleSort(nums []int)  {
	sorted := make([]int, len(nums))
	copy(sorted, nums)

	sort.Ints(sorted)

	n := len(nums)

	mid := (n - 1) / 2
	right := n - 1

	for i := 0; i < n; i++ {
		if i%2 == 0 {
			nums[i] = sorted[mid]
			mid--
		} else {
			nums[i] = sorted[right]
			right--
		}
	}
}
// @lc code=end


