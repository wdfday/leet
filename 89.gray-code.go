/*
 * @lc app=leetcode id=89 lang=golang
 *
 * [89] Gray Code
 */

// @lc code=start
func grayCode(n int) []int {
    total := 1 << n
    ret := make([]int, total)
    for i := 0; i < total; i++ {
        ret[i] = i ^ (i >> 1)
    }    
    return ret
}
// @lc code=end

