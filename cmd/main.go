package main

import (
	_ "github.com/EnricoPDG/meli-desafio/docs"
	"github.com/EnricoPDG/meli-desafio/logger"
	"github.com/EnricoPDG/meli-desafio/router"
	"github.com/EnricoPDG/meli-desafio/service"
	"github.com/joho/godotenv"
	"go.uber.org/zap"
	"os"
)

// @title Meli Desafio
// @version 1.0
// @description Meli Desafio - Enrico Papsch Di Giacomo
// @host localhost:8080

var log = logger.GetLogger()

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Error("Error loading .env file", zap.Error(err))
	}

	productFilePath := os.Getenv("PRODUCT_FILE_PATH")
	reviewFilePath := os.Getenv("REVIEW_FILE_PATH")
	sellerFilePath := os.Getenv("SELLER_FILE_PATH")

	productService := service.NewProductService(productFilePath, reviewFilePath)
	sellerService := service.NewSellerService(sellerFilePath)

	r := router.SetupRouter(productService, sellerService)

	if err := r.Run(":8080"); err != nil {
		log.Fatal("Error starting server", zap.Error(err))
	}
}
