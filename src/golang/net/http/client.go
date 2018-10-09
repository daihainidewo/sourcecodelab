// Package main main
// file create by daihao, time is 2018/9/25 12:08
package main

import (
	"bytes"
	"fmt"
	"net/http"
	"time"
)

func main() {
	for {
		go post()
		time.Sleep(10 * time.Millisecond)
	}
}

func post() {
	resp, err := http.Post("http://localhost:8080/server", "application/x-www-form-urlencoded", bytes.NewBufferString("id=1234567890"))
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(resp.Status)
	//resp.Body.Close()
}
