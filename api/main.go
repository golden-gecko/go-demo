package main

import (
	"fmt"
	"os"

	"github.com/gin-gonic/gin"

	log "github.com/sirupsen/logrus"

	"api/api"
	"api/roach"
)

func Run(host string, port int) {
	router := gin.Default()

	router.GET("/v1/healthcheck", api.Healthcheck)

	router.GET("/v1/users", api.GetUsers)
	router.GET("/v1/users/:userId", api.GetUser)

	router.POST("/v1/users", api.CreateUser)
	router.DELETE("/apv1/users/:userId", api.DeleteUser)

	router.GET("/v1/vehicles", api.GetVehicles)
	router.GET("/v1/vehicles/:vehicleId", api.GetVehicle)

	router.POST("/v1/vehicles", api.CreateVehicle)
	router.POST("/v1/vehicles/:plate/transits", api.CreateTransit)

	router.Run(fmt.Sprintf("%s:%d", host, port))
}

func main() {
	if err := roach.Connect(); err != nil {
		log.Error(err)
		os.Exit(1)
	}

	defer roach.Disconnect()

	Run("0.0.0.0", 6000)
}
