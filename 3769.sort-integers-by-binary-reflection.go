/*
 * @lc app=leetcode id=3769 lang=golang
 *
 * [3769] Sort Integers by Binary Reflection
 */

// @lc code=start
type pair struct {
	val, refl int
}

func sortByReflection(nums []int) []int {
	pairs := make([]pair, len(nums))
	for i, v := range nums {
		pairs[i] = pair{v, reflect(v)}
	}

	sort.Slice(pairs, func(i, j int) bool {
		if pairs[i].refl != pairs[j].refl {
			return pairs[i].refl < pairs[j].refl
		}
		return pairs[i].val < pairs[j].val
	})

	res := make([]int, len(nums))
	for i, p := range pairs {
		res[i] = p.val
	}
	return res
}
func reflect(n int) int {
	res := 0
	for n > 0 {
		res = res<<1 | (n & 1)
		n >>= 1
	}
	return res

}

// @lc code=end

