// Package serverfound serverfound
// file create by daihao, time is 2018/9/11 15:13
package serverfound
import (
	"context"
	"encoding/json"
	"github.com/coreos/etcd/clientv3"
	"log"
	"time"
)

// Service
type Service struct {
	ProcessID int
	info      ServiceInfo
	client    *clientv3.Client
}

// ServiceInfo
type ServiceInfo struct {
	ID   int
	IP   string
	Port int
}

// HeartBeat HeartBeat
func (s *Service) HeartBeat() () {
	// TODO
	for {
		key := ""
		info, _ := json.Marshal(s.info)
		_, err := s.client.Put(context.Background(), key, string(info))
		if err != nil {
			log.Println(err)
		}
		time.Sleep(10 * time.Second)
	}
}

// RegisterService RegisterService
func RegisterService(conf clientv3.Config) {
	// TODO
	cli, err := clientv3.New(conf)
	if err != nil {
		panic(err)
	}
	sev := new(Service)
	sev.client = cli

	go sev.HeartBeat()
}

