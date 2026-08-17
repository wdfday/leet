/*
 * @lc app=leetcode id=347 lang=golang
 *
 * [347] Top K Frequent Elements
 */

// @lc code=start
func topKFrequent(nums []int, k int) []int {
	res := []int{}

	m := make(map[int]int)

	for _, num := range nums {
		m[num]++
	}

	h := &Heap{}

	for k, v := range m {
		heap.Push(h, Element{
			Value: k,
			Freq:  v,
		},
		)
	}

	for _ = range k {
		e := heap.Pop(h).(Element)
		res = append(res, e.Value)
	}

	return res

}


type Element struct {
	Value int
	Freq  int
}

type Heap []Element

func (h Heap) Len() int           { return len(h) }
func (h Heap) Less(i, j int) bool { return h[i].Freq > h[j].Freq }
func (h Heap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *Heap) Push(x any) {
	// Push and Pop use pointer receivers because they modify the slice's length,
	// not just its contents.
	*h = append(*h, x.(Element))
}

func (h *Heap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

// @lc code=end
func topKFrequent(nums []int, k int) []int {
	freq := make(map[int]int)

	for _, num := range nums {
		freq[num]++
	}

	bucket := make([][]int, len(nums)+1)

	for num, count := range freq {
		bucket[count] = append(bucket[count], num)
	}

	res := make([]int, 0, k)

	for count := len(bucket) - 1; count >= 1 && len(res) < k; count-- {
		for _, num := range bucket[count] {
			res = append(res, num)

			if len(res) == k {
				return res
			}
		}
	}

	return res
}