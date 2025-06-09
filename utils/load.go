package utils

import (
	"encoding/json"
	"fmt"
	"os"
	"sync"
)

var mu sync.RWMutex

// LoadJSON reads a JSON file from the given filePath and unmarshal its content into data.
//
// The data parameter must be a pointer to the variable where the JSON content will be stored.
// It uses a read lock to ensure concurrent safe access to the file.
//
// Returns an error if the file cannot be read or the JSON cannot be unmarshalled.
func LoadJSON(filePath string, data interface{}) error {
	mu.RLock()
	defer mu.RUnlock()

	file, err := os.ReadFile(filePath)
	if err != nil {
		return fmt.Errorf("error reading file: %s", filePath)
	}

	if err := json.Unmarshal(file, &data); err != nil {
		return fmt.Errorf("error unmarshalling data")
	}

	return nil
}
