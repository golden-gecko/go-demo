package main

import (
	"bytes"
	"crypto/tls"
	"crypto/x509"
	"encoding/json"
	"fmt"
	"io"
	"math/rand"
	"net/http"
	"os"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"

	log "github.com/sirupsen/logrus"

	"services/common"
	"services/common/metrics"
	"services/common/model"
)

func Send(url string, data []byte) error {
	caCert, err := os.ReadFile("certs/ca-cert.pem")

	if err != nil {
        return err
    }

    caCertPool := x509.NewCertPool()
    caCertPool.AppendCertsFromPEM(caCert)

    client := &http.Client{
        Transport: &http.Transport{
            TLSClientConfig: &tls.Config{
                RootCAs:      caCertPool,
            },
        },
    }

	log.Info(url)

	response, err := client.Post(url, "application/json", bytes.NewBuffer(data))

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
        Timestamp: common.Timestamp(),
    }
}

func CreateTransit() model.Transit {
    c := model.Coordinate{
        Latitude:  14.0745211117 + rand.Float32() * (24.0299857927 - 14.0745211117),
        Longitude: 49.0273953314 + rand.Float32() * (54.8515359564 - 49.0273953314),
    }

    t := model.Transit{
        Plate:     common.RandomPlate(),
        Location:  c,
        Timestamp: common.Timestamp(),
    }

    return t
}

func CreateUser(users model.Users) model.User {
    return model.User{
        Name:     users.Users[rand.Intn(len(users.Users))],
        Password: common.RandomWord(20),
    }
}

func CreateVehicle(cars model.Cars) model.Vehicle {
    car := cars.Cars[rand.Intn(len(cars.Cars))]

    return model.Vehicle{
        Plate:    common.RandomPlate(),
        Brand:    car.Brand,
        Model:    car.Model,
        Year:     common.RandomYear(),
        Category: car.Category,
    }
}

func ProcessItems(receiverUrl string, names model.Names, dataProcessed prometheus.Counter, minInterval int, maxInterval int) {
	for {
		var items []model.Item

		for i := 0; i < 1 + rand.Intn(9); i++ {
			items = append(items, CreateItem(names))

			dataProcessed.Inc()
		}

		data, err := json.Marshal(items)

		if err != nil {
			panic(err)
		}

		if err := Send(receiverUrl + "/item", data); err != nil {
			log.Error(err)
		}

		common.SleepRange(minInterval, maxInterval)
	}
}

func ProcessTemperatures(receiverUrl string, dataProcessed prometheus.Counter, minInterval int, maxInterval int) {
	for {
		t1 := CreateTemperature("Room #1")
		t2 := CreateTemperature("Room #2")
		t3 := CreateTemperature("Room #3")

		dataProcessed.Inc()
		dataProcessed.Inc()
		dataProcessed.Inc()

		data, err := json.Marshal([]model.Temperature{t1, t2, t3})

		if err != nil {
			panic(err)
		}

		if err := Send(receiverUrl + "/temperature", data); err != nil {
			log.Error(err)
		}

		common.SleepRange(minInterval, maxInterval)
	}
}

func ProcessTransits(receiverUrl string, dataProcessed prometheus.Counter, minInterval int, maxInterval int) {
	for {
		var transits []model.Transit

		for i := 0; i < 1 + rand.Intn(9); i++ {
			transits = append(transits, CreateTransit())

			dataProcessed.Inc()
		}

		data, err := json.Marshal(transits)

		if err != nil {
			panic(err)
		}

		if err := Send(receiverUrl + "/transit", data); err != nil {
			log.Error(err)
		}

		common.SleepRange(minInterval, maxInterval)
	}
}

func ProcessUsers(apiUrl string, users model.Users, dataProcessed prometheus.Counter, minInterval int, maxInterval int) {
	for {
		data, err := json.Marshal(CreateUser(users))

		dataProcessed.Inc()

		if err != nil {
			panic(err)
		}

		if err := Send(apiUrl + "/users", data); err != nil {
			log.Error(err)
		}

		common.SleepRange(minInterval, maxInterval)
	}
}

func ProcessVehicles(apiUrl string, cars model.Cars, dataProcessed prometheus.Counter, minInterval int, maxInterval int) {
	for {
		data, err := json.Marshal(CreateVehicle(cars))

		dataProcessed.Inc()

		if err != nil {
			panic(err)
		}

		if err := Send(apiUrl + "/vehicles", data); err != nil {
			log.Error(err)
		}

		common.SleepRange(minInterval, maxInterval)
	}
}

func main() {
    apiUrl := fmt.Sprintf("%s/v1", os.Getenv("API_URL"))
    receiverUrl := fmt.Sprintf("%s/v1/data", os.Getenv("RECEIVER_URL"))

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

	dataProduced := promauto.NewCounter(prometheus.CounterOpts{
		Name: "producer_data_processed",
	})

	var forever chan struct {}

	go metrics.Serve(9106)

	go ProcessItems(receiverUrl, names, dataProduced, 1000, 2000)
	go ProcessTemperatures(receiverUrl, dataProduced, 1000, 2000)
	go ProcessTransits(receiverUrl, dataProduced, 2000, 4000)
	go ProcessUsers(apiUrl, users, dataProduced, 4000, 6000)
	go ProcessVehicles(apiUrl, cars, dataProduced, 6000, 8000)

	<- forever
}
