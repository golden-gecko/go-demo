package model

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_UserToJson(t *testing.T) {
    user := NewUser("Jack", "abc")
	userJson, _ := user.ToJSON()

    assert := assert.New(t)
    assert.Equal(string(userJson), `{"Name":"Jack","Password":"abc"}`)
}

func Test_VehicleToJson(t *testing.T) {
    user := NewVehicle("ABC123", "Chevrolet", "Camaro", 1992, "SUV")
	userJson, _ := user.ToJSON()

    assert := assert.New(t)
    assert.Equal(string(userJson), `{"Plate":"ABC123","Brand":"Chevrolet","Model":"Camaro","Year":1992,"Category":"SUV"}`)
}

func Test_VehicleTransitToJson(t *testing.T) {
    user := NewTransit("ABC123", 0.12, 0.34, "2022-01-01T12:16:34")
	userJson, _ := user.ToJSON()

    assert := assert.New(t)
    assert.Equal(string(userJson), `{"Plate":"ABC123","Location":{"Latitude":0.12,"Longitude":0.34},"Timestamp":"2022-01-01T12:16:34"}`)
}
