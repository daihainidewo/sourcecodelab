// Package main main
// file create by daihao, time is 2018/9/25 12:08
package main

import (
	"fmt"
	"net/http"
	"net/rpc"
	"os"
	"os/signal"
	"syscall"
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

	rpcserver := rpc.NewServer()
	err := rpcserver.Register(new(Server))
	if err != nil {
		fmt.Println(err)
		return
	}
	rpc.HandleHTTP()


	signCh := make(chan os.Signal)
	signal.Notify(signCh, os.Interrupt, os.Kill, syscall.SIGTERM)

	go startRouter(8080)

	<-signCh
}

func startRouter(port int) error {
	http.HandleFunc("/server", server)

	err := http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
	if err != nil {
		fmt.Println(err)
	}
	return err
}

// server server
func server(w http.ResponseWriter, r *http.Request) {
	// TODO
	fmt.Println("get a client")
	r.ParseForm()
	id := r.PostFormValue("id")
	w.Write([]byte(id))
}
