package main

import (
	"encoding/json"
	"io"
	"os"
	"path/filepath"

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

func GetUsers() (model.Users, error) {
	var users model.Users

	jsonFile, err := os.Open(filepath.Join("services", "producer", "data", "users.json"))

	if err != nil {
		return users, err
	}

	defer jsonFile.Close()

	byteValue, err := io.ReadAll(jsonFile)

	if err != nil {
		return users, err
	}

	if json.Unmarshal(byteValue, &users) != nil {
		return users, err
	}

	return users, nil
}
