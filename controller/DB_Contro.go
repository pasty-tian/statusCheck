package controller

import (
	"fmt"
	"github.com/statusCheck/common"
	"time"
	"context"
)

func PutKey() {
	CLI := common.GetDB()
	ctx, _ = context.WithTimeout(context.Background(), time.Second)
	resp, err := CLI.Get(ctx, "q1mi")
	if err != nil {
		fmt.Printf("get from etcd failed, err:%v\n", err)
		return
	}
	for _, ev := range resp.Kvs {
		fmt.Printf("%s:%s\n", ev.Key, ev.Value)
}
