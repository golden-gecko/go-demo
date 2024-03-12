package main

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"path/filepath"

	log "github.com/sirupsen/logrus"

	"services/common/model"
)

func OpenFile(filePath string) (*os.File, error) {
	log.Info(fmt.Sprintf("Opening file %s...", filePath))

	return os.Open(filePath)
}

func GetCars(dataPath string) (model.Cars, error) {
	var cars model.Cars

	jsonFile, err := OpenFile(filepath.Join(dataPath, "cars.json"))

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

func GetNames(dataPath string) (model.Names, error) {
	var names model.Names

	jsonFile, err := OpenFile(filepath.Join(dataPath, "names.json"))

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

func GetUsers(dataPath string) (model.Users, error) {
	var users model.Users

	jsonFile, err := OpenFile(filepath.Join(dataPath, "users.json"))

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
