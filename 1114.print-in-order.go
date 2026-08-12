/*
 * @lc app=leetcode id=1114 lang=golang
 *
 * [1114] Print in Order
 */

// @lc code=start
type Foo struct {
	ch12 chan struct{}
	ch23 chan struct{}
}

func NewFoo() *Foo {
	return &Foo{
		ch12: make(chan struct{}),
		ch23: make(chan struct{}),
	}
}

func (f *Foo) First(printFirst func()) {
	// Do not change this line
	printFirst()
	close(f.ch12)
}

func (f *Foo) Second(printSecond func()) {
	<-f.ch12
	/// Do not change this line
	printSecond()
	close(f.ch23)
}

func (f *Foo) Third(printThird func()) {
	<-f.ch23
	// Do not change this line
	printThird()
}
// @lc code=end

