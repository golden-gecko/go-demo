package main

import (
	"fmt"
	"os"

	"github.com/gin-gonic/gin"

	"api/api"
	"api/config"
	"api/queue"
)

func Run(host string, port int) {
	router := gin.Default()

	router.POST("/api/v1/data", api.CreateData)

	router.GET("/api/v1/users", api.GetUsers)
	router.GET("/api/v1/users/:userID", api.GetUser)

	router.GET("/api/v1/vehicles", api.GetVehicles)
	router.GET("/api/v1/vehicles/:vehicleID", api.GetVehicle)

	router.POST("/api/v1/vehicles", api.CreateVehicle)
	router.POST("/api/v1/vehicles/:plate/transits", api.CreateTransit)

	router.Run(fmt.Sprintf("%s:%d", host, port))
}

func main() {
	// db.Init(config.MONGO_DB_URI)
	// defer db.Deinit()

	// queue.InitKafka()
	// defer queue.DeinitKafka()

	err := queue.InitRabbit()

	if err != nil {
		os.Exit(1)
	}

	defer queue.DeinitRabbit()

	Run(config.API_HOST, config.API_PORT)
}
