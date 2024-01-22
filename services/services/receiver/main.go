package main

import (
	"fmt"

	"github.com/gin-gonic/gin"

	"services/common/rabbit"
)

func Run(host string, port int) {
	gin.SetMode(gin.ReleaseMode)

    router := gin.Default()
	router.ForwardedByClientIP = true
	router.SetTrustedProxies([]string{"127.0.0.1"})

    router.GET("/v1/healthcheck", Healthcheck)
    router.POST("/v1/data/:type", CreateData)

    router.Run(fmt.Sprintf("%s:%d", host, port))
}

func main() {
    if err := rabbit.Connect(); err != nil {
		panic(err)
    }

    defer rabbit.Disconnect()

    Run("0.0.0.0", 7000)
}
