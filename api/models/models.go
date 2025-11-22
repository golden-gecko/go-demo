package models

import (
	"encoding/json"

	log "github.com/sirupsen/logrus"
)

// ---------------------------------------------------------------------------

type Coordinate struct {
    Latitude  float32 `bson:"Latitude"`
    Longitude float32 `bson:"Longitude"`
}

type Transit struct {
    Plate     string     `bson:"Plate"`
    Location  Coordinate `bson:"Location"`
    Timestamp string     `bson:"Timestamp"`
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

    log.Println(json)

	return nil
}

// ---------------------------------------------------------------------------

type User struct {
    Name     string `bson:"Name"`
    Password string `bson:"Password"`
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

    log.Println(json)

	return nil
}

// ---------------------------------------------------------------------------

type Vehicle struct {
    Plate    string `bson:"Plate"`
    Brand    string `bson:"Brand"`
    Model    string `bson:"Model"`
    Year     int    `bson:"Year"`
    Category string `bson:"Category"`
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

    log.Println(json)

	return nil
}
