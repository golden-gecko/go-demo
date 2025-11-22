package api

import (
	"context"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"

	"api/db"
	"api/models"
	"api/roach"
)

// ---------------------------------------------------------------------------

func CreateTransit(c *gin.Context) {
	var transit models.Transit

	transit.Plate = c.Param("plate")

	if err := c.BindJSON(&transit); err != nil {
		panic(err)
	}

	_, err := db.TransitCollection.InsertOne(context.TODO(), transit)

	if err != nil {
		panic(err)
	}

	c.IndentedJSON(http.StatusCreated, nil)
}

// ---------------------------------------------------------------------------

func CreateUser(c *gin.Context) {
	var user models.User

	if err := c.BindJSON(&user); err != nil {
		log.Fatalln(err)
		c.JSON(http.StatusInternalServerError, map[string]error{"message": err})
	}

	if err := roach.CreateUser(user); err != nil {
		log.Fatalln(err)
		c.JSON(http.StatusInternalServerError, map[string]error{"message": err})
	}

	/*_, err := db.UserCollection.InsertOne(context.TODO(), user)

	if err != nil {
		log.Fatalln(err)
	}*/

	c.JSON(http.StatusCreated, map[string]error{})
}

func DeleteUser(c *gin.Context) {
	c.JSON(http.StatusOK, map[string]error{})
}

func GetUsers(c *gin.Context) {
	var users []models.User

	cursor := db.Find(db.UserCollection)
	defer cursor.Close(context.TODO())

	for cursor.Next(context.TODO()) {
		var user1 bson.D
		var user3 models.User

		if err := cursor.Decode(&user1); err != nil {
			panic(err)
		}

		user2, err := bson.Marshal(user1)

		if err != nil {
			panic(err)
		}

		if err := bson.Unmarshal(user2, &user3); err != nil {
			panic(err)
		}

		users = append(users, user3)
	}

	c.IndentedJSON(http.StatusOK, users)
}

func GetUser(c *gin.Context) {
	// userID := c.Param("userID")
}

// ---------------------------------------------------------------------------

func CreateVehicle(c *gin.Context) {
	var vehicle models.Vehicle

	if err := c.BindJSON(&vehicle); err != nil {
		panic(err)
	}

	_, err := db.VehicleCollection.InsertOne(context.TODO(), vehicle)

	if err != nil {
		panic(err)
	}

	c.IndentedJSON(http.StatusCreated, nil)
}

func GetVehicles(c *gin.Context) {
	var vehicles []models.Vehicle

	cursor := db.Find(db.VehicleCollection)
	defer cursor.Close(context.TODO())

	for cursor.Next(context.TODO()) {
		var vehicle1 bson.D
		var vehicle3 models.Vehicle

		if err := cursor.Decode(&vehicle1); err != nil {
			panic(err)
		}

		vehicle2, err := bson.Marshal(vehicle1)

		if err != nil {
			panic(err)
		}

		if err := bson.Unmarshal(vehicle2, &vehicle3); err != nil {
			panic(err)
		}

		vehicles = append(vehicles, vehicle3)
	}

	c.IndentedJSON(http.StatusOK, vehicles)
}

func GetVehicle(c *gin.Context) {
	// vehicleID := c.Param("vehicleID")
}
