/*
 * @lc app=leetcode id=284 lang=golang
 *
 * [284] Peeking Iterator
 */

// @lc code=start
/*   Below is the interface for Iterator, which is already defined for you.
 *
 *   type Iterator struct {
 *       
 *   }
 *
 *   func (this *Iterator) hasNext() bool {
 *		// Returns true if the iteration has more elements.
 *   }
 *
 *   func (this *Iterator) next() int {
 *		// Returns the next element in the iteration.
 *   }
 */

type PeekingIterator struct {
    iter *Iterator
	p int
}

func Constructor(iter *Iterator) *PeekingIterator {
	if iter == nil {
		return nil
	}
	var p int
	if iter.hasNext() {
		p = iter.next()
	}
    return &PeekingIterator{
		iter: iter,
		p: p,
	}
}

func (this *PeekingIterator) hasNext() bool {
    if this.p == 0 {
		return false
	}
	return true
}

func (this *PeekingIterator) next() int {
	res := this.p
	if this.iter.hasNext(){
		this.p = this.iter.next()
	} else {
		this.p = 0
	}
    return res
}

func (this *PeekingIterator) peek() int {
    return this.p
}
// @lc code=end

