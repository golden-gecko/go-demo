package main

import (
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"

	"services/common/model"
	"services/common/rabbit"
	"services/common/rest"
)

func Healthcheck(c *gin.Context) {
    c.IndentedJSON(http.StatusOK, nil)
}

func CreateDataItem(c *gin.Context) {
	var items []model.Item

	if err := c.ShouldBind(&items); err != nil {
		rest.ResponseError(c, err.Error())
		return
	}

	for _, item := range items {
		data, err := json.Marshal(&item)

		if err != nil {
			rest.ResponseError(c, err.Error())
			return
		}

		if err := rabbit.Write("items", data); err != nil {
			rest.ResponseError(c, err.Error())
			return
		}
	}

	c.IndentedJSON(http.StatusAccepted, nil)
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

		if err := rabbit.Write("temperatures", data); err != nil {
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

		if err := rabbit.Write("transits", data); err != nil {
			rest.ResponseError(c, err.Error())
			return
		}
	}

	c.IndentedJSON(http.StatusAccepted, nil)
}

func CreateData(c *gin.Context) {
    switch c.Param("type") {
		case "item":
			CreateDataItem(c)

		case "temperature":
			CreateDataTemperature(c)

		case "transit":
			CreateDataTransit(c)

		default:
			rest.ResponseValidationError(c, "invalid value: type")
			c.IndentedJSON(http.StatusBadRequest, nil)
    }
}
