// Package doc_tree doc_tree
// file create by daihao, time is 2018/8/27 14:45
package main

import (
	"path/filepath"
	"os"
	"sort"
	"time"
	"fmt"
	"sync"
	"strings"
)

// main main
func main() {
	start := time.Now()
	defer func() {
		end := time.Now()
		fmt.Printf("Operating time %gs\n", end.Sub(start).Seconds())
	}()
	p, _ := filepath.Abs("../../game")
	sli := NewFileNodeSlice(p)
	sli.ig.dir[".git"] = struct{}{}
	sli.ig.dir[".idea"] = struct{}{}
	sli.Walk(p)

	last := []string{`\src\golang\sort\sort.go`, `\src\golang\sort\zort.go`, `\src\golang\strings\strings.go`, `\src\golang\test\test.go`, `a/b/c/d`}
	sort.Strings(last)

	c, m, d := sli.Check(last, 0)
	for _, t := range c {
		fmt.Println("create", t)
	}
	for _, t := range m {
		fmt.Println("modify", t)
	}
	for _, t := range d {
		fmt.Println("delete", t)
	}

	fmt.Println("file num:", len(sli.Slice))
}

// Ignore
type Ignore struct {
	dir  map[string]struct{} // 结尾不带分割符
	file map[string]struct{}
}

// FileNodeSlice
type FileNodeSlice struct {
	Slice    []string // 相对于rootpath的相对路径
	sm       *sync.Mutex
	ig       *Ignore
	rootPath string
}

// NewFileNodeSlice new FileNodeSlice
func NewFileNodeSlice(rootpath string) *FileNodeSlice {
	return &FileNodeSlice{
		Slice: make([]string, 0, 10),
		sm:    new(sync.Mutex),
		ig: &Ignore{
			dir:  make(map[string]struct{}),
			file: make(map[string]struct{}),
		},
		rootPath: rootpath,
	}
}

// String String
func (fns *FileNodeSlice) String() (string) {
	ret := "paths:\n"
	for _, d := range fns.Slice {
		ret += d + "\n"
	}
	return ret
}

// Len Len
func (fns *FileNodeSlice) Len() (int) {
	return len(fns.Slice)
}

// Less Less
func (fns *FileNodeSlice) Less(i, j int) (bool) {
	return fns.Slice[i] < fns.Slice[j]
}

// Swap Swap
func (fns *FileNodeSlice) Swap(i, j int) () {
	fns.Slice[i], fns.Slice[j] = fns.Slice[j], fns.Slice[i]
}

// Check 查出新建文件，修改文件，删除文件, last必须是sort后的相对路径序列，timestamp是上次修改的时间戳，返回路径为绝对路径
func (fns *FileNodeSlice) Check(last []string, timestamp int64) (crt []string, mod []string, del []string) {
	// TODO 
	a, b := 0, 0
	crt = make([]string, 0)
	mod = make([]string, 0)
	del = make([]string, 0)
	for a < len(fns.Slice) && b < len(last) {
		if fns.Slice[a] == last[b] {
			ap := filepath.Join(fns.rootPath, last[b])
			info, _ := os.Lstat(ap)
			if info.ModTime().Unix() > timestamp {
				mod = append(mod, ap)
			}
			a++
			b++
		} else if fns.Slice[a] > last[b] {
			ap := filepath.Join(fns.rootPath, last[b])
			del = append(del, ap)
			b++
		} else {
			ap := filepath.Join(fns.rootPath, fns.Slice[a])
			crt = append(crt, ap)
			a++
		}
	}

	if a == len(fns.Slice) {
		del = append(del, last[b:]...)
	} else {
		crt = append(crt, fns.Slice[a:]...)
	}

	return
}

// Sort Sort
func (fns *FileNodeSlice) Sort() () {
	sort.Sort(fns)
}

// Add Add
func (fns *FileNodeSlice) Add(n string) () {
	fns.sm.Lock()
	defer fns.sm.Unlock()
	fns.Slice = append(fns.Slice, n)
}

// Adds Adds
func (fns *FileNodeSlice) Adds(n ...string) () {
	fns.sm.Lock()
	defer fns.sm.Unlock()
	fns.Slice = append(fns.Slice, n...)
}

// ToStringArray ToStringArray
func (fns *FileNodeSlice) ToStringArray() ([]string) {
	return fns.Slice
}

// Walk Walk
func (fns *FileNodeSlice) Walk(path string) ([]string) {
	wg := new(sync.WaitGroup)
	token := NewTokenBucket(1000)
	token.Get()
	wg.Add(1)
	go fns.w(path, wg, &token)
	wg.Wait()
	fns.Sort()
	return fns.ToStringArray()
}

// walk walk
func (fns *FileNodeSlice) w(path string, wg *sync.WaitGroup, token *TokenBucket) error {
	defer token.Put()
	defer wg.Done()
	f, err := fns.readDirNames(path)
	if err != nil {
		fmt.Println(err)
		return err
	}
	tmp := make([]string, 0, len(f))

	for _, d := range f {
		filename := filepath.Join(path, d)
		info, err := os.Lstat(filename)
		if err != nil {
			continue
		}
		if info.IsDir() {
			_, ok := fns.ig.dir[d]
			if ok {
				continue
			}
			wg.Add(1)
			token.Get()
			go func(path string) {
				fns.w(path, wg, token)
			}(filename)
		} else {
			_, ok := fns.ig.file[d]
			if ok {
				continue
			}
			tmps := strings.TrimPrefix(filename, fns.rootPath)
			tmp = append(tmp, tmps)
		}
	}
	fns.Adds(tmp...)
	return nil
}

func (fns *FileNodeSlice) readDirNames(dirname string) ([]string, error) {
	f, err := os.Open(dirname)
	if err != nil {
		return nil, err
	}
	names, err := f.Readdirnames(-1)
	f.Close()
	if err != nil {
		return nil, err
	}
	return names, nil
}

type TokenBucket chan struct{}

func NewTokenBucket(size int) TokenBucket {
	tb := make(chan struct{}, size)
	for i := 0; i < size; i++ {
		tb <- struct{}{}
	}

	return TokenBucket(tb)
}

func (tb TokenBucket) Get() {
	<-tb
}

func (tb TokenBucket) Put() {
	tb <- struct{}{}
}
