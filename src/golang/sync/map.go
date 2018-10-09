// Package sync sync
// file create by daihao, time is 2018/10/8 17:47
package main

import (
	"fmt"
	"sort"
	"strconv"
	"sync"
	"time"
)

func main() {
	m := sync.Map{}
	for i := 0; i < 1e4; i++ {
		go func(i int) {
			m.Store(i, "a"+strconv.Itoa(i))
		}(i)
	}
	time.Sleep(1000 * time.Millisecond)
	type tmp struct {
		key   interface{}
		value interface{}
	}
	s := make([]tmp, 0, 1e5)
	//s := make(map[interface{}]interface{})
	m.Range(func(key, value interface{}) bool {
		s = append(s, tmp{key, value})
		return true
	})
	sort.Slice(s, func(i, j int) bool {
		if s[i].key.(int) < s[j].key.(int) {
			return true
		}
		return false
	})
	for i, d := range s {
		if i%1000 == 0 {
			fmt.Println(d)
			continue
		}
		fmt.Print(d)
	}
	//fmt.Println(s)
	//fmt.Println(m)
}
