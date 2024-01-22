package common

import (
	"fmt"
	"math/rand"
	"time"

	log "github.com/sirupsen/logrus"
)

func Sleep(interval int) {
	log.Info(fmt.Sprintf("Sleeping for %d ms", interval))

	time.Sleep(time.Duration(interval) * time.Millisecond)
}

func SleepRange(minInterval int, maxInterval int) {
	interval := minInterval + rand.Intn(maxInterval - minInterval)

	log.Info(fmt.Sprintf("Sleeping for %d ms", interval))

	time.Sleep(time.Duration(interval) * time.Millisecond)
}

func Timestamp() string {
	return time.Now().UTC().Format("2006-01-02 15:04:05.000000")
}

func StringToTimestamp(timestamp string) (time.Time, error) {
	return time.Parse("2006-01-02 15:04:05.000000", timestamp)
}
