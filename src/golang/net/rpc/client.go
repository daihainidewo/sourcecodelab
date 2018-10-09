// Package main main
// file create by daihao, time is 2018/9/25 12:08
package main

import (
	"fmt"
	"net/rpc"
)

// Opt
type Opt struct {
	A, B int
}

type Server int

// Add Add
func (s *Server) Add(args *Opt, ret *int) (error) {
	*ret = args.A + args.B
	return nil
}

// Sub Sub
func (s *Server) Sub(args *Opt, ret *int) (error) {
	*ret = args.A - args.B
	return nil
}

func main() {
	client, err := rpc.DialHTTP("tcp", "127.0.0.1:8080")
	if err != nil {
		fmt.Println(err)
		return
	}

	a := Opt{5, 1}
	ret := 0
	err = client.Call("Server.Add", a, &ret)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("%d+%d=%d", a.A, a.B, ret)
	err = client.Call("Server.Sub", a, &ret)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("%d-%d=%d", a.A, a.B, ret)

}
