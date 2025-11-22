package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"math/rand"
	"net/http"
	"time"

	log "github.com/sirupsen/logrus"

	"services/common/model"
)

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

func CreateItem(names model.Names) model.Item {
	i := model.Item{
		Name: names.Names[rand.Intn(len(names.Names))],
	}

	i.Attributes = make(map[string]int)

	if rand.Intn(2) == 0 {
		i.Attributes["Attack"] = rand.Intn(10)
	}

	if rand.Intn(2) == 0 {
		i.Attributes["Defence"] = rand.Intn(10)
	}

	if rand.Intn(2) == 0 {
		i.Attributes["Magic"] = rand.Intn(10)
	}

	return i
}

func CreateTemperature(location string) model.Temperature {
    return model.Temperature{
        Location:  location,
        Value:     rand.Float32() * 100,
        Timestamp: time.Now().Format(time.RFC3339),
    }
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

func CreateUser(users model.Users) model.User {
    return model.User{
        Name:     users.Users[rand.Intn(len(users.Users))],
        Password: RandomWord(20),
    }
}

func CreateVehicle(cars model.Cars) model.Vehicle {
    car := cars.Cars[rand.Intn(len(cars.Cars))]

    return model.Vehicle{
        Plate:    RandomPlate(),
        Brand:    car.Brand,
        Model:    car.Model,
        Year:     RandomYear(),
        Category: car.Category,
    }
}

func Sleep(minInterval int, maxInterval int) {
	interval := minInterval + rand.Intn(maxInterval - minInterval)

	log.Info(fmt.Sprintf("Sleeping for %d ms", interval))

	time.Sleep(time.Duration(interval) * time.Millisecond)
}

func ProcessItems(receiverUrl string, names model.Names, minInterval int, maxInterval int) {
	for {
		var items []model.Item

		for i := 0; i < 1 + rand.Intn(9); i++ {
			items = append(items, CreateItem(names))
		}

		data, err := json.Marshal(items)

		if err != nil {
			panic(err)
		}

		if err := Send(receiverUrl + "/item", data); err != nil {
			panic(err)
		}

		Sleep(minInterval, maxInterval)
	}
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

func ProcessUsers(apiUrl string, users model.Users, minInterval int, maxInterval int) {
	for {
		data, err := json.Marshal(CreateUser(users))

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

    users, err := GetUsers()

	if err != nil {
		panic(err)
	}

	var forever chan struct {}

	go ProcessItems(receiverUrl, names, 1000, 2000)
	go ProcessTemperatures(receiverUrl, 1000, 2000)
	go ProcessTransits(receiverUrl, 2000, 4000)
	go ProcessUsers(apiUrl, users, 4000, 6000)
	go ProcessVehicles(apiUrl, cars, 6000, 8000)

	<- forever
}
