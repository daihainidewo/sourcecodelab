// Package context context
// file create by daihao, time is 2018/9/13 11:57
package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	c := context.WithValue(context.Background(), "abc", 123)
	fmt.Println(c.Value("abc"))
	c, cancel := context.WithTimeout(c, 100*time.Millisecond)
	c = context.WithValue(c, "abcd", 1235)
	fmt.Println(c.Deadline())
	cancel()
	time.Sleep(2 * time.Second)
	fmt.Println(c.Value("abcd"))
	fmt.Println(c.Value("abc"))

	

}
