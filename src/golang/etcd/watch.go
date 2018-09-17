// Package main main
// file create by daihao, time is 2018/9/11 14:48
package main

import (
	"context"
	"fmt"
	"github.com/coreos/etcd/clientv3"
	"time"
)

func main() {
	start := time.Now()
	client, err := clientv3.New(clientv3.Config{
		Endpoints:   []string{"localhost:2379", "localhost:2380"},
		DialTimeout: 5 * time.Second,
	})
	if err != nil {
		fmt.Println("new client error")
		return
	}
	defer client.Close()
	fmt.Println("connect success", time.Now().Sub(start).String())

	ctx, cancle := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancle()

	key := "/test/"
	wat := client.Watch(ctx, key)
	go func() {
		for {
			fmt.Println(len(wat))
			time.Sleep(1 * time.Second)
		}
	}()
	for wresp := range wat {
		for _, ev := range wresp.Events {
			fmt.Printf("%s %q : %q\n", ev.Type, ev.Kv.Key, ev.Kv.Value)
		}
	}
}
