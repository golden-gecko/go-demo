package common

import (
	"fmt"
	"math/rand"
	"time"

	log "github.com/sirupsen/logrus"
)

func RandomPlate() string {
	const characters = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	const digits = "0123456789"

	return RandomString(4, characters) + RandomString(8, digits)
}

func RandomString(length int, charset string) string {
	b := make([]byte, length)

	for i := range b {
		b[i] = charset[rand.Intn(len(charset))]
	}

	return string(b)
}

func RandomWord(length int) string {
	const characters = "abcdefghijklmnopqrstuvwxyz"

	return RandomString(length, characters)
}

func RandomYear() int {
	return 1980 + rand.Intn(30)
}

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
