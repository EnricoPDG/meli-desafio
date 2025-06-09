package response

import (
	"github.com/EnricoPDG/meli-desafio/model"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

type JSONSuccessResultProduct struct {
	Message   string        `json:"message" example:"success"`
	Data      model.Product `json:"data"`
	Timestamp string        `json:"timestamp" example:"2025-06-05T18:44:26.131446-03:00"`
}

type JSONSuccessResultProducts struct {
	Message          string           `json:"message" example:"success"`
	Data             []*model.Product `json:"data"`
	Timestamp        string           `json:"timestamp" example:"2025-06-05T18:44:26.131446-03:00"`
	NumberOfProducts int              `json:"number_of_products" example:"10"`
}

func SuccessResponseProduct(c *gin.Context, product model.Product) {
	obj := JSONSuccessResultProduct{
		Message:   "success",
		Data:      product,
		Timestamp: time.Now().String(),
	}
	c.JSON(http.StatusOK, obj)
}

func SuccessResponseProducts(c *gin.Context, product []*model.Product) {
	obj := JSONSuccessResultProducts{
		Message:          "success",
		NumberOfProducts: len(product),
		Data:             product,
		Timestamp:        time.Now().String(),
	}
	c.JSON(http.StatusOK, obj)
}
