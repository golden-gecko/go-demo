package main

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"math/rand"
	"net/http"
	"os"
	"time"

	log "github.com/sirupsen/logrus"

	"producer/models"
)

func RandomString(length int, charset string) string {
	b := make([]byte, length)

	for i := range b {
		b[i] = charset[rand.Intn(len(charset))]
	}

	return string(b)
}

func RandomBrand() string {
	brands := []string{
		"Abarth",
		"Alfa Romeo",
		"Aston Martin",
		"Audi",
		"Bentley",
		"BMW",
		"Bugatti",
		"Cadillac",
		"Chevrolet",
		"Chrysler",
		"Citroën",
		"Dacia",
		"Daewoo",
		"Daihatsu",
		"Dodge",
		"Donkervoort",
		"DS",
		"Ferrari",
		"Fiat",
		"Fisker",
		"Ford",
		"Honda",
		"Hummer",
		"Hyundai",
		"Infiniti",
		"Iveco",
		"Jaguar",
		"Jeep",
		"Kia",
		"KTM",
		"Lada",
		"Lamborghini",
		"Lancia",
		"Land Rover",
		"Landwind",
		"Lexus",
		"Lotus",
		"Maserati",
		"Maybach",
		"Mazda",
		"McLaren",
		"Mercedes-Benz",
		"MG",
		"Mini",
		"Mitsubishi",
		"Morgan",
		"Nissan",
		"Opel",
		"Peugeot",
		"Porsche",
		"Renault",
		"Rolls-Royce",
		"Rover",
		"Saab",
		"Seat",
		"Skoda",
		"Smart",
		"SsangYong",
		"Subaru",
		"Suzuki",
		"Tesla",
		"Toyota",
		"Volkswagen",
		"Volvo",
	}

	return brands[rand.Intn(len(brands))]
}

func RandomName() string {
	names := []string{
		"Alanh",
		"Amanda",
		"Amy",
		"Betty",
		"Bruce",
		"Becky",
		"Collins",
		"Craig",
		"Clark",
		"Donald",
		"David",
		"Drew",
		"Emily",
		"Edward",
		"Eric",
		"Frank",
		"Finch",
		"Francis",
		"Gary",
		"Glen",
		"George",
		"Henry",
		"Howard",
		"Helen",
		"Irene",
		"Iris",
		"Ivan",
		"Jack",
		"John",
		"Jenny",
		"Karen",
		"Kim",
		"Kenneth",
		"Larry",
		"Leslie",
		"Lisa",
		"Martin",
		"Mary",
		"Miller",
		"Nick",
		"Nancy",
		"Nathan",
		"Oliver",
		"Octavio",
		"Oscar",
		"Peter",
		"Paula",
		"Pamela",
		"Quin",
		"Ronald",
		"Ron",
		"Randy",
		"Samnuel",
		"Scott",
		"Sharon",
		"Teresa",
		"Terry",
		"Thomas",
		"Uriel",
		"Usher",
		"Victor",
		"Virginia",
		"Watson",
		"Wendy",
		"Wood",
		"Xavier",
		"Yasmin",
		"Yvette",
		"Yuri",
		"Zack",
		"Zelda",
		"Zane",
	}

	return names[rand.Intn(len(names))]
}

func RandomPlate() string {
	const characters = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	const digits = "0123456789"

	return RandomString(2, characters) + RandomString(3, digits)
}

func RandomWord(length int) string {
	const characters = "abcdefghijklmnopqrstuvwxyz"

	return RandomString(length, characters)
}

func Send(url string, data []byte) error {
	body := bytes.NewBuffer(data)
	resp, err := http.Post(url, "application/json", body)

	if err != nil {
		log.Error(err)
		return err
	}

	if resp.StatusCode == 200 || resp.StatusCode == 201 || resp.StatusCode == 202 {
		log.Println(resp.StatusCode)
	} else {
		log.Error(resp.StatusCode)
	}

	return nil
}

func CreateUser() models.User {
	u := models.User{
		Name:     RandomName(),
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
		Year:     1980 + rand.Intn(30),
		Category: car.Category,
	}

	return t
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
		Latitude:  14.0745211117 + rand.Float32()*(24.0299857927-14.0745211117),
		Longitude: 49.0273953314 + rand.Float32()*(54.8515359564-49.0273953314),
	}

	t := models.Transit{
		Plate:     RandomPlate(),
		Location:  c,
		Timestamp: time.Now().Format(time.RFC3339),
	}

	return t
}

func GetCars() models.Cars {
	jsonFile, err := os.Open("data/cars.json")

	if err != nil {
		log.Error(err)
		os.Exit(1)
	}

	defer jsonFile.Close()

	var cars models.Cars

	byteValue, _ := ioutil.ReadAll(jsonFile)

	json.Unmarshal(byteValue, &cars)

	return cars
}

func main() {
	apiUrl := "http://haproxy:9020/v1"
	receiverUrl := "http://haproxy:9050/v1/data"

	cars := GetCars()

	minInterval := 100
	maxInterval := 500

	for {
		if true {
			u1 := CreateUser()

			data, err := json.Marshal(u1)

			if err != nil {
				log.Error(err)
				os.Exit(1)
			}

			Send(apiUrl+"/users", data)
		}

		if true {
			v1 := CreateVehicle(cars)

			data, err := json.Marshal(v1)

			if err != nil {
				log.Error(err)
				os.Exit(1)
			}

			Send(apiUrl+"/vehicles", data)
		}

		if true {
			t1 := CreateTemperature("Room #1")
			t2 := CreateTemperature("Room #2")
			t3 := CreateTemperature("Room #3")

			data, err := json.Marshal([]models.Temperature{t1, t2, t3})

			if err != nil {
				log.Error(err)
				os.Exit(1)
			}

			Send(receiverUrl+"/temperature", data)
		}

		if true {
			var transits []models.Transit

			for i := 0; i < 1+rand.Intn(9); i++ {
				transits = append(transits, CreateTransit())
			}

			data, err := json.Marshal(transits)

			if err != nil {
				log.Error(err)
				os.Exit(1)
			}

			Send(receiverUrl+"/transit", data)
		}

		sleep := minInterval + rand.Intn(maxInterval-minInterval)

		log.Debug("Sleeping for", sleep, "ms")

		time.Sleep(time.Duration(sleep) * time.Millisecond)
	}
}
