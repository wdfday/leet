/*
 * @lc app=leetcode id=1122 lang=golang
 *
 * [1122] Relative Sort Array
 */

// @lc code=start
func relativeSortArray(arr1 []int, arr2 []int) []int {
	m := make(map[int]int)
	for i, v := range arr2 {
		m[v] = len(arr2) - i
	}

	slices.SortFunc(arr1, func(a, b int) int {
		if cmp.Compare(m[b],m[a]) != 0 {
			return cmp.Compare(m[b],m[a])
		} 
		return cmp.Compare(a,b)
	})
    return arr1
}
// @lc code=end

