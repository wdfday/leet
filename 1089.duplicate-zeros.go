/*
 * @lc app=leetcode id=1089 lang=golang
 *
 * [1089] Duplicate Zeros
 */

// @lc code=start
func duplicateZeros(arr []int)  {

	for i := 0; i < len(arr); i++ {
		if arr[i] == 0 {
			cp := append(arr[:i], 0)
			cp = append(cp, arr[i:]...)
			copy(arr, cp)
			i++
		}
	}
    
}
// @lc code=end

