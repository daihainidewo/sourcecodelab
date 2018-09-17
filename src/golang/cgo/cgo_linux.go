// +build linux
// Package main main
// file create by daihao, time is 2018/9/4 11:56
package main

/*
#include "main.c"
*/
import "C"
// #cgo CFLAGS: -I .
// #cgo LDFLAGS: -L . -lfoo
import "fmt"

func main() {
	p := C.Hello()
	i := 0
	for *C.char((C.int(p) + C.int(i))) != C.char(0) {
		fmt.Print(C.char(*p))
		i++
	}
}
