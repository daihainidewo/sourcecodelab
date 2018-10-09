// Package atomic atomic
// file create by daihao, time is 2018/10/9 15:34
package main

import (
	"fmt"
	"sync/atomic"
	"time"
)

func main() {
	v := new(atomic.Value)
	start := time.Now().Add(100 * time.Millisecond)
	i := 0
	for start.UnixNano() > time.Now().UnixNano() {
		go func(i int) {
			v.Store(i)
		}(i)
		i++
	}
	fmt.Println(v.Load())
}
