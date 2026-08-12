/*
 * @lc app=leetcode id=922 lang=golang
 *
 * [922] Sort Array By Parity II
 */

// @lc code=start
func sortArrayByParityII(nums []int) []int {
	for i :=0; i < len(nums); i++ {
		if isOdd(i) && isEven(nums[i]) {
			for j := i+1; j < len(nums); j+=2 {
				if isOdd(nums[j]) {
					nums[i], nums[j] = nums[j], nums[i]
					break
				}
			}
		} else if isEven(i) && isOdd(nums[i]) {
			for j := i+1; j < len(nums); j+=2 {
				if isEven(nums[j]) {
					nums[i], nums[j] = nums[j], nums[i]
					break
				}
			}
		}
	}
    return nums
}

func isEven(i int) bool {
	return i %2 ==0 
}

func isOdd(i int) bool {
	return i %2 == 1
}
// @lc code=end

