package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"math/rand"
	"net/http"
	"os"
	"path/filepath"
	"time"

	log "github.com/sirupsen/logrus"

	"services/common/model"
)

func GetCars() (model.Cars, error) {
    var cars model.Cars

    jsonFile, err := os.Open(filepath.Join("services", "producer", "data", "cars.json"))

    if err != nil {
		return cars, err
    }

    defer jsonFile.Close()

    byteValue, err := io.ReadAll(jsonFile)

	if err != nil {
		return cars, err
    }

    if json.Unmarshal(byteValue, &cars) != nil {
		return cars, err
	}

    return cars, nil
}

func GetNames() (model.Names, error) {
    var names model.Names

    jsonFile, err := os.Open(filepath.Join("services", "producer", "data", "names.json"))

    if err != nil {
		return names, err
    }

    defer jsonFile.Close()

    byteValue, err := io.ReadAll(jsonFile)

	if err != nil {
		return names, err
    }

    if json.Unmarshal(byteValue, &names) != nil {
		return names, err
	}

    return names, nil
}

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

func Send(url string, data []byte) error {
	log.Info(url)

	response, err := http.Post(url, "application/json", bytes.NewBuffer(data))

    if err != nil {
        return err
    }

	body, err := io.ReadAll(response.Body)

    if err != nil {
		log.Info(fmt.Sprintf("%d %s", response.StatusCode, string(body)))
    } else {
		log.Info(response.StatusCode)
	}

    return nil
}

func CreateTemperature(location string) model.Temperature {
    t := model.Temperature{
        Location:  location,
        Value:     rand.Float32() * 100,
        Timestamp: time.Now().Format(time.RFC3339),
    }

    return t
}

func CreateTransit() model.Transit {
    c := model.Coordinate{
        Latitude:  14.0745211117 + rand.Float32() * (24.0299857927 - 14.0745211117),
        Longitude: 49.0273953314 + rand.Float32() * (54.8515359564 - 49.0273953314),
    }

    t := model.Transit{
        Plate:     RandomPlate(),
        Location:  c,
        Timestamp: time.Now().Format(time.RFC3339),
    }

    return t
}

func CreateUser(names model.Names) model.User {
    u := model.User{
        Name:     names.Names[rand.Intn(len(names.Names))],
        Password: RandomWord(20),
    }

    return u
}

func CreateVehicle(cars model.Cars) model.Vehicle {
    car := cars.Cars[rand.Intn(len(cars.Cars))]

    t := model.Vehicle{
        Plate:    RandomPlate(),
        Brand:    car.Brand,
        Model:    car.Model,
        Year:     RandomYear(),
        Category: car.Category,
    }

    return t
}

func Sleep(minInterval int, maxInterval int) {
	interval := minInterval + rand.Intn(maxInterval - minInterval)

	log.Info(fmt.Sprintf("Sleeping for %d ms", interval))

	time.Sleep(time.Duration(interval) * time.Millisecond)
}

func ProcessTemperatures(receiverUrl string, minInterval int, maxInterval int) {
	for {
		t1 := CreateTemperature("Room #1")
		t2 := CreateTemperature("Room #2")
		t3 := CreateTemperature("Room #3")

		data, err := json.Marshal([]model.Temperature{t1, t2, t3})

		if err != nil {
			panic(err)
		}

		if err := Send(receiverUrl + "/temperature", data); err != nil {
			panic(err)
		}

		Sleep(minInterval, maxInterval)
	}
}

func ProcessTransits(receiverUrl string, minInterval int, maxInterval int) {
	for {
		var transits []model.Transit

		for i := 0; i < 1 + rand.Intn(9); i++ {
			transits = append(transits, CreateTransit())
		}

		data, err := json.Marshal(transits)

		if err != nil {
			panic(err)
		}

		if err := Send(receiverUrl + "/transit", data); err != nil {
			panic(err)
		}

		Sleep(minInterval, maxInterval)
	}
}

func ProcessUsers(apiUrl string, names model.Names, minInterval int, maxInterval int) {
	for {
		data, err := json.Marshal(CreateUser(names))

		if err != nil {
			panic(err)
		}

		if err := Send(apiUrl + "/users", data); err != nil {
			panic(err)
		}

		Sleep(minInterval, maxInterval)
	}
}

func ProcessVehicles(apiUrl string, cars model.Cars, minInterval int, maxInterval int) {
	for {
		data, err := json.Marshal(CreateVehicle(cars))

		if err != nil {
			panic(err)
		}

		if err := Send(apiUrl + "/vehicles", data); err != nil {
			panic(err)
		}

		Sleep(minInterval, maxInterval)
	}
}

func main() {
    apiUrl := "http://haproxy:6100/v1"
    receiverUrl := "http://haproxy:7100/v1/data"

    cars, err := GetCars()

	if err != nil {
		panic(err)
	}

    names, err := GetNames()

	if err != nil {
		panic(err)
	}

	var forever chan struct {}

	go ProcessTemperatures(receiverUrl, 100, 200)
	go ProcessTransits(receiverUrl, 2000, 4000)
	go ProcessUsers(apiUrl, names, 4000, 6000)
	go ProcessVehicles(apiUrl, cars, 6000, 8000)

	<- forever
}
