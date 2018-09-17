// +build linux
// Package doc_tree doc_tree
// file create by daihao, time is 2018/8/27 14:50
package main

import (
	"os"
	"syscall"
)

// GetFileTime GetFileTime
func GetFileTime(info os.FileInfo) (ct int64, lwt int64, lat int64) {
	// TODO
	info, _ := os.Stat(path)
	filesys := info.Sys().(*syscall.Stat_t)
	ct = filesys.Ctim.Nano()
	lwt = filesys.Mtim.Nano()
	lat = filesys.Atim.Nano()
}
