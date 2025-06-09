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

func TestSuccessResponseReviews(t *testing.T) {
	gin.SetMode(gin.TestMode)

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	reviews := []*model.Review{
		{
			ID:        uuid.New(),
			ProductID: uuid.New(),
			Rating:    1.0,
			Content:   "test",
			Author:    "test",
			CreateAt:  time.Now(),
			UpdateAt:  time.Now(),
		},
	}

	response.SuccessResponseReviews(c, reviews)

	assert.Equal(t, http.StatusOK, w.Code)
}
