package models

import (
    "testing"

    "github.com/stretchr/testify/assert"
)

func Test_UserToJson(t *testing.T) {
    user := NewUser("Jack", "abc")

    assert := assert.New(t)
    assert.Equal(string(user.ToJSON()), `{"Name":"Jack","Password":"abc"}`)
}

func Test_VehicleToJson(t *testing.T) {
    user := NewVehicle("ABC123", "Chevrolet", "Camaro", 1992)

    assert := assert.New(t)
    assert.Equal(string(user.ToJSON()), `{"Plate":"ABC123","Brand":"Chevrolet","Model":"Camaro","Year":1992}`)
}

func Test_VehicleTransitToJson(t *testing.T) {
    user := NewTransit("ABC123", 0.12, 0.34, "2022-01-01T12:16:34")

    assert := assert.New(t)
    assert.Equal(string(user.ToJSON()), `{"Plate":"ABC123","Location":{"Latitude":0.12,"Longitude":0.34},"Timestamp":"2022-01-01T12:16:34"}`)
}
