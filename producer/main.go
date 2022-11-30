package main

import (
	"bytes"
	"encoding/json"
	"log"
	"math/rand"
	"net/http"
	"os"
	"time"
)

type Coordinate struct {
	Latitude  float32 `bson:"Latitude"`
	Longitude float32 `bson:"Longitude"`
}

type Temperature struct {
	Location  string  `bson:"Location"`
	Value     float32 `bson:"Value"`
	Timestamp string  `bson:"Timestamp"`
}

type Transit struct {
	Plate     string     `bson:"Plate"`
	Location  Coordinate `bson:"Location"`
	Timestamp string     `bson:"Timestamp"`
}

type User struct {
	Name     string `bson:"Name"`
	Password string `bson:"Password"`
}

func RandomString(length int, charset string) string {
	b := make([]byte, length)

	for i := range b {
		b[i] = charset[rand.Intn(len(charset))]
	}

	return string(b)
}

func RandomPlate() string {
	const characters = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	const digits = "0123456789"

	return RandomString(3, characters) + RandomString(6, digits)
}

func RandomWord(length int) string {
	const characters = "abcdefghijklmnopqrstuvwxyz"

	return RandomString(length, characters)
}

func Send(url string, data []byte) error {
	body := bytes.NewBuffer(data)
	resp, err := http.Post(url, "application/json", body)

	if err != nil {
		log.Println(err)
		return err
	}

	if resp.StatusCode != 200 && resp.StatusCode != 201 && resp.StatusCode != 202 {
		log.Println(resp.StatusCode)
	}

	return nil
}

func CreateUser() User {
	t := User{
		Name:     RandomWord(10),
		Password: RandomWord(20),
	}

	return t
}

func CreateTemperature(location string) Temperature {
	t := Temperature{
		Location:  location,
		Value:     rand.Float32() * 100,
		Timestamp: time.Now().Format(time.RFC3339),
	}

	return t
}

func CreateTransit() Transit {
	c := Coordinate{
		Latitude:  14.0745211117 + rand.Float32()*(24.0299857927-14.0745211117),
		Longitude: 49.0273953314 + rand.Float32()*(54.8515359564-49.0273953314),
	}

	t := Transit{
		Plate:     RandomPlate(),
		Location:  c,
		Timestamp: time.Now().Format(time.RFC3339),
	}

	return t
}

func main() {
	apiUrl := "http://192.168.0.213:9010/api/v1"
	recceiverUrl := "http://haproxy:9000/api/v1/data"

	minInterval := 100
	maxInterval := 500

	for {
		if true {
			u1 := CreateUser()

			data, err := json.Marshal(u1)

			if err != nil {
				os.Exit(1)
			}

			Send(apiUrl+"/users", data)
		}

		if true {
			t1 := CreateTemperature("Room #1")
			t2 := CreateTemperature("Room #2")
			t3 := CreateTemperature("Room #3")

			data, err := json.Marshal([]Temperature{t1, t2, t3})

			if err != nil {
				os.Exit(1)
			}

			Send(recceiverUrl+"/temperature", data)
		}

		if true {
			var transits []Transit

			for i := 0; i < 1+rand.Intn(9); i++ {
				transits = append(transits, CreateTransit())
			}

			data, err := json.Marshal(transits)

			if err != nil {
				os.Exit(1)
			}

			Send(recceiverUrl+"/transit", data)
		}

		sleep := minInterval + rand.Intn(maxInterval-minInterval)

		log.Println("Sleeping for", sleep, "ms")

		time.Sleep(time.Duration(sleep) * time.Millisecond)
	}
}
