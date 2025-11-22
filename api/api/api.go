package api

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"

	log "github.com/sirupsen/logrus"

	"api/db"
	"api/models"
	"api/roach"
)

// ---------------------------------------------------------------------------

func Healthcheck(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, nil)
}

// ---------------------------------------------------------------------------

func CreateTransit(c *gin.Context) {
	var transit models.Transit

	transit.Plate = c.Param("plate")

	if err := c.BindJSON(&transit); err != nil {
		log.Error(err)
		panic(err)
	}

	_, err := db.TransitCollection.InsertOne(context.TODO(), transit)

	if err != nil {
		log.Error(err)
		panic(err)
	}

	c.IndentedJSON(http.StatusCreated, nil)
}

// ---------------------------------------------------------------------------

func CreateUser(c *gin.Context) {
	var user models.User

	if err := c.BindJSON(&user); err != nil {
		log.Error(err)
		c.JSON(http.StatusInternalServerError, map[string]error{"message": err})
	}

	if err := roach.CreateUser(user); err != nil {
		log.Error(err)
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
			log.Error(err)
			panic(err)
		}

		user2, err := bson.Marshal(user1)

		if err != nil {
			log.Error(err)
			panic(err)
		}

		if err := bson.Unmarshal(user2, &user3); err != nil {
			log.Error(err)
			panic(err)
		}

		users = append(users, user3)
	}

	c.IndentedJSON(http.StatusOK, users)
}

func GetUser(c *gin.Context) {
	userId, err := uuid.Parse(c.Param("userId"))

	if err != nil {
		log.Error(err)
		c.IndentedJSON(http.StatusBadRequest, err)
		return
	}

	user, err := roach.GetUser(userId)

	if err != nil {
		log.Error(err)
		c.IndentedJSON(http.StatusUnprocessableEntity, err)
		return
	}

	c.IndentedJSON(http.StatusUnprocessableEntity, user)
}

// ---------------------------------------------------------------------------

func CreateVehicle(c *gin.Context) {
	var vehicle models.Vehicle

	if err := c.BindJSON(&vehicle); err != nil {
		log.Error(err)
		c.JSON(http.StatusInternalServerError, map[string]error{"message": err})
	}

	if err := roach.CreateVehicle(vehicle); err != nil {
		log.Error(err)
		c.JSON(http.StatusInternalServerError, map[string]error{"message": err})
	}

	c.JSON(http.StatusCreated, map[string]error{})
}

func GetVehicles(c *gin.Context) {
	var vehicles []models.Vehicle

	cursor := db.Find(db.VehicleCollection)
	defer cursor.Close(context.TODO())

	for cursor.Next(context.TODO()) {
		var vehicle1 bson.D
		var vehicle3 models.Vehicle

		if err := cursor.Decode(&vehicle1); err != nil {
			log.Error(err)
			panic(err)
		}

		vehicle2, err := bson.Marshal(vehicle1)

		if err != nil {
			log.Error(err)
			panic(err)
		}

		if err := bson.Unmarshal(vehicle2, &vehicle3); err != nil {
			log.Error(err)
			panic(err)
		}

		vehicles = append(vehicles, vehicle3)
	}

	c.IndentedJSON(http.StatusOK, vehicles)
}

func GetVehicle(c *gin.Context) {
	// vehicleID := c.Param("vehicleID")
}
