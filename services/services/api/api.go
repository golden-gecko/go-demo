package main

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"

	"services/common/model"
	"services/common/mongo"
	"services/common/rest"
	"services/common/sanitize"
	"services/common/sql"
	"services/common/validate"
)

func Healthcheck(c *gin.Context) {
    c.IndentedJSON(http.StatusOK, nil)
}

func CreateTransit(c *gin.Context) {
    var transit model.Transit

    transit.Plate = sanitize.String(c.Param("plate"))

    if validate.NonEmptyString(transit.Plate) == false {
		rest.ResponseValidationError(c, "invalid value: no name")
        return
    }

	if err := c.BindJSON(&transit); err != nil {
		rest.ResponseError(c, err.Error())
		return
    }

    if _, err := mongo.TransitCollection.InsertOne(context.TODO(), transit); err != nil {
		rest.ResponseError(c, err.Error())
		panic(err)
    }

	rest.ResponseNoBody(c, http.StatusCreated)
}

func CreateUser(c *gin.Context) {
    var user model.User

    if err := c.BindJSON(&user); err != nil {
		rest.ResponseError(c, err.Error())
        return
    }

    if validate.NonEmptyString(user.Name) == false {
		rest.ResponseValidationError(c, "invalid value: no name")
        return
    }

    if validate.NonEmptyString(user.Password) == false {
		rest.ResponseValidationError(c, "invalid value: no password")
        return
    }

    if err := sql.CreateUser(user); err != nil {
		rest.ResponseError(c, err.Error())
		panic(err)
    }

	rest.ResponseNoBody(c, http.StatusCreated)
}

func GetUsers(c *gin.Context) {
    var users []model.User

    cursor, err := mongo.Find(mongo.UserCollection)

	if err != nil {
		rest.ResponseError(c, err.Error())
		return
	}

    defer cursor.Close(context.TODO())

    for cursor.Next(context.TODO()) {
        var user1 bson.D
        var user3 model.User

        if err := cursor.Decode(&user1); err != nil {
			rest.ResponseError(c, err.Error())
			return
        }

        user2, err := bson.Marshal(user1)

        if err != nil {
			rest.ResponseError(c, err.Error())
			return
        }

        if err := bson.Unmarshal(user2, &user3); err != nil {
			rest.ResponseError(c, err.Error())
			return
        }

        users = append(users, user3)
    }

    c.IndentedJSON(http.StatusOK, users)
}

func GetUser(c *gin.Context) {
    userId, err := uuid.Parse(c.Param("userId"))

    if err != nil {
		rest.ResponseError(c, err.Error())
		return
    }

    user, err := sql.GetUser(userId)

    if err != nil {
		rest.ResponseError(c, err.Error())
		return
    }

    c.IndentedJSON(http.StatusOK, user)
}

func CreateVehicle(c *gin.Context) {
    var vehicle model.Vehicle

    if err := c.BindJSON(&vehicle); err != nil {
		rest.ResponseError(c, err.Error())
		return
    }

    if validate.NonEmptyString(vehicle.Plate) == false {
		rest.ResponseValidationError(c, "invalid value: no plate")
		return
    }

    if validate.NonEmptyString(vehicle.Brand) == false {
		rest.ResponseValidationError(c, "invalid value: no brand")
		return
    }

    if validate.NonEmptyString(vehicle.Model) == false {
		rest.ResponseValidationError(c, "invalid value: no model")
		return
    }

    if validate.NonEmptyString(vehicle.Category) == false {
		rest.ResponseValidationError(c, "invalid value: no category")
		return
    }

    if err := sql.CreateVehicle(vehicle); err != nil {
		rest.ResponseError(c, err.Error())
		panic(err)
    }

	rest.ResponseNoBody(c, http.StatusCreated)
}

func GetVehicles(c *gin.Context) {
    var vehicles []model.Vehicle

    cursor, err := mongo.Find(mongo.VehicleCollection)
	
	if err != nil {
		rest.ResponseError(c, err.Error())
		return
	}

    defer cursor.Close(context.TODO())

    for cursor.Next(context.TODO()) {
        var vehicle1 bson.D
        var vehicle3 model.Vehicle

        if err := cursor.Decode(&vehicle1); err != nil {
			rest.ResponseError(c, err.Error())
			return
        }

        vehicle2, err := bson.Marshal(vehicle1)

        if err != nil {
			rest.ResponseError(c, err.Error())
			return
        }

        if err := bson.Unmarshal(vehicle2, &vehicle3); err != nil {
			rest.ResponseError(c, err.Error())
			return
        }

        vehicles = append(vehicles, vehicle3)
    }

    c.IndentedJSON(http.StatusOK, vehicles)
}
