/*
 * @lc app=leetcode id=3678 lang=golang
 *
 * [3678] Smallest Absent Positive Greater Than Average
 */

// @lc code=start
func smallestAbsent(nums []int) int {
	m := make(map[int]bool)
	sum := 0
	for _, v := range nums {
		sum += v 
		if v > 0 {
			m[v] = true
		}
	} 
	sum /= len(nums)
	for i := max(1, sum+1); ;i++{
		if m[i] == false {
			return i
		}
	}
	return 1
}
// @lc code=end

