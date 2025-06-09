package utils_test

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"testing"

	"github.com/EnricoPDG/meli-desafio/utils"
	"github.com/stretchr/testify/assert"
)

type MockData struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func TestLoadJSON(t *testing.T) {
	validData := MockData{Name: "John", Age: 30}
	validJSON, _ := json.Marshal(validData)

	invalidJSON := []byte(`{"name": "John", "age": }`)

	os.WriteFile("valid.json", validJSON, 0644)
	os.WriteFile("invalid.json", invalidJSON, 0644)
	defer os.Remove("valid.json")
	defer os.Remove("invalid.json")

	tests := []struct {
		name        string
		filePath    string
		expectError bool
		expectedErr error
		expected    *MockData
	}{
		{
			name:        "file does not exist",
			filePath:    "nonexistent.json",
			expectError: true,
			expectedErr: fmt.Errorf("error reading file: %s", "nonexistent.json"),
		},
		{
			name:        "invalid JSON content",
			filePath:    "invalid.json",
			expectError: true,
			expectedErr: errors.New("error unmarshalling data"),
		},
		{
			name:        "valid JSON content",
			filePath:    "valid.json",
			expectError: false,
			expected:    &validData,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var result MockData
			err := utils.LoadJSON(tt.filePath, &result)

			if tt.expectError {
				assert.Error(t, err)
				assert.Contains(t, err.Error(), tt.expectedErr.Error())
			} else {
				assert.NoError(t, err)
				assert.Equal(t, *tt.expected, result)
			}
		})
	}
}
