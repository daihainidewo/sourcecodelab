// Package main main
// file create by daihao, time is 2018/9/25 12:08
package main

import (
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

func main() {

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
