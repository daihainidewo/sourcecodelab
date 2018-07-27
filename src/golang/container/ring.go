// Package main main
// file create by daihao, time is 2018/7/27 14:39
package main

import (
	"container/ring"
	"fmt"
)

func main() {
	arr := [...]int{1, 2, 3, 4, 5, 6}
	r := ring.New(len(arr))
	fmt.Println(r.Len())
	rt := r
	for i := 0; i < len(arr); i++ {
		rt.Value = arr[i]
		rt = rt.Next()
	}
	r.Do(func(i interface{}) {
		fmt.Println(i.(int) * 2)
	})
	r1 := ring.New(1)
	r1.Value = 123
	r.Link(r1)
	r.Do(func(i interface{}) {
		fmt.Println(i)
	})

}
