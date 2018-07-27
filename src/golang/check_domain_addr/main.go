// file create by daihao, time is 2018/7/26 10:59
package main

import (
	"flag"
	"net"
	"fmt"
	"net/http"
	"io/ioutil"
)

type data struct {
}

var domain = flag.String("d", "www.baidu.com", "domain")

func main() {
	flag.Parse()
	rec, err := net.LookupIP(*domain)
	if err != nil {
		fmt.Println("parse domain error", err)
		return
	}
	for _, d := range rec {
		fmt.Println("ip", d)
		u := "http://freeapi.ipip.net/" + d.String()
		resp, err := http.Get(u)
		if err != nil {
			fmt.Printf("get ip data error, please try again, error massage %s\n", err)
			continue
		}
		defer resp.Body.Close()
		html, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			fmt.Printf("parse ip error, paease try again, error massage %s\n", err)
			continue
		}
		if len(html) == 0 {
			continue
		}
		shtml := string(html[1:len(html)-2])
		fmt.Println(shtml)
	}
}
