// Package serverfound serverfound
// file create by daihao, time is 2018/9/11 15:27
package serverfound

import (
	"context"
	"github.com/coreos/etcd/clientv3"
	"github.com/coreos/etcd/mvcc/mvccpb"
	"sync"
)

// Server
type Server struct {
	info   ServiceInfo
	Status int
}

// ServerPool
type ServerPool struct {
	services map[string]*Server
	client   *clientv3.Client
	rwsm     *sync.RWMutex
}

// NewServerPool new ServerPool
func NewServerPool(conf clientv3.Config) *ServerPool {
	// TODO
	cli, err := clientv3.New(conf)
	if err != nil {
		panic(err)
	}
	ret := new(ServerPool)
	ret.client = cli
	ret.rwsm = new(sync.RWMutex)
	ret.services = make(map[string]*Server)
	return ret
}

func Default(conf clientv3.Config) {
	// TODO
	sync.Once{}.Do(func() {
		NewServerPool(conf).Watcher()
	})
}

// Watcher Watcher
func (sp *ServerPool) Watcher() () {
	// TODO
	key := ""
	wc := sp.client.Watch(context.Background(), key)
	for msg := range wc {
		for _, ev := range msg.Events {
			switch ev.Type {
			case mvccpb.PUT:
				sp.AddServer(ev.Kv)
			case mvccpb.DELETE:
				sp.RemoveServer(ev.Kv)
			default:

			}
		}
	}
}
func (sp *ServerPool) AddServer(kv *mvccpb.KeyValue) {

}
func (sp *ServerPool) RemoveServer(kv *mvccpb.KeyValue) {

}
