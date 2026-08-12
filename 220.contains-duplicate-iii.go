/*
 * @lc app=leetcode id=220 lang=golang
 *
 * [220] Contains Duplicate III
 */

// @lc code=start
func containsNearbyAlmostDuplicate(nums []int, indexDiff int, valueDiff int) bool {
	if valueDiff < 0 {
		return false
	}

	bucket := make(map[int]int)
	size := valueDiff + 1

	getID := func(x int) int {
		if x < 0 {
			return (x+1)/size - 1
		}
		return x / size
	}

	for i := 0; i < len(nums); i++ {
		id := getID(nums[i])

		if _, ok := bucket[id]; ok {
			return true
		}

		if v, ok := bucket[id-1]; ok && abs(nums[i]-v) <= valueDiff {
			return true
		}

		if v, ok := bucket[id+1]; ok && abs(nums[i]-v) <= valueDiff {
			return true
		}

		bucket[id] = nums[i]

		if i >= indexDiff {
			delete(bucket, getID(nums[i-indexDiff]))
		}
	}

	return false
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

// @lc code=end

