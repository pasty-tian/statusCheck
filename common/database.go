package common

import (
	"fmt"
	clientv3 "go.etcd.io/etcd/client/v3"
	"time"
)

var CLI *clientv3.Client

func Init() *clientv3.Client {
	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   []string{"127.0.0.1:2379"},
		DialTimeout: 5 * time.Second,
	})
	if err != nil {
		// handle error!
		fmt.Printf("connect to etcd failed, err:%v\n", err)
		//panic("dfas", err)
	}
	
	fmt.Println("connect to etcd success")
	CLI = cli
	return CLI
}

func GetDB() *clientv3.Client {
	return CLI
}
