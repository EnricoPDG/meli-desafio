package v1_test

import (
	"errors"
	"fmt"
	"github.com/EnricoPDG/meli-desafio/mocks"
	"github.com/EnricoPDG/meli-desafio/model"
	"github.com/EnricoPDG/meli-desafio/response"
	"github.com/EnricoPDG/meli-desafio/router"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestGetSellerByID(t *testing.T) {
	validUUID := uuid.New()
	validReq, _ := http.NewRequest("GET", fmt.Sprintf("/api/v1/sellers/%s", validUUID), nil)
	validSeller := model.Seller{
		ID:       validUUID,
		Nickname: "test",
		Name:     "test",
		Email:    "test@test.com",
		Phone:    "013 12345-6789",
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
		CreateAt: time.Time{},
		UpdateAt: time.Time{},
	}
	validResponseData := response.JSONSuccessResultSeller{
		Message:   "success",
		Data:      &validSeller,
		Timestamp: time.Now().String(),
	}

	invalidUUID := "123"
	badReqInvalidUUID, _ := http.NewRequest("GET", fmt.Sprintf("/api/v1/sellers/%s", invalidUUID), nil)

	tt := []struct {
		name           string
		req            *http.Request
		expectedStatus int
		responseData   interface{}
		sellerData     *model.Seller
		err            error
	}{
		{
			name:           "valid request 200",
			req:            validReq,
			expectedStatus: http.StatusOK,
			responseData:   validResponseData,
			sellerData:     &validSeller,
			err:            nil,
		},
		{
			name:           "bad request invalid uuid 400",
			req:            badReqInvalidUUID,
			expectedStatus: http.StatusBadRequest,
			responseData:   nil,
			sellerData:     nil,
			err:            errors.New("bad request, invalid uuid"),
		},
		{
			name:           "product not found 404",
			req:            validReq,
			expectedStatus: http.StatusNotFound,
			responseData:   nil,
			sellerData:     nil,
			err:            errors.New("product not found"),
		},
		{
			name:           "application error 500",
			req:            validReq,
			expectedStatus: http.StatusInternalServerError,
			responseData:   nil,
			sellerData:     nil,
			err:            errors.New("application error"),
		},
	}

	sellerMock := &mocks.SellerServiceMock{}
	productMock := &mocks.ProductServiceMock{}

	r := router.SetupRouter(productMock, sellerMock)

	for _, test := range tt {
		t.Run(test.name, func(t *testing.T) {
			sellerMock.Data = test.sellerData
			sellerMock.Error = test.err

			w := httptest.NewRecorder()
			r.ServeHTTP(w, test.req)

			assert.Equal(t, test.expectedStatus, w.Code)
		})
	}
}
