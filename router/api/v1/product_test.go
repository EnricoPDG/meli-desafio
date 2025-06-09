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

func TestGetProductByID(t *testing.T) {
	validUUID := uuid.New()
	validReq, _ := http.NewRequest("GET", fmt.Sprintf("/api/v1/products/%s", validUUID), nil)
	validProductData := model.Product{
		ID:               validUUID,
		Title:            "test",
		Description:      "test",
		ShortDescription: "test",
		Price: model.Price{
			Amount:   1.00,
			Currency: "BRL",
		},
		Images: []model.Image{
			{URL: "test", Size: "test"},
		},
		PaymentMethods: []string{"pix"},
		Rating:         1.0,
		Stock:          1,
		SellerID:       validUUID,
		Condition:      "test",
		SoldQuantity:   1,
		CreateAt:       time.Now(),
		UpdateAt:       time.Now(),
	}
	validResponseData := response.JSONSuccessResultProduct{
		Message:   "success",
		Data:      validProductData,
		Timestamp: time.Now().String(),
	}

	invalidUUID := "123"
	badReqInvalidUUID, _ := http.NewRequest("GET", fmt.Sprintf("/api/v1/products/%s", invalidUUID), nil)

	tt := []struct {
		name           string
		req            *http.Request
		expectedStatus int
		responseData   interface{}
		productData    *model.Product
		err            error
	}{
		{
			name:           "valid request 200",
			req:            validReq,
			expectedStatus: http.StatusOK,
			responseData:   validResponseData,
			productData:    &validProductData,
			err:            nil,
		},
		{
			name:           "bad request invalid uuid 400",
			req:            badReqInvalidUUID,
			expectedStatus: http.StatusBadRequest,
			responseData:   nil,
			productData:    nil,
			err:            errors.New("bad request, invalid uuid"),
		},
		{
			name:           "product not found 404",
			req:            validReq,
			expectedStatus: http.StatusNotFound,
			responseData:   nil,
			productData:    nil,
			err:            errors.New("product not found"),
		},
		{
			name:           "application error 500",
			req:            validReq,
			expectedStatus: http.StatusInternalServerError,
			responseData:   nil,
			productData:    nil,
			err:            errors.New("application error"),
		},
	}

	sellerMock := &mocks.SellerServiceMock{}
	productMock := &mocks.ProductServiceMock{}

	r := router.SetupRouter(productMock, sellerMock)

	for _, test := range tt {
		t.Run(test.name, func(t *testing.T) {
			productMock.Data = test.productData
			productMock.Error = test.err

			w := httptest.NewRecorder()
			r.ServeHTTP(w, test.req)

			assert.Equal(t, test.expectedStatus, w.Code)
		})
	}
}

func TestGetReviewsByProductID(t *testing.T) {
	validUUID := uuid.New()
	validReq, _ := http.NewRequest("GET", fmt.Sprintf("/api/v1/products/%s/reviews", validUUID), nil)
	validReviewsData := []*model.Review{
		{
			ID:        validUUID,
			ProductID: validUUID,
			Rating:    1.0,
			Content:   "test",
			Author:    "test",
			CreateAt:  time.Now(),
			UpdateAt:  time.Now(),
		},
		{
			ID:        validUUID,
			ProductID: validUUID,
			Rating:    2.0,
			Content:   "test2",
			Author:    "test2",
			CreateAt:  time.Now(),
			UpdateAt:  time.Now(),
		},
	}
	validResponseData := response.JSONSuccessResultReview{
		Data:      validReviewsData,
		Message:   "success",
		Timestamp: time.Now().String(),
	}

	invalidUUID := "123"
	badReqInvalidUUID, _ := http.NewRequest("GET", fmt.Sprintf("/api/v1/products/%s/reviews", invalidUUID), nil)

	tt := []struct {
		name           string
		req            *http.Request
		expectedStatus int
		responseData   interface{}
		reviewsData    []*model.Review
		err            error
	}{
		{
			name:           "valid request 200",
			req:            validReq,
			expectedStatus: http.StatusOK,
			responseData:   validResponseData,
			reviewsData:    validReviewsData,
			err:            nil,
		},
		{
			name:           "bad request invalid uuid 400",
			req:            badReqInvalidUUID,
			expectedStatus: http.StatusBadRequest,
			responseData:   nil,
			reviewsData:    nil,
			err:            errors.New("bad request, invalid uuid"),
		},
		{
			name:           "product not found 404",
			req:            validReq,
			expectedStatus: http.StatusNotFound,
			responseData:   nil,
			reviewsData:    nil,
			err:            errors.New("product not found"),
		},
		{
			name:           "application error 500",
			req:            validReq,
			expectedStatus: http.StatusInternalServerError,
			responseData:   nil,
			reviewsData:    nil,
			err:            errors.New("application error"),
		},
	}

	sellerMock := &mocks.SellerServiceMock{}
	productMock := &mocks.ProductServiceMock{}

	r := router.SetupRouter(productMock, sellerMock)

	for _, test := range tt {
		t.Run(test.name, func(t *testing.T) {
			productMock.Data = test.reviewsData
			productMock.Error = test.err

			w := httptest.NewRecorder()
			r.ServeHTTP(w, test.req)

			assert.Equal(t, test.expectedStatus, w.Code)
		})
	}
}

func TestListProducts(t *testing.T) {
	validPage := 1
	validLimit := 10
	validReq, _ := http.NewRequest("GET", fmt.Sprintf("/api/v1/products?page=%d&limit=%d", validPage, validLimit), nil)
	validProductsData := []*model.Product{
		{
			ID:               uuid.New(),
			Title:            "test",
			Description:      "test",
			ShortDescription: "test",
			Price: model.Price{
				Amount:   1.00,
				Currency: "BRL",
			},
			Images: []model.Image{
				{URL: "test", Size: "test"},
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
				Amount:   1.00,
				Currency: "BRL",
			},
			Images: []model.Image{
				{URL: "test", Size: "test"},
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

	validResponseData := response.JSONSuccessResultProducts{
		Message:   "success",
		Data:      validProductsData,
		Timestamp: time.Now().String(),
	}

	invalidPage := "invalid"
	badReqInvalidQuery, _ := http.NewRequest("GET", fmt.Sprintf("/api/v1/products?page=%s&limit=%d", invalidPage, validLimit), nil)

	tt := []struct {
		name           string
		req            *http.Request
		expectedStatus int
		responseData   interface{}
		productsData   []*model.Product
		err            error
	}{
		{
			name:           "valid request 200",
			req:            validReq,
			expectedStatus: http.StatusOK,
			responseData:   validResponseData,
			productsData:   validProductsData,
			err:            nil,
		},
		{
			name:           "bad request invalid query parameter 400",
			req:            badReqInvalidQuery,
			expectedStatus: http.StatusBadRequest,
			responseData:   nil,
			productsData:   nil,
			err:            errors.New("bad request invalid query parameter"),
		},
		{
			name:           "product not found 404",
			req:            validReq,
			expectedStatus: http.StatusNotFound,
			responseData:   nil,
			productsData:   nil,
			err:            errors.New("products not found"),
		},
		{
			name:           "application error 500",
			req:            validReq,
			expectedStatus: http.StatusInternalServerError,
			responseData:   nil,
			productsData:   nil,
			err:            errors.New("application error"),
		},
	}

	sellerMock := &mocks.SellerServiceMock{}
	productMock := &mocks.ProductServiceMock{}

	r := router.SetupRouter(productMock, sellerMock)

	for _, test := range tt {
		t.Run(test.name, func(t *testing.T) {
			productMock.Data = test.productsData
			productMock.Error = test.err

			w := httptest.NewRecorder()
			r.ServeHTTP(w, test.req)

			assert.Equal(t, test.expectedStatus, w.Code)
		})
	}
}
