package main

import (
	"os"
	"time"

	log "github.com/sirupsen/logrus"

	"aggregator/roach"
)

func ProcessCountByYear() {

}

func ProcessCountByModel() {

}

func main() {
	if err := roach.Connect(); err != nil {
		log.Error(err)
		os.Exit(1)
	}

	defer roach.Disconnect()

	interval := 1000

	for {
		ProcessCountByYear()
		ProcessCountByModel()

		log.Debug("Sleeping for", interval, "ms")

		time.Sleep(time.Duration(interval) * time.Millisecond)
	}
}
