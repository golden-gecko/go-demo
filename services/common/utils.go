package common

import (
	"fmt"
	"time"

	log "github.com/sirupsen/logrus"
)

func Sleep(interval int) {
	log.Info(fmt.Sprintf("Sleeping for %d ms", interval))

	time.Sleep(time.Duration(interval) * time.Millisecond)
}
