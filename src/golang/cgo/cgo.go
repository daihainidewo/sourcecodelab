
// +build windows

// Package cgo cgo
// file create by daihao, time is 2018/7/30 17:05
package main

/*
#include <windows.h>
*/
import "C"

func Msgbox(body, title string) int {
	C.MessageBox(nil, (*C.CHAR)(C.CString(body)), (*C.CHAR)(C.CString(title)), 0)
	return 0
}

func main() {
	Msgbox("title", "body")
}
