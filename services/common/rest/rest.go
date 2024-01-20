package rest

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func ResponseError(c *gin.Context, err string) {
	c.JSON(http.StatusUnprocessableEntity, map[string]string{"message": err})
}

func ResponseNoBody(c *gin.Context, code int) {
	c.Status(code)
}

func ResponseValidationError(c *gin.Context, message string) {
	c.JSON(http.StatusBadRequest, map[string]string{"message": message})
}
