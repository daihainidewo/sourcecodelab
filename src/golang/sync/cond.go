// Package main main
// file create by daihao, time is 2018/10/9 10:55
package main

import (
	"log"
	"sync"
)

func main() {
	var m sync.Mutex
	c := sync.NewCond(&m)
	n := 200
	running := make(chan bool, n)
	awake := make(chan bool, n)
	for i := 0; i < n; i++ {
		go func() {
			m.Lock()
			running <- true
			c.Wait()
			awake <- true
			m.Unlock()
		}()
	}
	for i := 0; i < n; i++ {
		<-running // Wait for everyone to run.
	}
	for n > 0 {
		select {
		case <-awake:
			log.Fatal("goroutine not asleep")
		default:
		}
		m.Lock()
		c.Signal()
		m.Unlock()
		<-awake // Will deadlock if no goroutine wakes up
		select {
		case <-awake:
			log.Fatal("too many goroutines awake")
		default:
		}
		n--
	}
	c.Signal()
}
