// +build windows

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
	filesys := info.Sys().(*syscall.Win32FileAttributeData)
	ct = filesys.CreationTime.Nanoseconds()
	lwt = filesys.LastWriteTime.Nanoseconds()
	lat = filesys.LastAccessTime.Nanoseconds()
	return
}
