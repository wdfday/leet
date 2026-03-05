/*
 * @lc app=leetcode id=41 lang=golang
 *
 * [41] First Missing Positive
 */

// @lc code=start
func firstMissingPositive(nums []int) int {
	n := len(nums)

	// Bước 0: Check có số 1 không, không có thì return 1 ngay
	has1 := false
	for i := 0; i < n; i++ {
		if nums[i] == 1 {
			has1 = true
			break
		}
	}
	if !has1 {
		return 1
	}

	// Bước 1: Replace số <= 0 hoặc > n thành 1
	for i := 0; i < n; i++ {
		if nums[i] <= 0 || nums[i] > n {
			nums[i] = 1
		}
	}

	// Bước 2: Marking - đổi dấu số tại index k-1 nếu có số k
	for i := 0; i < n; i++ {
		val := abs(nums[i])
		idx := val - 1
		if nums[idx] > 0 {
			nums[idx] = -nums[idx]
		}
	}

	// Bước 3: Tìm index đầu tiên còn dương
	for i := 0; i < n; i++ {
		if nums[i] > 0 {
			return i + 1
		}
	}

	return n + 1
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

// @lc code=end

// Solution 2: Swap approach (slower but more intuitive)
// func firstMissingPositive(nums []int) int {
// 	n := len(nums)
//
// 	// Bước 1: Swap mỗi số về đúng vị trí của nó
// 	// nums[i] = k thì nên ở index k-1
// 	for i := 0; i < n; i++ {
// 		// Swap bừa cho đến khi không swap được
// 		for nums[i] > 0 && nums[i] <= n && nums[nums[i]-1] != nums[i] {
// 			// Swap nums[i] về vị trí nums[i]-1
// 			targetIdx := nums[i] - 1
// 			nums[i], nums[targetIdx] = nums[targetIdx], nums[i]
// 		}
// 	}
//
// 	// Bước 2: Tìm index đầu tiên không có đúng giá trị
// 	for i := 0; i < n; i++ {
// 		if nums[i] != i+1 {
// 			return i + 1
// 		}
// 	}
//
// 	// Nếu đủ hết 1→n, thì missing là n+1
// 	return n + 1
// }


