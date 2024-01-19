package main

import (
	"fmt"

	"github.com/gin-gonic/gin"

	"services/common/queue"
)

func Run(host string, port int) {
    router := gin.Default()

    router.GET("/v1/healthcheck", Healthcheck)
    router.POST("/v1/data/:type", CreateData)

    router.Run(fmt.Sprintf("%s:%d", host, port))
}

func main() {
    if err := queue.Connect(); err != nil {
		panic(err)
    }

    defer queue.Disconnect()

    Run("0.0.0.0", 7000)
}
