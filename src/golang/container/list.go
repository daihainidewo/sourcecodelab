// Package main main
// file create by daihao, time is 2018/7/27 12:19
package main

import (
	"container/list"
	"fmt"
)

func main() {
	l := list.New()
	l.Init()
	l.PushBack("abcd")
	l.PushBack("qwer")
	fmt.Println(l.Len())
	fmt.Println(l.Front())
	fmt.Println(l.Back())
	fmt.Println(l.Remove(l.Back()))
	fmt.Println(l)
}
