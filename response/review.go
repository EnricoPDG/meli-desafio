package response

import (
	"github.com/EnricoPDG/meli-desafio/model"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

type JSONSuccessResultReview struct {
	Message   string          `json:"message" example:"success"`
	Data      []*model.Review `json:"data"`
	Timestamp string          `json:"timestamp" example:"2025-06-05T18:44:26.131446-03:00"`
}

func SuccessResponseReviews(c *gin.Context, Reviews []*model.Review) {
	obj := JSONSuccessResultReview{
		Message:   "success",
		Data:      Reviews,
		Timestamp: time.Now().String(),
	}
	c.JSON(http.StatusOK, obj)
}
