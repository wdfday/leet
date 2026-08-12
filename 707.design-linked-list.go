/*
 * @lc app=leetcode id=707 lang=golang
 *
 * [707] Design Linked List
 */

// @lc code=start
type MyLinkedList struct {
	head *MyLinkedList // dummy head, val không dùng
	next *MyLinkedList
	val  int
	size int
}

func Constructor() MyLinkedList {
	return MyLinkedList{size: 0, next: nil}
}

// helper: trả về node đứng TRƯỚC index (index=-1 -> trả về chính dummy head)
func (this *MyLinkedList) nodeBefore(index int) *MyLinkedList {
	cur := this
	for i := 0; i < index; i++ {
		cur = cur.next
	}
	return cur
}

func (this *MyLinkedList) Get(index int) int {
	if index < 0 || index >= this.size {
		return -1
	}
	return this.nodeBefore(index).next.val
}

func (this *MyLinkedList) AddAtHead(val int) {
	this.AddAtIndex(0, val)
}

func (this *MyLinkedList) AddAtTail(val int) {
	this.AddAtIndex(this.size, val)
}

func (this *MyLinkedList) AddAtIndex(index int, val int) {
	if index > this.size {
		return
	}
	if index < 0 {
		index = 0
	}
	prev := this.nodeBefore(index)
	newNode := &MyLinkedList{val: val, next: prev.next}
	prev.next = newNode
	this.size++
}

func (this *MyLinkedList) DeleteAtIndex(index int) {
	if index < 0 || index >= this.size {
		return
	}
	prev := this.nodeBefore(index)
	prev.next = prev.next.next
	this.size--
}

/**
 * Your MyLinkedList object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Get(index);
 * obj.AddAtHead(val);
 * obj.AddAtTail(val);
 * obj.AddAtIndex(index,val);
 * obj.DeleteAtIndex(index);
 */
// @lc code=end

