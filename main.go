package main

import (
	"github.com/gin-gonic/gin"
	"github.com/statusCheck/modules"
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
