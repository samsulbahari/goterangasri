package helper

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func Validatorerror(c *gin.Context, err error) {

	var jsErr *json.UnmarshalTypeError
	if errors.As(err, &jsErr) {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": "the json is invalid",
		})
	}
	errorMessages := []string{}
	for _, e := range err.(validator.ValidationErrors) {
		errorMessage := fmt.Sprintf("error on field %s , condition  %s", e.Field(), e.ActualTag())
		errorMessages = append(errorMessages, errorMessage)
	}
	c.JSON(http.StatusBadRequest, gin.H{
		"status":  false,
		"message": errorMessages,
	})
	return
}
