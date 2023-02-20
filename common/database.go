package common

import (
	"context"
	"fmt"
	clientv3 "go.etcd.io/etcd/client/v3"
	"log"
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

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	moduleMap := make(map[string]string)
	moduleMap["redis"] = "0"
	moduleMap["mysql"] = "0"
	moduleMap["minio"] = "0"
	for k, v := range moduleMap {
		_, err = cli.Put(ctx, k, v)
		if err != nil {
			log.Printf("err", err)
		}
	}

	getResponse, _ := cli.Get(ctx, "mysql")
	if getResponse != nil {
		fmt.Println(getResponse)
	}
	for _, kv := range getResponse.Kvs {
		fmt.Printf("%s=%s\n", kv.Key, kv.Value)
	}
	fmt.Println("connect to etcd success")
	CLI = cli
	return CLI
}
