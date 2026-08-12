/*
 * @lc app=leetcode id=84 lang=golang
 *
 * [84] Largest Rectangle in Histogram
 */

// @lc code=start


// largestRectangleAreaStack: O(n) time, O(n) space using a monotonic increasing stack.
// stack holds indices with increasing heights[idx]. When a shorter bar arrives,
// pop and finalize each taller bar: its width spans from the new top of the
// stack (exclusive) to the current index (exclusive).
func largestRectangleArea(heights []int) int {
	stack := make([]int, 0, len(heights))
	res := 0
	for i := 0; i <= len(heights); i++ {
		h := 0
		if i < len(heights) {
			h = heights[i]
		}
		for len(stack) > 0 && heights[stack[len(stack)-1]] >= h {
			top := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			width := i
			if len(stack) > 0 {
				width = i - stack[len(stack)-1] - 1
			}
			if area := heights[top] * width; area > res {
				res = area
			}
		}
		stack = append(stack, i)
	}
	return res
}

// @lc code=end



// func largestRectangleArea(heights []int) int {
// 	m := len(heights)
// 	dp := make([][]int, m)
// 	for i := range m {
// 		dp[i] = make([]int, m)
// 	}
// 	res := 0
// 	for i := 0; i < m; i++ {
// 		dp[i][i] = heights[i]
// 		if dp[i][i] > res {
// 			res =  dp[i][i] 
// 		}
// 		for j := i + 1; j < m; j++ {
// 			h := min(dp[i][j-1]/(j-i), heights[j])
// 			dp[i][j] = (j+1-i) * h
// 			if dp[i][j] > res {
// 				res = dp[i][j]
// 			}
// 		}
// 	}
// 	return res
// }

// largestRectangleAreaDP1D: same O(n^2) time recurrence as largestRectangleArea,
// // but O(n) space. Row i of the original 2D dp only ever reads dp[i][j-1] from
// // the same row, never row i-1, so a single array can be overwritten in place
// // on every outer iteration.
// func largestRectangleArea(heights []int) int {
// 	m := len(heights)
// 	dp := make([]int, m)
// 	res := 0
// 	for i := 0; i < m; i++ {
// 		dp[i] = heights[i]
// 		if dp[i] > res {
// 			res = dp[i]
// 		}
// 		for j := i + 1; j < m; j++ {
// 			h := min(dp[j-1]/(j-i), heights[j])
// 			dp[j] = (j + 1 - i) * h
// 			if dp[j] > res {
// 				res = dp[j]
// 			}
// 		}
// 	}
// 	return res
// }