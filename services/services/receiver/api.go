package main

import (
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"

	amqp "github.com/rabbitmq/amqp091-go"
	kafka "github.com/segmentio/kafka-go"

	"services/common/queue"
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

func WriteToKafka(data []byte) error {
    _, err := queue.KafkaClient.WriteMessages(
        kafka.Message{
            Value: data,
        },
    )

	return err
}

func WriteToRabbit(queueName string, data []byte) error {
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

	return err
}

func ResponseNoBody(c *gin.Context, code int) {
	c.JSON(code, nil)
}

func ResponseError(c *gin.Context, code int, err string) {
	c.JSON(code, map[string]string{"message": err})
}

func Healthcheck(c *gin.Context) {
    c.IndentedJSON(http.StatusOK, nil)
}

func CreateDataTemperature(c *gin.Context) {
	var temperatures []Temperature

	if err := c.ShouldBind(&temperatures); err != nil {
		ResponseError(c, http.StatusUnprocessableEntity, err.Error())
		return
	}

	for _, temperature := range temperatures {
		data, err := json.Marshal(&temperature)

		if err != nil {
			ResponseError(c, http.StatusUnprocessableEntity, err.Error())
			return
		}

		if err := WriteToRabbit("temperatures", data); err != nil {
			ResponseError(c, http.StatusUnprocessableEntity, err.Error())
			return
		}
	}

	c.IndentedJSON(http.StatusAccepted, nil)
}

func CreateDataTransit(c *gin.Context) {
	var transits []Transit

	if err := c.ShouldBind(&transits); err != nil {
		ResponseError(c, http.StatusUnprocessableEntity, err.Error())
		return
	}

	for _, transit := range transits {
		data, err := json.Marshal(&transit)

		if err != nil {
			ResponseError(c, http.StatusUnprocessableEntity, err.Error())
			return
		}

		if err := WriteToRabbit("transits", data); err != nil {
			ResponseError(c, http.StatusUnprocessableEntity, err.Error())
			return
		}
	}

	c.IndentedJSON(http.StatusAccepted, nil)
}

func CreateData(c *gin.Context) {
    switch c.Param("type") {
		case "temperature":
			CreateDataTemperature(c)

		case "transit":
			CreateDataTransit(c)

		default:
			c.IndentedJSON(http.StatusBadRequest, nil)
    }
}
