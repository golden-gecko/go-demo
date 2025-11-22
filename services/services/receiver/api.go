package main

import (
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"

	amqp "github.com/rabbitmq/amqp091-go"

	"services/common/model"
	"services/common/queue"
	"services/common/rest"
)

/*
func WriteToKafka(data []byte) error {
    _, err := queue.KafkaClient.WriteMessages(
        kafka.Message{
            Value: data,
        },
    )

	return err
}
*/

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

func Healthcheck(c *gin.Context) {
    c.IndentedJSON(http.StatusOK, nil)
}

func CreateDataTemperature(c *gin.Context) {
	var temperatures []model.Temperature

	if err := c.ShouldBind(&temperatures); err != nil {
		rest.ResponseError(c, err.Error())
		return
	}

	for _, temperature := range temperatures {
		data, err := json.Marshal(&temperature)

		if err != nil {
			rest.ResponseError(c, err.Error())
			return
		}

		if err := WriteToRabbit("temperatures", data); err != nil {
			rest.ResponseError(c, err.Error())
			return
		}
	}

	c.IndentedJSON(http.StatusAccepted, nil)
}

func CreateDataTransit(c *gin.Context) {
	var transits []model.Transit

	if err := c.ShouldBind(&transits); err != nil {
		rest.ResponseError(c, err.Error())
		return
	}

	for _, transit := range transits {
		data, err := json.Marshal(&transit)

		if err != nil {
			rest.ResponseError(c, err.Error())
			return
		}

		if err := WriteToRabbit("transits", data); err != nil {
			rest.ResponseError(c, err.Error())
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
			rest.ResponseValidationError(c, "invalid value: type")
			c.IndentedJSON(http.StatusBadRequest, nil)
    }
}
