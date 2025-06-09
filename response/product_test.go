package response_test

import (
	"github.com/EnricoPDG/meli-desafio/model"
	"github.com/EnricoPDG/meli-desafio/response"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestSuccessResponseProduct(t *testing.T) {
	gin.SetMode(gin.TestMode)

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	productList := model.Product{
		ID:               uuid.New(),
		Title:            "test",
		Description:      "test",
		ShortDescription: "test",
		Price: model.Price{
			Amount:   1.0,
			Currency: "BRL",
		},
		Images: []model.Image{
			{
				URL:  "test",
				Size: "test",
			},
		},
		PaymentMethods: []string{"pix"},
		Rating:         1.0,
		Stock:          1,
		SellerID:       uuid.New(),
		Condition:      "test",
		SoldQuantity:   1,
		CreateAt:       time.Now(),
		UpdateAt:       time.Now(),
	}

	response.SuccessResponseProduct(c, productList)

	assert.Equal(t, http.StatusOK, w.Code)
}

func TestSuccessResponseProducts(t *testing.T) {
	gin.SetMode(gin.TestMode)

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	productList := []*model.Product{
		{
			ID:               uuid.New(),
			Title:            "test",
			Description:      "test",
			ShortDescription: "test",
			Price: model.Price{
				Amount:   1.0,
				Currency: "BRL",
			},
			Images: []model.Image{
				{
					URL:  "test",
					Size: "test",
				},
			},
			PaymentMethods: []string{"pix"},
			Rating:         1.0,
			Stock:          1,
			SellerID:       uuid.New(),
			Condition:      "test",
			SoldQuantity:   1,
			CreateAt:       time.Now(),
			UpdateAt:       time.Now(),
		},
		{
			ID:               uuid.New(),
			Title:            "test",
			Description:      "test",
			ShortDescription: "test",
			Price: model.Price{
				Amount:   1.0,
				Currency: "BRL",
			},
			Images: []model.Image{
				{
					URL:  "test",
					Size: "test",
				},
			},
			PaymentMethods: []string{"pix"},
			Rating:         1.0,
			Stock:          1,
			SellerID:       uuid.New(),
			Condition:      "test",
			SoldQuantity:   1,
			CreateAt:       time.Now(),
			UpdateAt:       time.Now(),
		},
		{
			ID:               uuid.New(),
			Title:            "test",
			Description:      "test",
			ShortDescription: "test",
			Price: model.Price{
				Amount:   1.0,
				Currency: "BRL",
			},
			Images: []model.Image{
				{
					URL:  "test",
					Size: "test",
				},
			},
			PaymentMethods: []string{"pix"},
			Rating:         1.0,
			Stock:          1,
			SellerID:       uuid.New(),
			Condition:      "test",
			SoldQuantity:   1,
			CreateAt:       time.Now(),
			UpdateAt:       time.Now(),
		},
		{
			ID:               uuid.New(),
			Title:            "test",
			Description:      "test",
			ShortDescription: "test",
			Price: model.Price{
				Amount:   1.0,
				Currency: "BRL",
			},
			Images: []model.Image{
				{
					URL:  "test",
					Size: "test",
				},
			},
			PaymentMethods: []string{"pix"},
			Rating:         1.0,
			Stock:          1,
			SellerID:       uuid.New(),
			Condition:      "test",
			SoldQuantity:   1,
			CreateAt:       time.Now(),
			UpdateAt:       time.Now(),
		},
		{
			ID:               uuid.New(),
			Title:            "test",
			Description:      "test",
			ShortDescription: "test",
			Price: model.Price{
				Amount:   1.0,
				Currency: "BRL",
			},
			Images: []model.Image{
				{
					URL:  "test",
					Size: "test",
				},
			},
			PaymentMethods: []string{"pix"},
			Rating:         1.0,
			Stock:          1,
			SellerID:       uuid.New(),
			Condition:      "test",
			SoldQuantity:   1,
			CreateAt:       time.Now(),
			UpdateAt:       time.Now(),
		},
	}

	response.SuccessResponseProducts(c, productList)

	assert.Equal(t, http.StatusOK, w.Code)
}
