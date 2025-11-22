package api

import (
	"context"
	"net/http"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"

	"api/db"
	"api/models"
	"api/roach"
)

// ---------------------------------------------------------------------------

func ValidateString(value string) bool {
    return len(strings.Trim(value, " ")) > 0
}

// ---------------------------------------------------------------------------

func Healthcheck(c *gin.Context) {
    c.IndentedJSON(http.StatusOK, nil)
}

// ---------------------------------------------------------------------------

func ResponseNoBody(c *gin.Context, code int) {
	c.JSON(code, nil)
}

func ResponseError(c *gin.Context, code int, err string) {
	c.JSON(code, map[string]string{"message": err})
}

// ---------------------------------------------------------------------------

func CreateTransit(c *gin.Context) {
    var transit models.Transit

    transit.Plate = c.Param("plate")

    if ValidateString(transit.Plate) == false {
		ResponseError(c, http.StatusBadRequest, "invalid value: no name")
        return
    }

	if err := c.BindJSON(&transit); err != nil {
		ResponseError(c, http.StatusUnprocessableEntity, err.Error())
		return
    }

    if _, err := db.TransitCollection.InsertOne(context.TODO(), transit); err != nil {
		ResponseError(c, http.StatusUnprocessableEntity, err.Error())
		os.Exit(1)
    }

	ResponseNoBody(c, http.StatusCreated)
}

// ---------------------------------------------------------------------------

func CreateUser(c *gin.Context) {
    var user models.User

    if err := c.BindJSON(&user); err != nil {
		ResponseError(c, http.StatusUnprocessableEntity, err.Error())
        return
    }

    if ValidateString(user.Name) == false {
		ResponseError(c, http.StatusBadGateway, "invalid value: no name")
        return
    }

    if ValidateString(user.Password) == false {
		ResponseError(c, http.StatusBadGateway, "invalid value: no password")
        return
    }

    if err := roach.CreateUser(user); err != nil {
		ResponseError(c, http.StatusUnprocessableEntity, err.Error())
		os.Exit(1)
    }

	ResponseNoBody(c, http.StatusCreated)
}

func GetUsers(c *gin.Context) {
    var users []models.User

    cursor := db.Find(db.UserCollection)
    defer cursor.Close(context.TODO())

    for cursor.Next(context.TODO()) {
        var user1 bson.D
        var user3 models.User

        if err := cursor.Decode(&user1); err != nil {
			ResponseError(c, http.StatusUnprocessableEntity, err.Error())
			return
        }

        user2, err := bson.Marshal(user1)

        if err != nil {
			ResponseError(c, http.StatusUnprocessableEntity, err.Error())
			return
        }

        if err := bson.Unmarshal(user2, &user3); err != nil {
			ResponseError(c, http.StatusUnprocessableEntity, err.Error())
			return
        }

        users = append(users, user3)
    }

    c.IndentedJSON(http.StatusOK, users)
}

func GetUser(c *gin.Context) {
    userId, err := uuid.Parse(c.Param("userId"))

    if err != nil {
		ResponseError(c, http.StatusBadRequest, err.Error())
		return
    }

    user, err := roach.GetUser(userId)

    if err != nil {
		ResponseError(c, http.StatusUnprocessableEntity, err.Error())
		return
    }

    c.IndentedJSON(http.StatusOK, user)
}

// ---------------------------------------------------------------------------

func CreateVehicle(c *gin.Context) {
    var vehicle models.Vehicle

    if err := c.BindJSON(&vehicle); err != nil {
		ResponseError(c, http.StatusUnprocessableEntity, err.Error())
		return
    }

    if ValidateString(vehicle.Plate) == false {
		ResponseError(c, http.StatusBadRequest, "invalid value: no plate")
		return
    }

    if ValidateString(vehicle.Brand) == false {
		ResponseError(c, http.StatusBadRequest, "invalid value: no brand")
		return
    }

    if ValidateString(vehicle.Model) == false {
		ResponseError(c, http.StatusBadRequest, "invalid value: no model")
		return
    }

    if ValidateString(vehicle.Category) == false {
		ResponseError(c, http.StatusBadRequest,"invalid value: no category")
		return
    }

    if err := roach.CreateVehicle(vehicle); err != nil {
		ResponseError(c, http.StatusUnprocessableEntity, err.Error())
		os.Exit(1)
    }

	ResponseNoBody(c, http.StatusCreated)
}

func GetVehicles(c *gin.Context) {
    var vehicles []models.Vehicle

    cursor := db.Find(db.VehicleCollection)
    defer cursor.Close(context.TODO())

    for cursor.Next(context.TODO()) {
        var vehicle1 bson.D
        var vehicle3 models.Vehicle

        if err := cursor.Decode(&vehicle1); err != nil {
			ResponseError(c, http.StatusUnprocessableEntity, err.Error())
			return
        }

        vehicle2, err := bson.Marshal(vehicle1)

        if err != nil {
			ResponseError(c, http.StatusUnprocessableEntity, err.Error())
			return
        }

        if err := bson.Unmarshal(vehicle2, &vehicle3); err != nil {
			ResponseError(c, http.StatusUnprocessableEntity, err.Error())
			return
        }

        vehicles = append(vehicles, vehicle3)
    }

    c.IndentedJSON(http.StatusOK, vehicles)
}
