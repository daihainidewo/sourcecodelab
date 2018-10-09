// Package main main
// file create by daihao, time is 2018/10/9 11:52
package main

import (
	"fmt"
	"sync"
)

func main() {
	p := sync.Pool{}
	p.New = func() interface{} {
		return 0
	}
	p.Put(1)
	fmt.Println(p.Get().(int))
	fmt.Println(p.Get().(int))
	fmt.Println(p.Get().(int))
	fmt.Println(p.Get().(int))
	fmt.Println(p.Get().(int))
}
