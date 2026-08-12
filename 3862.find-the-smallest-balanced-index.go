/*
 * @lc app=leetcode id=3862 lang=golang
 *
 * [3862] Find the Smallest Balanced Index
 */

// @lc code=start

// func smallestBalancedIndex(nums []int) int {
// 	n := len(nums)
// 	if n < 2 {
// 		return -1
// 	}

// 	const CAP = int(2e18) // ngưỡng chặn để tránh tràn số

// 	leftSum := make([]int, n)
// 	for i := 1; i < n; i++ {
// 		leftSum[i] = leftSum[i-1] + nums[i-1]
// 	}

// 	rightProd := make([]int, n)
// 	rightProd[n-1] = 1
// 	for i := n - 2; i >= 0; i-- {
// 		if rightProd[i+1] > CAP/nums[i+1] { // check tràn số trước khi nhân
// 			rightProd[i] = CAP + 1 // đánh dấu "quá lớn", chắc chắn không thể == leftSum
// 		} else {
// 			rightProd[i] = rightProd[i+1] * nums[i+1]
// 		}
// 	}

//		for i := 0; i < n; i++ {
//			if leftSum[i] == rightProd[i] {
//				return i
//			}
//		}
//		return -1
//	}
func smallestBalancedIndex(nums []int) int {
	if len(nums) < 2 {
		return -1
	}
	l, r := 0, len(nums)-1
	leftSum, rightProd := 0, 1
	const CAP = int(1 << 62) // ngưỡng chặn, đủ lớn để không nhầm với leftSum thật

	for l < r {
		if leftSum < rightProd {
			leftSum += nums[l]
			l++
		} else {
			if rightProd > CAP/nums[r] { // sắp tràn số nếu nhân tiếp
				rightProd = CAP
			} else {
				rightProd *= nums[r]
			}
			r--
		}
	}

	if leftSum == rightProd {
		return l
	}
	return -1
}

// @lc code=end

