/*
 * @lc app=leetcode id=380 lang=golang
 *
 * [380] Insert Delete GetRandom O(1)
 */

// @lc code=start
type RandomizedSet struct {
	nums []int
	idx  map[int]int
}

func Constructor() RandomizedSet {
	return RandomizedSet{
		nums: make([]int, 0),
		idx:  make(map[int]int),
	}

}

func (this *RandomizedSet) Insert(val int) bool {
	if _, exists := this.idx[val]; exists {
		return false
	}
	this.idx[val] = len(this.nums)
	this.nums = append(this.nums, val)
	return true
}

func (this *RandomizedSet) Remove(val int) bool {
	removeIdx, exists := this.idx[val]
	if !exists {
		return false
	}

	lastIdx := len(this.nums) - 1
	if removeIdx != lastIdx {
		lastVal := this.nums[lastIdx]
		this.nums[removeIdx] = lastVal
		this.idx[lastVal] = removeIdx
	}

	this.nums = this.nums[:lastIdx]
	delete(this.idx, val)
	return true
}

func (this *RandomizedSet) GetRandom() int {
	return this.nums[rand.Intn(len(this.nums))]
}

/**
 * Your RandomizedSet object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Insert(val);
 * param_2 := obj.Remove(val);
 * param_3 := obj.GetRandom();
 */
// @lc code=end

