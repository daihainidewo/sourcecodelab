// Package flags flags
// file create by daihao, time is 2018/8/30 15:13
package main

import "os"

// ArgUnit
type ArgUnit struct {
	Name    string
	Value   interface{}
	Usage   string
	Default interface{}
}

// Args
type Args struct {
	args map[string]map[string]interface{}
}

// Parse Parse
func (a *Args) Parse() () {
	// TODO
	opt := os.Args[1]
	args := make([]string, 0, len(os.Args)-2)
	if len(os.Args) > 2 {
		args = os.Args[2:]
	}
}

func main() {

}
