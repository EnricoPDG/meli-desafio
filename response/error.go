package response

import (
	"github.com/EnricoPDG/meli-desafio/logger"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
	"time"
)

var log = logger.GetLogger()

type JSONErrorResponse struct {
	Message   string `json:"message" example:"something went wrong"`
	Timestamp string `json:"timestamp" example:"2025-06-05T18:44:26.131446-03:00"`
}

func BadRequestResponse(c *gin.Context, err error) {
	obj := JSONErrorResponse{
		Message:   err.Error(),
		Timestamp: time.Now().String(),
	}
	c.JSON(http.StatusBadRequest, obj)
}

func NotFoundResponse(c *gin.Context, err error) {
	obj := JSONErrorResponse{
		Message:   err.Error(),
		Timestamp: time.Now().String(),
	}
	c.JSON(http.StatusNotFound, obj)
}

func InternalServerErrorResponse(c *gin.Context, err error) {
	obj := JSONErrorResponse{
		Message:   err.Error(),
		Timestamp: time.Now().String(),
	}
	c.JSON(http.StatusInternalServerError, obj)
}

func ChooseErrorResponse(c *gin.Context, err error) {
	errorMsg := err.Error()
	switch {
	case strings.Contains(errorMsg, "application error"):
		InternalServerErrorResponse(c, err)
	case strings.Contains(errorMsg, "not found"):
		NotFoundResponse(c, err)
	case strings.Contains(errorMsg, "bad request") || strings.Contains(errorMsg, "invalid UUID"):
		BadRequestResponse(c, err)
	default:
		InternalServerErrorResponse(c, err)
	}
}
