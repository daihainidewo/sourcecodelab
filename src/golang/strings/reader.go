// Package main main
// file create by daihao, time is 2018/7/27 18:13
package main

import (
	"strings"
	"fmt"
)

// main main
func main() {
	// TODO
	r := strings.NewReader("abc")
	fmt.Println(r.Len(), r.Size())
	b := []byte("1234")
	row, _ := r.Read(b)
	fmt.Println(row, string(b))
	row, _ = r.ReadAt(b, 1)
	fmt.Println(row, string(b))
	c, _ := r.ReadByte()
	fmt.Println(c, string(c))
}
