package main

import (
	"fmt"
	"os"

	"github.com/gin-gonic/gin"

	log "github.com/sirupsen/logrus"

	"receiver/api"
	"receiver/queue"
)

func Run(host string, port int) {
	router := gin.Default()

	router.POST("/api/v1/data/:type", api.CreateData)

	router.Run(fmt.Sprintf("%s:%d", host, port))
}

func main() {
	err := queue.InitRabbit()

	if err != nil {
		log.Error(err)
		os.Exit(1)
	}

	defer queue.DeinitRabbit()

	Run("0.0.0.0", 7000)
}
