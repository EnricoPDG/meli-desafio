package response_test

import (
	"errors"
	"github.com/EnricoPDG/meli-desafio/response"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestBadRequestResponse(t *testing.T) {
	gin.SetMode(gin.TestMode)

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	response.BadRequestResponse(c, errors.New("bad request"))

	assert.Equal(t, http.StatusBadRequest, w.Code)
}

func TestNotFoundResponse(t *testing.T) {
	gin.SetMode(gin.TestMode)

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	response.NotFoundResponse(c, errors.New("not found"))

	assert.Equal(t, http.StatusNotFound, w.Code)
}

func TestInternalServerErrorResponse(t *testing.T) {
	gin.SetMode(gin.TestMode)

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	response.InternalServerErrorResponse(c, errors.New("not found"))

	assert.Equal(t, http.StatusInternalServerError, w.Code)
}

func TestChooseErrorResponse(t *testing.T) {
	gin.SetMode(gin.TestMode)

	test := []struct {
		name         string
		expectedCode int
		errorMsg     string
	}{
		{
			name:         "bad request error",
			expectedCode: http.StatusBadRequest,
			errorMsg:     "bad request",
		},
		{
			name:         "internal server error",
			expectedCode: http.StatusInternalServerError,
			errorMsg:     "application error",
		},
		{
			name:         "not found error",
			expectedCode: http.StatusNotFound,
			errorMsg:     "not found",
		},
		{
			name:         "invalid UUID error",
			expectedCode: http.StatusBadRequest,
			errorMsg:     "bad request invalid UUID",
		},
		{
			name:         "generic error",
			expectedCode: http.StatusInternalServerError,
			errorMsg:     "unexpected error",
		},
	}

	for _, tt := range test {
		t.Run(tt.name, func(t *testing.T) {
			w := httptest.NewRecorder()
			c, _ := gin.CreateTestContext(w)

			response.ChooseErrorResponse(c, errors.New(tt.errorMsg))

			assert.Equal(t, tt.expectedCode, w.Code)
		})
	}
}
