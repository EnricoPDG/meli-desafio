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

func TestSuccessResponseSeller(t *testing.T) {
	gin.SetMode(gin.TestMode)

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	seller := model.Seller{
		ID:       uuid.New(),
		Nickname: "test",
		Name:     "test",
		Email:    "test",
		Phone:    "test",
		Reputation: model.Reputation{
			Level: "test",
			Transaction: model.Transaction{
				Completed: 1,
				Canceled:  1,
				Rating:    1.0,
			},
		},
		Address: model.Address{
			City:    "test",
			State:   "test",
			Country: "test",
		},
		CreateAt: time.Now(),
		UpdateAt: time.Now(),
	}

	response.SuccessResponseSeller(c, &seller)

	assert.Equal(t, http.StatusOK, w.Code)
}
