package main

import "container/heap"

/*
 * @lc app=leetcode id=295 lang=golang
 *
 * [295] Find Median from Data Stream
 */

// @lc code=start
type MaxHeap []int

func (h MaxHeap) Len() int            { return len(h) }
func (h MaxHeap) Less(i, j int) bool  { return h[i] > h[j] } // đảo dấu để thành max-heap
func (h MaxHeap) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *MaxHeap) Push(x interface{}) { *h = append(*h, x.(int)) }
func (h *MaxHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}

type MinHeap []int

func (h MinHeap) Len() int            { return len(h) }
func (h MinHeap) Less(i, j int) bool  { return h[i] < h[j] }
func (h MinHeap) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *MinHeap) Push(x interface{}) { *h = append(*h, x.(int)) }
func (h *MinHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}

type MedianFinder struct {
	maxHeap *MaxHeap 
	minHeap *MinHeap 
}

func Constructor() MedianFinder {
	return MedianFinder{
		maxHeap: &MaxHeap{},
		minHeap: &MinHeap{},
	}
}

func (this *MedianFinder) AddNum(num int) {
	heap.Push(this.maxHeap, num)
	heap.Push(this.minHeap, heap.Pop(this.maxHeap))

	if this.minHeap.Len() > this.maxHeap.Len() {
		heap.Push(this.maxHeap, heap.Pop(this.minHeap))
	}
}

func (this *MedianFinder) FindMedian() float64 {
	if this.maxHeap.Len() > this.minHeap.Len() {
		return float64((*this.maxHeap)[0])
	}
	return float64((*this.maxHeap)[0]+(*this.minHeap)[0]) / 2.0
}
/**
 * Your MedianFinder object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AddNum(num);
 * param_2 := obj.FindMedian();
 */
// @lc code=end

