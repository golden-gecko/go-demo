package main

import (
	"bytes"
	"encoding/json"
	"io"
	"math/rand"
	"net/http"
	"os"
	"path/filepath"
	"sync"
	"time"

	log "github.com/sirupsen/logrus"

	"producer/models"
)

func GetCars() (models.Cars, error) {
    var cars models.Cars

    jsonFile, err := os.Open(filepath.Join("data", "cars.json"))

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

func GetNames() (models.Names, error) {
    var names models.Names

    jsonFile, err := os.Open(filepath.Join("data", "names.json"))

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
	log.Println(url)

	response, err := http.Post(url, "application/json", bytes.NewBuffer(data))

    if err != nil {
        return err
    }

	body, err := io.ReadAll(response.Body)

    if err != nil {
		log.Println(response.StatusCode)
    } else {
		log.Println(response.StatusCode, string(body))
	}

    return nil
}

func CreateTemperature(location string) models.Temperature {
    t := models.Temperature{
        Location:  location,
        Value:     rand.Float32() * 100,
        Timestamp: time.Now().Format(time.RFC3339),
    }

    return t
}

func CreateTransit() models.Transit {
    c := models.Coordinate{
        Latitude:  14.0745211117 + rand.Float32() * (24.0299857927 - 14.0745211117),
        Longitude: 49.0273953314 + rand.Float32() * (54.8515359564 - 49.0273953314),
    }

    t := models.Transit{
        Plate:     RandomPlate(),
        Location:  c,
        Timestamp: time.Now().Format(time.RFC3339),
    }

    return t
}

func CreateUser(names models.Names) models.User {
    u := models.User{
        Name:     names.Names[rand.Intn(len(names.Names))],
        Password: RandomWord(20),
    }

    return u
}

func CreateVehicle(cars models.Cars) models.Vehicle {
    car := cars.Cars[rand.Intn(len(cars.Cars))]

    t := models.Vehicle{
        Plate:    RandomPlate(),
        Brand:    car.Brand,
        Model:    car.Model,
        Year:     RandomYear(),
        Category: car.Category,
    }

    return t
}

func ProcessTemperatures(wg *sync.WaitGroup, receiverUrl string) error {
	for {
		t1 := CreateTemperature("Room #1")
		t2 := CreateTemperature("Room #2")
		t3 := CreateTemperature("Room #3")

		data, err := json.Marshal([]models.Temperature{t1, t2, t3})

		if err != nil {
			return err
		}

		if err := Send(receiverUrl + "/temperature", data); err != nil {
			log.Error(err)
			os.Exit(1)
		}

		minInterval := 1000
		maxInterval := 2000

        interval := minInterval + rand.Intn(maxInterval - minInterval)

        log.Info("Sleeping for ", interval, " ms")

        time.Sleep(time.Duration(interval) * time.Millisecond)
	}
}

func ProcessTransits(wg *sync.WaitGroup, receiverUrl string) error {
	for {
		var transits []models.Transit

		for i := 0; i < 1 + rand.Intn(9); i++ {
			transits = append(transits, CreateTransit())
		}

		data, err := json.Marshal(transits)

		if err != nil {
			return err
		}

		if err := Send(receiverUrl + "/transit", data); err != nil {
			log.Error(err)
			os.Exit(1)
		}

		minInterval := 1000
		maxInterval := 2000

        interval := minInterval + rand.Intn(maxInterval - minInterval)

        log.Info("Sleeping for ", interval, " ms")

        time.Sleep(time.Duration(interval) * time.Millisecond)
	}
}

func ProcessUsers(wg *sync.WaitGroup, apiUrl string, names models.Names) error {
	for {
		data, err := json.Marshal(CreateUser(names))

		if err != nil {
			return err
		}

		if err := Send(apiUrl + "/users", data); err != nil {
			log.Error(err)
			os.Exit(1)
		}

		minInterval := 1000
		maxInterval := 2000

        interval := minInterval + rand.Intn(maxInterval - minInterval)

        log.Info("Sleeping for ", interval, " ms")

        time.Sleep(time.Duration(interval) * time.Millisecond)
	}
}

func ProcessVehicles(wg *sync.WaitGroup, apiUrl string, cars models.Cars) error {
	for {
		data, err := json.Marshal(CreateVehicle(cars))

		if err != nil {
			return err
		}

		if err := Send(apiUrl + "/vehicles", data); err != nil {
			log.Error(err)
			os.Exit(1)
		}

		minInterval := 1000
		maxInterval := 2000

        interval := minInterval + rand.Intn(maxInterval - minInterval)

        log.Info("Sleeping for ", interval, " ms")

        time.Sleep(time.Duration(interval) * time.Millisecond)
	}
}

func main() {
    apiUrl := "http://haproxy:9020/v1"
    receiverUrl := "http://haproxy:9050/v1/data"

    cars, err := GetCars()

	if err != nil {
		log.Error(err)
		os.Exit(1)
	}

    names, err := GetNames()

	if err != nil {
		log.Error(err)
		os.Exit(1)
	}

	var wg sync.WaitGroup

	wg.Add(4)

	go ProcessTemperatures(&wg, receiverUrl)
	go ProcessTransits(&wg, receiverUrl)
	go ProcessUsers(&wg, apiUrl, names)
	go ProcessVehicles(&wg, apiUrl, cars)

	wg.Wait()
}
