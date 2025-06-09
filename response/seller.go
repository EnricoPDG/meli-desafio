package response

import (
	"github.com/EnricoPDG/meli-desafio/model"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

type JSONSuccessResultSeller struct {
	Message   string        `json:"message" example:"success"`
	Data      *model.Seller `json:"data"`
	Timestamp string        `json:"timestamp" example:"2025-06-05T18:44:26.131446-03:00"`
}

func SuccessResponseSeller(c *gin.Context, Seller *model.Seller) {
	obj := JSONSuccessResultSeller{
		Message:   "success",
		Data:      Seller,
		Timestamp: time.Now().String(),
	}
	c.JSON(http.StatusOK, obj)
}
