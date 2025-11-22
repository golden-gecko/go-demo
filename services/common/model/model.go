package model

import (
	"encoding/json"

	log "github.com/sirupsen/logrus"
)

type Car struct {
    Year     int    `bson:"Year"`
    Brand    string `bson:"Brand"`
    Model    string `bson:"Model"`
    Category string `bson:"Category"`
}

type Cars struct {
    Cars []Car `json:"cars"`
}

type Coordinate struct {
    Latitude  float32 `bson:"Latitude"`
    Longitude float32 `bson:"Longitude"`
}

type Item struct {
	Name       string         `bson:"Name"`
    Attributes map[string]int `bson:"Attributes"`
}

type Names struct {
    Names []string `json:"names"`
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

type Users struct {
    Users []string `json:"users"`
}

type Vehicle struct {
    Plate    string `bson:"Plate"`
    Brand    string `bson:"Brand"`
    Model    string `bson:"Model"`
    Year     int    `bson:"Year"`
    Category string `bson:"Category"`
}

type VehicleCountByModel struct {
    Brand string `bson:"Brand"`
    Model string `bson:"Model"`
    Count int    `bson:"Count"`
}

type VehicleCountByYear struct {
    Year  int `bson:"Year"`
    Count int `bson:"Count"`
}

func NewTransit(plate string, latitude float32, longitude float32, timestamp string) Transit {
    return Transit{plate, Coordinate{latitude, longitude}, timestamp}
}

func (transit Transit) ToJSON() ([]byte, error) {
    b, err := json.Marshal(transit)

    if err != nil {
		return nil, err
    }

    return b, nil
}

func (transit Transit) Print() error {
	json, err := transit.ToJSON()

	if err != nil {
		return err
	}

    log.Info(json)

	return nil
}

func NewUser(name string, password string) User {
    return User{name, password}
}

func (user User) ToJSON() ([]byte, error) {
    b, err := json.Marshal(user)

    if err != nil {
		return nil, err
    }

    return b, nil
}

func (user User) Print() error {
	json, err := user.ToJSON()

	if err != nil {
		return err
	}

    log.Info(json)

	return nil
}

func NewVehicle(plate string, brand string, model string, year int, category string) Vehicle {
    return Vehicle{plate, brand, model, year, category}
}

func (vehicle Vehicle) ToJSON() ([]byte, error) {
    b, err := json.Marshal(vehicle)

    if err != nil {
		return nil, err
    }

    return b, nil
}

func (vehicle Vehicle) Print() error {
	json, err := vehicle.ToJSON()

	if err != nil {
		return err
	}

    log.Info(json)

	return nil
}
