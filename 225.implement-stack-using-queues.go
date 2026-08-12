/*
 * @lc app=leetcode id=225 lang=golang
 *
 * [225] Implement Stack using Queues
 */

// @lc code=start
type MyStack struct {
	data []int
}

func Constructor() MyStack {
	return MyStack{data: []int{}}
}

func (this *MyStack) Push(x int) {
	this.data = append(this.data, x)
}

func (this *MyStack) Pop() int {
	n := len(this.data)
	if n == 0 {
		return 0
	}
	val := this.data[n-1]
	this.data = this.data[:n-1]
	return val
}

func (this *MyStack) Top() int {
	n := len(this.data)
	if n == 0 {
		return 0
	}
	return this.data[n-1]
}

func (this *MyStack) Empty() bool {
	return len(this.data) == 0
}

/**
 * Your MyStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.Empty();
 */
// @lc code=end

