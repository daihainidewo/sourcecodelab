// Package test test
// file create by daihao, time is 2018/8/20 17:21
package main

import (
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
	A, B int
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

	//a := map[int]*stt{1: {b: 2}}
	//a[2] = &stt{2, 3}
	//fmt.Println(a)
	//fmt.Printf("%#v\n", a)
	//j, err := json.Marshal(&a)
	//if err != nil {
	//	fmt.Println(err)
	//	return
	//}
	//fmt.Println(string(j))
	//b := map[int]*stt{}
	//json.Unmarshal(j, b)
	//fmt.Println(b)

	//a := 1
	//b := "12"
	//Add(&a)
	//Add(b)
	//fmt.Println(a, b)
	//c := stt{1, 2}
	//d, _ := json.Marshal(c)
	//fmt.Println(string(d))
	//e, _ := json.Marshal(&c)
	//fmt.Println(string(e))

	//echo "{\"name\":\"pcgameq_panda_gift_donate\",\"data\":\"{\"__plat\":\"ios\",\"uid\":\"137914784\",\"anchor\":\"108930328\",\"roomid\":\"3497620\",\"giftid\":\"5aba0ed6ea3d187b391b2293\",\"price\":\"50\",\"count\":\"1\",\"total\":\"50\",\"ip\":\"10.131.7.41\",\"time\":1538136530.2132,\"channel\":\"\",\"unique\":\"5bae19d2dcf63a25092a8046_1\",\"pdft\":\"\",\"cate\":\"gbady\",\"extra_anchor\":\"\",\"fb\":\"\",\"lotteryChance\":\"\"}",\"host\":\"pt5v.plat.bjtb.pdtv.it\",\"key\":\"\",\"time\":\"2018-09-28 20:08:50\",\"requestid\":\"1538136530017-14713101-30284-4d4dc38519ab7a2f\"}"


}

// Add Add
func Add(a interface{}) {
	// TODO
	switch a.(type) {
	case *int:
		b := *a.(*int) + 1
		a = &b
	case *string:
		a = *a.(*string) + "123"
	default:
		fmt.Println("other type")
	}
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
