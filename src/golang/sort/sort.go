// Package sort sort
// file create by daihao, time is 2018/7/27 10:20
package main

import (
	"sort"
	"fmt"
)

// kv
type kv struct {
	Key   int
	Value string
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

// Sort Sort
func (k *kvslice) Sort() {
	// TODO
	sort.Sort(k)
}

// String String
func (k *kvslice) String() string {
	// TODO
	str := ""
	str += fmt.Sprintf("size:%d, cap:%d, slice:[", k.size, k.cap)
	for i, d := range k.Sli {
		if i != k.size {
			str += fmt.Sprintf("{Key:%d, Value:%s}", d.Key, d.Value)
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

// main main
func main() {
	// TODO
	a := NewKvslice(5)
	a.Set(2, "123")
	a.Set(3, "234")
	a.Set(1, "345")
	a.Set(9, "456")
	a.Set(8, "567")
	a.Set(7, "678")
	fmt.Println(a)
	a.Sort()
	fmt.Println(a)
}
