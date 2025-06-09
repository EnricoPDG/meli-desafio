package service_test

import (
	"github.com/EnricoPDG/meli-desafio/model"
	"github.com/EnricoPDG/meli-desafio/service"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestGetSellerByID(t *testing.T) {
	idFound := uuid.New()
	idNotFound := uuid.New()

	sellers := []*model.Seller{
		{ID: idFound, Name: "Test Seller"},
	}

	sellerFile := createTempJSONFile(t, sellers)
	defer os.Remove(sellerFile)

	tests := []struct {
		name          string
		sellerID      uuid.UUID
		filePath      string
		expectedName  string
		expectedError string
	}{
		{
			name:         "success - seller found",
			sellerID:     idFound,
			filePath:     sellerFile,
			expectedName: "Test Seller",
		},
		{
			name:          "fail - seller not found",
			sellerID:      idNotFound,
			filePath:      sellerFile,
			expectedError: "seller not found",
		},
		{
			name:          "fail - error loading JSON file",
			sellerID:      idFound,
			filePath:      "invalid/path.json",
			expectedError: "application error",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			service := service.NewSellerService(tt.filePath)

			seller, err := service.GetSellerByID(tt.sellerID)

			if tt.expectedError != "" {
				assert.Nil(t, seller)
				assert.EqualError(t, err, tt.expectedError)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.expectedName, seller.Name)
			}
		})
	}
}
