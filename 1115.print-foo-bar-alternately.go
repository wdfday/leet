/*
 * @lc app=leetcode id=1115 lang=golang
 *
 * [1115] Print FooBar Alternately
 */

// @lc code=start
type FooBar struct {
	n int
}

func NewFooBar(n int) *FooBar {
	return &FooBar{n: n}
}

func (fb *FooBar) Foo(printFoo func()) {
	for i := 0; i < fb.n; i++ {
		// printFoo() outputs "foo". Do not change or remove this line.
        printFoo()
	}
}

func (fb *FooBar) Bar(printBar func()) {
	for i := 0; i < fb.n; i++ {
		// printBar() outputs "bar". Do not change or remove this line.
        printBar()
	}
}
// @lc code=end

