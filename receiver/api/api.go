package api

import (
	"encoding/json"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"

	amqp "github.com/rabbitmq/amqp091-go"
	kafka "github.com/segmentio/kafka-go"

	"receiver/queue"
)

type Coordinate struct {
	Latitude  float32 `bson:"Latitude"`
	Longitude float32 `bson:"Longitude"`
}

type Temperature struct {
	Location  string  `bson:"Location"`
	Value     float32 `bson:"Value"`
	Timestamp string  `bson:"Timestamp"`
}

type Transit struct {
	Plate     string     `bson:"Plate"`
	Location  Coordinate `bson:"Location"`
	Timestamp string     `bson:"Timestamp"`
}

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

func WriteToRabbit(queueName string, data []byte) {
	err := queue.Channel.PublishWithContext(
		queue.Context,
		"",
		queueName,
		false, // mandatory
		false, // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        data,
		},
	)

	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
}

func Healthcheck(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, nil)
}

func CreateData(c *gin.Context) {
	switch c.Param("type") {
	case "temperature":
		var temperatures []Temperature

		if err := c.ShouldBind(&temperatures); err != nil {
			log.Println(err)
		} else {
			// log.Println(temperatures)

			for _, temperature := range temperatures {
				data, err := json.Marshal(&temperature)

				// log.Println(temperature)
				// log.Println(data)

				if err != nil {
					os.Exit(1)
				}

				WriteToRabbit("temperatures", data)

				c.IndentedJSON(http.StatusAccepted, nil)
			}
		}
	case "transit":
		var transits []Transit

		if err := c.ShouldBind(&transits); err != nil {
			log.Println(err)
		} else {
			// log.Println(transits)

			for _, transit := range transits {
				data, err := json.Marshal(&transit)

				// log.Println(transit)
				// log.Println(data)

				if err != nil {
					log.Println(err)
					os.Exit(1)
				}

				WriteToRabbit("transits", data)

				c.IndentedJSON(http.StatusAccepted, nil)
			}
		}
	default:
		c.IndentedJSON(http.StatusBadRequest, nil)
	}
}
