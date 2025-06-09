package main

import (
	"encoding/json"
	"fmt"
	"github.com/EnricoPDG/meli-desafio/model"
	"os"
	"time"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/google/uuid"
)

func generateFakeSeller() model.Seller {
	now := time.Now()
	return model.Seller{
		ID:       uuid.New(),
		Nickname: gofakeit.Name(),
		Name:     gofakeit.Name(),
		Email:    gofakeit.Email(),
		Phone:    gofakeit.Phone(),
		Reputation: model.Reputation{
			Level: gofakeit.RandomString([]string{"platinum", "gold", "silver", "bronze"}),
			Transaction: model.Transaction{
				Completed: gofakeit.Number(0, 1000),
				Canceled:  gofakeit.Number(0, 1000),
				Rating:    gofakeit.Float64Range(1.0, 5.0),
			},
		},
		Address: model.Address{
			City:    gofakeit.City(),
			State:   gofakeit.State(),
			Country: gofakeit.Country(),
		},
		CreateAt: now,
		UpdateAt: now,
	}
}

func generateFakeProduct(sellerID uuid.UUID) model.Product {
	now := time.Now()
	return model.Product{
		ID:               uuid.New(),
		Title:            gofakeit.RandomString([]string{"iphone", "tv", "macbook", "laptot", "refrigerator"}),
		Description:      gofakeit.Paragraph(1, 2, 10, " "),
		ShortDescription: gofakeit.Sentence(5),
		Price: model.Price{
			Amount:   gofakeit.Price(10, 5000),
			Currency: "BRL",
		},
		Images: []model.Image{
			{URL: gofakeit.ImageURL(640, 480), Size: "medium"},
			{URL: gofakeit.ImageURL(300, 300), Size: "small"},
		},
		PaymentMethods: []string{"credit", "debit", "pix"},
		Rating:         gofakeit.Float64Range(1.0, 5.0),
		Stock:          gofakeit.Number(0, 200),
		SellerID:       sellerID,
		Condition:      gofakeit.RandomString([]string{"new", "used", "refurbished"}),
		SoldQuantity:   gofakeit.Number(0, 1000),
		CreateAt:       now.Add(-time.Duration(gofakeit.Number(1, 365)) * 24 * time.Hour),
		UpdateAt:       now,
	}
}

func generateFakeReview(productID uuid.UUID) model.Review {
	now := time.Now()
	return model.Review{
		ID:        uuid.New(),
		ProductID: productID,
		Rating:    gofakeit.Float64Range(1.0, 5.0),
		Content:   gofakeit.Sentence(10),
		Author:    gofakeit.Name(),
		CreateAt:  now.Add(-time.Duration(gofakeit.Number(1, 90)) * 24 * time.Hour),
		UpdateAt:  now,
	}
}

func saveJSON(filename string, data interface{}) {
	file, err := os.Create(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	enc := json.NewEncoder(file)
	enc.SetIndent("", " ")
	if err := enc.Encode(data); err != nil {
		panic(err)
	}
}

func main() {
	gofakeit.Seed(0)

	var products []model.Product
	var reviews []model.Review
	var sellers []model.Seller

	for s := 0; s < 20; s++ {
		seller := generateFakeSeller()
		sellers = append(sellers, seller)
		// Gera 1000 produtos por vendedor
		for i := 0; i < 1000; i++ {
			product := generateFakeProduct(seller.ID)
			products = append(products, product)

			// Gera entre 1 e 5 reviews por produto
			for j := 0; j < gofakeit.Number(1, 5); j++ {
				reviews = append(reviews, generateFakeReview(product.ID))
			}
		}
	}

	saveJSON("data/products.json", products)
	saveJSON("data/reviews.json", reviews)
	saveJSON("data/sellers.json", sellers)

	fmt.Println("Arquivos gerados com sucesso")
}
