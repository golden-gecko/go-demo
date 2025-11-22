package main

import (
	"fmt"

	"github.com/gin-gonic/gin"

	"services/common/sql"
)

func Run(host string, port int) {
    router := gin.Default()

    router.GET("/v1/healthcheck", Healthcheck)
    router.GET("/v1/users", GetUsers)
    router.GET("/v1/users/:userId", GetUser)
    router.POST("/v1/users", CreateUser)
    router.GET("/v1/vehicles", GetVehicles)
    router.POST("/v1/vehicles", CreateVehicle)
    router.POST("/v1/vehicles/:plate/transits", CreateTransit)

    router.Run(fmt.Sprintf("%s:%d", host, port))
}

func main() {
	err := sql.Connect();

    if err != nil {
		panic(err)
    }

    defer sql.Disconnect()

    Run("0.0.0.0", 6000)
}
