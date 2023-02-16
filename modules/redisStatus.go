package modules

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis"
)

func RedisStatusCheck(c *gin.Context) {
	client := redis.NewClient(&redis.Options{
		Addr:     "redis-15188.c91.us-east-1-3.ec2.cloud.redislabs.com:15188",
		Password: "gOBv7mC9KVEQW7z04zARtZK5hd1j0Poe",
		DB:       0,
	})
	defer client.Close()
	sts, err := client.Get("health").Result()
	if err != nil {
		fmt.Println("redis connection failed：", err)
		c.JSON(500, ("redis status failed"))
		return
	}
	fmt.Println("redis is health: ", sts)
	c.JSON(200, ("redis status ok"))

}
