// Package main main
// file create by daihao, time is 2018/7/27 18:10
package main

import (
	"encoding/json"
	"fmt"
)

type abs struct {
	A1 string  `json:"a1"`
	A2 int     `json:"a2"`
	A3 float64 `json:"a3"`
}

// String String
func (o *abs) String() string {
	// TODO
	return fmt.Sprintf("{a1:%s, a2:%d, a3:%f}", o.A1, o.A2, o.A3)
}

// main main
func main() {
	// TODO

	a := abs{"abc", 1, 1.0}
	b := new(abs)
	//fmt.Println(a.String())
	c, e := json.Marshal(a)
	if e != nil {
		fmt.Println(e)
		return
	}
	fmt.Println(string(c))
	e = json.Unmarshal(c, b)
	if e != nil {
		fmt.Println(e)
		return
	}
	fmt.Println(b)
	d := abs{}
	e = json.Unmarshal(c, d)
	if e != nil {
		fmt.Println(e)
		return
	}
	fmt.Println(d)
}
