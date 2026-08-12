/*
 * @lc app=leetcode id=706 lang=golang
 *
 * [706] Design HashMap
 */

// @lc code=start
type entry struct {
	key, value int
	next       *entry
}

const numBuckets = 1000

type MyHashMap struct {
	buckets [numBuckets]*entry
}

func Constructor() MyHashMap {
	return MyHashMap{}
}

func (this *MyHashMap) hash(key int) int {
	return key % numBuckets
}

func (this *MyHashMap) Put(key int, value int) {
	idx := this.hash(key)
	// nếu key đã tồn tại thì update value
	for e := this.buckets[idx]; e != nil; e = e.next {
		if e.key == key {
			e.value = value
			return
		}
	}
	// chưa có thì thêm node mới vào đầu bucket
	newEntry := &entry{key: key, value: value, next: this.buckets[idx]}
	this.buckets[idx] = newEntry
}

func (this *MyHashMap) Get(key int) int {
	idx := this.hash(key)
	for e := this.buckets[idx]; e != nil; e = e.next {
		if e.key == key {
			return e.value
		}
	}
	return -1
}

func (this *MyHashMap) Remove(key int) {
	idx := this.hash(key)
	var prev *entry
	for e := this.buckets[idx]; e != nil; e = e.next {
		if e.key == key {
			if prev == nil {
				this.buckets[idx] = e.next
			} else {
				prev.next = e.next
			}
			return
		}
		prev = e
	}
}

/**
 * Your MyHashMap object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Put(key,value);
 * param_2 := obj.Get(key);
 * obj.Remove(key);
 */
// @lc code=end

