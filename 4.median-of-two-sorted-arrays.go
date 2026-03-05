package go
/*
 * @lc app=leetcode id=4 lang=golang
 *
 * [4] Median of Two Sorted Arrays
 */

// @lc code=start
func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {

	x, y := len(nums1), len(nums2)

	// ensure len nums1 < len nums2
	if x > y {
		return findMedianSortedArrays(nums2, nums1)
	}

	low, high := 0, x

	// left == right || == right + 1
	totalLeft := (x + y + 1) / 2

	for low <= high {
		partition1 := (low + high) / 2
		partition2 := totalLeft - partition1

		var maxLeft1 int
		if partition1 == 0 {
			maxLeft1 = math.MinInt64
		} else {
			maxLeft1 = nums1[partition1 - 1]
		}

		var minRight1 int
		if partition1 == x {
			minRight1 = math.MaxInt64
		} else {
			minRight1 = nums1[partition1]
		}

		var maxLeft2 int
		if partition2 == 0 {
			maxLeft2 = math.MinInt64
		} else {
			maxLeft2 = nums2[partition2 - 1]
		}

		var minRight2 int
		if partition2 == y {
			minRight2 = math.MaxInt64
		} else {
			minRight2 = nums2[partition2]
		}

		// check if we have found the correct partition
		if maxLeft1 <= minRight2 && maxLeft2 <= minRight1 {
			if (x + y) % 2 == 0 {
				return float64(max(maxLeft1, maxLeft2) + min(minRight1, minRight2)) / 2.0
			} else {
				return float64(max(maxLeft1, maxLeft2))
			}
		} else if maxLeft1 > minRight2 {
			high = partition1 - 1
		} else {
			low = partition1 + 1
		}
	}

	return 0.0

}

// @lc code=end

