// Package test test
// file create by daihao, time is 2018/8/20 17:21
package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"sort"
)

//
//import (
//	"encoding/json"
//	"fmt"
//)
//
//// FileTree
//type FileTree struct {
//	Dir  []*FileTree `json:"dir"`
//	File []string    `json:"file"`
//	Name string      `json:"name"`
//}
//
//// NewFileTree new FileTree
//func NewFileTree(name string) *FileTree {
//	return &FileTree{
//		Dir:  make([]*FileTree, 0),
//		File: make([]string, 0),
//		Name: name,
//	}
//}
//
//// String String
////func (ft *FileTree) String() (string) {
////	// TODO
////
////	path := ft.name
////	//for _, d := range ft.dir {
////	//	path = path + "/" + d.String()
////	//}
////	//for _, d := range ft.file {
////	//	path = d
////	//}
////	return path
////}
//
//func main() {
//
//	a := NewFileTree("a")
//	b := NewFileTree("b")
//	e := NewFileTree("e")
//	a.Dir = append(a.Dir, b)
//	a.Dir = append(a.Dir, e)
//	a.File = append(a.File, "c")
//	b.File = append(b.File, "d")
//	j, err := json.Marshal(a)
//	if err != nil {
//		fmt.Println("json marshal", err)
//		return
//	}
//	fmt.Println(a.Name)
//	fmt.Println(string(j))
//	//fmt.Println(fmt.Println(string(j)))
//	//abspath, err := filepath.Abs("../")
//	//if err != nil {
//	//	fmt.Println(err)
//	//	return
//	//}
//	//fi, err := os.Stat(abspath)
//	//if err != nil {
//	//	fmt.Println(err)
//	//	return
//	//}
//	//
//	//ft := NewFileTree(fi.Name())
//	//if fi.IsDir() {
//	//	tmp := NewFileTree(fi.Name())
//	//	ft.dir = append(ft.dir, tmp)
//	//} else {
//	//	ft.file = append(ft.file, ft.name)
//	//}
//
//}

//func main() {
//	rootpath, _ := filepath.Abs(".")
//
//	root := FileNode{"projects", rootpath, []*FileNode{}}
//	fileInfo, _ := os.Lstat(rootpath)
//
//	walk(rootpath, fileInfo, &root)
//
//	data, _ := json.Marshal(root)
//
//	fmt.Printf("%s", data)
//}
// FileNode
type FileNode struct {
	Name     string      `json:"name"`
	Path     string      `json:"path"`
	Children []*FileNode `json:"children"`
}

// walk walk
func walk(path string, node *FileNode) {
	// TODO
	files, _ := ioutil.ReadDir(path)
	for _, file := range files {
		fpath := filepath.Join(path, file.Name())

		child := FileNode{file.Name(), fpath, []*FileNode{}}
		node.Children = append(node.Children, &child)

		if file.IsDir() {
			walk(fpath, &child)
		}
	}
}

// In
type In interface {
	In()
}

// Out
type Out struct {
	In
	a int
}

// Ini
type Ini struct {
}

// In In
func (i *Ini) In() () {
	// TODO
	fmt.Println("hello")
}

// stt
type stt struct {
	a, b int
}

// CheckDirUpdate CheckDirUpdate
func main() { //([]string, error) {
	//
	//a := new(Out)
	//	//a.In = new(Ini)
	//	//a.In.In()
	//	//go func() {
	//	//	for {
	//	//		fmt.Println("1")
	//	//		time.Sleep(1 * time.Second)
	//	//	}
	//	//}()
	//	//time.Sleep(1 * time.Hour)

	a := map[int]*stt{1: {b: 2}}
	a[2] = &stt{2, 3}
	fmt.Println(a)
	fmt.Printf("%#v\n", a)
	j, err := json.Marshal(&a)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(j))
	b := map[int]*stt{}
	json.Unmarshal(j, b)
	fmt.Println(b)
}

//type FileNode struct {
//	Name      string      `json:"name"`
//	Path      string      `json:"path"`
//	FileNodes []*FileNode `json:"children"`
//}

//func walk(path string, info os.FileInfo, node *FileNode) {
//	// 列出当前目录下的所有目录、文件
//	files := listFiles(path)
//
//	// 遍历这些文件
//	for _, filename := range files {
//		// 拼接全路径
//		fpath := filepath.Join(path, filename)
//
//		// 构造文件结构
//		fio, _ := os.Lstat(fpath)
//
//		// 将当前文件作为子节点添加到目录下
//		child := FileNode{filename, fpath, []*FileNode{}}
//		node.FileNodes = append(node.FileNodes, &child)
//
//		// 如果遍历的当前文件是个目录，则进入该目录进行递归
//		if fio.IsDir() {
//			walk(fpath, fio, &child)
//		}
//	}
//
//	return
//}

func listFiles(dirname string) []string {
	f, _ := os.Open(dirname)

	names, _ := f.Readdirnames(-1)
	f.Close()

	sort.Strings(names)

	return names
}
