// Package http http
// file create by daihao, time is 2018/7/30 15:02
package main

import (
	"strconv"
	"fmt"
)

// main main
func main() {
	// TODO
	b := make([]byte, 0, 100)
	b = strconv.AppendInt(b, int64(8*14-1), 16)
	fmt.Println(string(b))
}
