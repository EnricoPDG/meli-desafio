package logger

import (
	"go.uber.org/zap"
	"sync"
)

var (
	log  *zap.Logger
	once sync.Once
)

// GetLogger returns a singleton instance of a zap.Logger.
//
// Later calls return the same logger instance.
func GetLogger() *zap.Logger {
	once.Do(func() {
		var err error
		log, err = zap.NewProduction()
		if err != nil {
			panic(err)
		}
	})

	return log
}
