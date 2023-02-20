package controller

import (
	"context"
	"fmt"
	clientv3 "go.etcd.io/etcd/client/v3"
)

func Watcher(client *clientv3.Client, key string) {

	// 监听这个chan
	watchChan := client.Watch(context.Background(), key)

	for watchResponse := range watchChan {

		for _, event := range watchResponse.Events {
			fmt.Printf("Type:%s,Key:%s,Value:%s\n", event.Type, event.Kv.Key, event.Kv.Value)
			/**
			Type:PUT,Key:/ns/service,Value:127.0.0.1:8000
			Type:PUT,Key:/ns/service,Value:127.0.0.1:8001
			Type:PUT,Key:/ns/service,Value:127.0.0.1:8002
			Type:PUT,Key:/ns/service,Value:127.0.0.1:8003
			...
			*/
		}
	}
}
