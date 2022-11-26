package api

import (
	"context"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"

	amqp "github.com/rabbitmq/amqp091-go"
	kafka "github.com/segmentio/kafka-go"

	"api/db"
	"api/models"
	"api/queue"
)

// ---------------------------------------------------------------------------

func WriteToKafka(data []byte) {
	_, err := queue.KafkaClient.WriteMessages(
		kafka.Message{
			Value: data,
		},
	)

	if err != nil {
		panic(err)
	}
}

func WriteToRabbit(data []byte) {
	maxRetries := 3

	for i := 0; i < maxRetries; i++ {
		err := queue.Channel.PublishWithContext(
			queue.Context,
			"",
			queue.Queue.Name,
			false, // mandatory
			false, // immediate
			amqp.Publishing{
				ContentType: "text/plain",
				Body:        data,
			},
		)

		if err != nil {
			log.Println("Failed to send message. Waiting 3 seconds...")

			time.Sleep(3 * time.Second)

			queue.InitRabbit()
		} else {
			return
		}
	}

	log.Println("Failed to send message 3 times. Exiting...")

	os.Exit(1)
}

func CreateData(c *gin.Context) {
	data, err := ioutil.ReadAll(c.Request.Body)

	if err != nil {
		log.Println(err)
	}

	// WriteToKafka(data)
	WriteToRabbit(data)

	c.IndentedJSON(http.StatusCreated, nil)
}

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
		panic(err)
	}

	_, err := db.UserCollection.InsertOne(context.TODO(), user)

	if err != nil {
		panic(err)
	}

	c.IndentedJSON(http.StatusCreated, nil)
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
