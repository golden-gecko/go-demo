package main

import (
	"bytes"
	"encoding/json"
	"log"
	"math/rand"
	"net/http"
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

func Send(body *bytes.Buffer) {
	resp, err := http.Post("http://192.168.10.26:5000/api/v1/data", "application/json", body)

	if err != nil {
		panic(err)
	}

	if resp.StatusCode != 200 {
		log.Println(resp.StatusCode)
	}
}

func CreateTemperature() {
	t := Temperature{
		Location:  "Room",
		Value:     rand.Float32() * 100,
		Timestamp: time.Now().Format(time.RFC3339),
	}

	temperature, err := json.Marshal(t)

	if err != nil {
		panic(err)
	}

	body := bytes.NewBuffer(temperature)

	Send(body)
}

func CreateTransit() {
	c := Coordinate{
		Latitude:  14.0745211117 + rand.Float32()*(24.0299857927-14.0745211117),
		Longitude: 49.0273953314 + rand.Float32()*(54.8515359564-49.0273953314),
	}

	t := Transit{
		Plate:     RandomPlate(),
		Location:  c,
		Timestamp: time.Now().Format(time.RFC3339),
	}

	transit, err := json.Marshal(t)

	if err != nil {
		panic(err)
	}

	body := bytes.NewBuffer(transit)

	Send(body)
}

func main() {
	for {
		if rand.Intn(2) == 0 {
			CreateTemperature()
		}

		if rand.Intn(2) == 0 {
			CreateTransit()
		}

		time.Sleep(time.Duration(rand.Intn(2)) * time.Second)
	}
}
