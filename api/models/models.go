package models

import (
	"encoding/json"
	"log"
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

func (transit Transit) ToJSON() []byte {
	b, err := json.Marshal(transit)

	if err != nil {
		panic(err)
	}

	return b
}

func (transit Transit) Print() {
	log.Println(string(transit.ToJSON()))
}

// ---------------------------------------------------------------------------

type User struct {
	Name string `bson:"Name"`
}

func NewUser(name string) User {
	return User{name}
}

func (user User) ToJSON() []byte {
	b, err := json.Marshal(user)

	if err != nil {
		panic(err)
	}

	return b
}

func (user User) Print() {
	log.Println(string(user.ToJSON()))
}

// ---------------------------------------------------------------------------

type Vehicle struct {
	Plate string `bson:"Plate"`
	Brand string `bson:"Brand"`
	Model string `bson:"Model"`
	Year  int    `bson:"Year"`
}

func NewVehicle(plate string, brand string, model string, year int) Vehicle {
	return Vehicle{plate, brand, model, year}
}

func (vehicle Vehicle) ToJSON() []byte {
	b, err := json.Marshal(vehicle)

	if err != nil {
		panic(err)
	}

	return b
}

func (vehicle Vehicle) Print() {
	log.Println(string(vehicle.ToJSON()))
}
