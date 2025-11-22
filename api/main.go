package main

import (
	"fmt"
	"os"

	"github.com/gin-gonic/gin"

	"api/api"
	"api/roach"
)

func Run(host string, port int) {
	router := gin.Default()

	router.GET("/api/v1/users", api.GetUsers)
	router.GET("/api/v1/users/:userID", api.GetUser)

	router.POST("/api/v1/users", api.CreateUser)
	router.DELETE("/api/v1/users/:userID", api.DeleteUser)

	router.GET("/api/v1/vehicles", api.GetVehicles)
	router.GET("/api/v1/vehicles/:vehicleID", api.GetVehicle)

	router.POST("/api/v1/vehicles", api.CreateVehicle)
	router.POST("/api/v1/vehicles/:plate/transits", api.CreateTransit)

	router.Run(fmt.Sprintf("%s:%d", host, port))
}

func main() {
	if err := roach.Connect(); err != nil {
		os.Exit(1)
	}

	defer roach.Disconnect()

	Run("0.0.0.0", 5000)
}
