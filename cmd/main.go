package main

import (
	_ "github.com/EnricoPDG/meli-desafio/docs"
	"github.com/EnricoPDG/meli-desafio/logger"
	"github.com/EnricoPDG/meli-desafio/router"
	"github.com/EnricoPDG/meli-desafio/service"
	"go.uber.org/zap"
)

// @title Meli Desafio
// @version 1.0
// @description Meli Desafio - Enrico Papsch Di Giacomo
// @host localhost:8080

var log = logger.GetLogger()

func main() {
	productFilePath := "data/products.json"
	reviewFilePath := "data/reviews.json"
	sellerFilePath := "data/sellers.json"

	productService := service.NewProductService(productFilePath, reviewFilePath)
	sellerService := service.NewSellerService(sellerFilePath)

	r := router.SetupRouter(productService, sellerService)

	if err := r.Run(":8080"); err != nil {
		log.Fatal("Error starting server", zap.Error(err))
	}
}
