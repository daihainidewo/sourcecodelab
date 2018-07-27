// Package container container
// file create by daihao, time is 2018/7/27 10:15
package main

import (
	"fmt"
	"container/heap"
)

// kv
type kv struct {
	Key   int
	Value string
}

// String String
func (k *kv) String() string {
	// TODO
	return fmt.Sprintf("{Key:%d, Value:%s}", k.Key, k.Value)
}

// kvslice
type kvslice struct {
	Sli  []kv
	cap  int
	size int
}

// NewKvslice new Kcslice
func NewKvslice(cap int) *kvslice {
	// TODO
	return &kvslice{
		cap:  cap,
		size: 0,
		Sli:  make([]kv, cap),
	}
}

// Set Set
func (k *kvslice) Set(key int, v string) {
	// TODO
	if k.size == k.cap {
		tmp := k.Sli
		k.cap = 2 * k.cap
		k.Sli = make([]kv, k.cap)
		for i, d := range tmp {
			k.Sli[i] = d
		}
	}
	k.Sli[k.size] = kv{key, v}
	k.size++
}

// String String
func (k *kvslice) String() string {
	// TODO
	str := ""
	str += fmt.Sprintf("size:%d, cap:%d, slice:[", k.size, k.cap)
	for i, d := range k.Sli {
		if i != k.size {
			str += fmt.Sprint(d)
		} else {
			break
		}
		if i != k.size-1 {
			str += ","
		}
	}
	str += "]"
	return str
}

// Swap Swap
func (k *kvslice) Swap(i, j int) {
	// TODO
	k.Sli[i], k.Sli[j] = k.Sli[j], k.Sli[i]
}

// Len Len
func (k *kvslice) Len() int {
	// TODO
	return k.size
}

// Less Less
func (k *kvslice) Less(i, j int) bool {
	// TODO
	return k.Sli[i].Key < k.Sli[j].Key
}

// Push Push
func (k *kvslice) Push(a interface{}) {
	// TODO
	b := a.(kv)
	k.Set(b.Key, b.Value)
}

// Pop Pop
func (k *kvslice) Pop() interface{} {
	// TODO
	k.size--
	return k.Sli[k.size]
}

// HeapExe HeapExe
func HeapExe() {
	// TODO
	a := NewKvslice(5)
	a.Set(2, "123")
	a.Set(3, "234")
	a.Set(1, "345")
	a.Set(9, "456")
	a.Set(8, "567")
	a.Set(7, "678")
	fmt.Println(a)
	heap.Init(a)
	fmt.Println(a)
	heap.Push(a, kv{111, "11111111"})
	fmt.Println(a)
	b := heap.Pop(a)
	fmt.Println(a)
	fmt.Println(b)
	b = heap.Pop(a)
	fmt.Println(a)
	fmt.Println(b)
	b = heap.Pop(a)
	fmt.Println(a)
	fmt.Println(b)
	b = heap.Pop(a)
	fmt.Println(a)
	fmt.Println(b)
}

func main() {
	HeapExe()
}
