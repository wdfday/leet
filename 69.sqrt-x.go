// package main

// import "fmt"

/*
 * @lc app=leetcode id=69 lang=golang
 *
 * [69] Sqrt(x)
 */

// @lc code=start
func mySqrt(x int) int {
	l, r := 0, x
	for l < r {
		mid := (l + r + 1) / 2
		if mid > x/mid {
			r = mid - 1
		} else {
			l = mid
		}
	}
	return l
}

// @lc code=end

// func main() {
// 	fmt.Println(mySqrt(10))
// }
