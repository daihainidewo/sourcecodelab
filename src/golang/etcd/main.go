// Package etcd etcd
// file create by daihao, time is 2018/9/10 15:54
package main

import (
	"context"
	"fmt"
	"github.com/coreos/etcd/clientv3"
	"time"
)

func main() {
	start := time.Now()
	fmt.Println("start", start)

	client, err := clientv3.New(clientv3.Config{
		Endpoints:   []string{"localhost:2379", "localhost:2380"},
		DialTimeout: 5 * time.Second,
	})
	if err != nil {
		fmt.Println("connect failed error %s", err)
	}
	defer client.Close()
	fmt.Println("connect success", time.Now().Sub(start).String())

	ctx, cancle := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancle()

	key := "/test/"

	//_, err = client.Put(ctx, key, "test")
	//if err != nil {
	//	fmt.Println(err)
	//	return
	//}
	resp, err := client.Get(ctx, key)
	if err != nil {
		fmt.Println("get failed error %s", err)
		return
	}
	for _, d := range resp.Kvs {
		fmt.Printf("key: %s, value: %s", d.Key, d.Value)
	}

}

// Hello Hello
func Hello()  {
    // TODO

}
