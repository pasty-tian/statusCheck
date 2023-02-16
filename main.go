package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/statusCheck/modules"
	clientv3 "go.etcd.io/etcd/client/v3"
	"time"
)

func main() {
	r := gin.Default()
	r.GET("/health/redis", modules.RedisStatusCheck)
	r.GET("/health/mysql", modules.MysqlStatusCheck)
	err := r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
	if err != nil {
		panic(err)
	}
}

func Init() {
	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   []string{"127.0.0.1:2379"},
		DialTimeout: 5 * time.Second,
	})
	if err != nil {
		// handle error!
		fmt.Printf("connect to etcd failed, err:%v\n", err)
		return
	}
	fmt.Println("connect to etcd success")
	defer cli.Close()
}
