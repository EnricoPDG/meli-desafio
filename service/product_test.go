package service_test

import (
	"encoding/json"
	"github.com/EnricoPDG/meli-desafio/model"
	"github.com/EnricoPDG/meli-desafio/service"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
	"time"
)

func createTempJSONFile(t *testing.T, structData interface{}) string {
	data, err := json.Marshal(structData)
	assert.NoError(t, err)

	tmpfile, err := os.CreateTemp("", "temp-*.json")
	assert.NoError(t, err)

	_, err = tmpfile.Write(data)
	assert.NoError(t, err)

	tmpfile.Close()
	return tmpfile.Name()
}

func TestGetProductByID(t *testing.T) {
	existingID := uuid.New()
	productList := []*model.Product{
		{
			ID:               existingID,
			Title:            "test",
			Description:      "test",
			ShortDescription: "test",
			Price: model.Price{
				Amount:   1,
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

	validFile := createTempJSONFile(t, productList)
	defer os.Remove(validFile)

	emptyFile := createTempJSONFile(t, []*model.Product{})
	defer os.Remove(emptyFile)

	tests := []struct {
		name          string
		filePath      string
		searchID      uuid.UUID
		expectedError string
		expectedTitle string
	}{
		{
			name:          "success - product found",
			filePath:      validFile,
			searchID:      existingID,
			expectedError: "",
			expectedTitle: "test",
		},
		{
			name:          "fail - product not found",
			filePath:      validFile,
			searchID:      uuid.New(),
			expectedError: "product not found",
		},
		{
			name:          "fail - error loading JSON file",
			filePath:      "invalid/file/path.json",
			searchID:      existingID,
			expectedError: "application error",
		},
		{
			name:          "fail - empty product list",
			filePath:      emptyFile,
			searchID:      uuid.New(),
			expectedError: "product not found",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			service := service.NewProductService(tt.filePath, "")

			product, err := service.GetProductByID(tt.searchID)

			if tt.expectedError != "" {
				assert.Nil(t, product)
				assert.EqualError(t, err, tt.expectedError)
			} else {
				assert.NoError(t, err)
				assert.NotNil(t, product)
				assert.Equal(t, tt.expectedTitle, product.Title)
			}
		})
	}
}

func TestGetReviewsByProductID(t *testing.T) {
	productID := uuid.New()

	reviewList := []*model.Review{
		{
			ID:        uuid.New(),
			ProductID: productID,
			Rating:    1.0,
			Content:   "test",
			Author:    "test",
			CreateAt:  time.Now(),
			UpdateAt:  time.Now(),
		},
		{
			ID:        uuid.New(),
			ProductID: uuid.New(),
			Rating:    2.0,
			Content:   "test",
			Author:    "test",
			CreateAt:  time.Now(),
			UpdateAt:  time.Now(),
		},
	}

	reviewsFileWithMatch := createTempJSONFile(t, reviewList)
	defer os.Remove(reviewsFileWithMatch)

	reviewsFileEmpty := createTempJSONFile(t, []*model.Review{})
	defer os.Remove(reviewsFileEmpty)

	tests := []struct {
		name             string
		filePath         string
		productID        uuid.UUID
		expectedError    string
		expectedReviewCt int
	}{
		{
			name:             "success - reviews found",
			filePath:         reviewsFileWithMatch,
			productID:        productID,
			expectedError:    "",
			expectedReviewCt: 1,
		},
		{
			name:             "fail - no reviews for product",
			filePath:         reviewsFileWithMatch,
			productID:        uuid.New(),
			expectedError:    "not found any review",
			expectedReviewCt: 0,
		},
		{
			name:             "fail - error loading JSON file",
			filePath:         "invalid/path.json",
			productID:        productID,
			expectedError:    "application error",
			expectedReviewCt: 0,
		},
		{
			name:             "fail - empty review list",
			filePath:         reviewsFileEmpty,
			productID:        productID,
			expectedError:    "not found any review",
			expectedReviewCt: 0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			service := service.NewProductService("", tt.filePath)

			reviews, err := service.GetReviewsByProductID(tt.productID)

			if tt.expectedError != "" {
				assert.Nil(t, reviews)
				assert.EqualError(t, err, tt.expectedError)
			} else {
				assert.NoError(t, err)
				assert.Len(t, reviews, tt.expectedReviewCt)
			}
		})
	}
}

func TestListProducts(t *testing.T) {
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

	productFile := createTempJSONFile(t, productList)
	defer os.Remove(productFile)

	tests := []struct {
		name          string
		page          int
		limit         int
		expectedCount int
		expectedError string
		filePath      string
	}{
		{
			name:          "success - first page full",
			page:          1,
			limit:         2,
			expectedCount: 2,
			filePath:      productFile,
		},
		{
			name:          "success - second page partial",
			page:          2,
			limit:         3,
			expectedCount: 2,
			filePath:      productFile,
		},
		{
			name:          "fail - page out of range",
			page:          10,
			limit:         2,
			expectedError: "bad request, page out of range",
			filePath:      productFile,
		},
		{
			name:          "fail - error loading JSON",
			page:          1,
			limit:         2,
			expectedError: "application error",
			filePath:      "invalid/path.json",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			service := service.NewProductService(tt.filePath, "")

			products, err := service.ListProducts(tt.page, tt.limit)

			if tt.expectedError != "" {
				assert.Nil(t, products)
				assert.EqualError(t, err, tt.expectedError)
			} else {
				assert.NoError(t, err)
				assert.Len(t, products, tt.expectedCount)
			}
		})
	}
}
