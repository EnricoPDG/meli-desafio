package model

import (
	"github.com/google/uuid"
	"time"
)

type Product struct {
	ID               uuid.UUID `uri:"id" json:"id" form:"id" example:"123e4567-e89b-12d3-a456-426655440000"`
	Title            string    `json:"title" form:"title"`
	Description      string    `json:"description" form:"description"`
	ShortDescription string    `json:"short_description" form:"short_description"`
	Price            Price     `json:"price" form:"price"`
	Images           []Image   `json:"images" form:"images"`
	PaymentMethods   []string  `json:"payment_methods" form:"payment_methods"`
	Rating           float64   `json:"rating" form:"rating"`
	Stock            int       `json:"stock" form:"stock"`
	SellerID         uuid.UUID `json:"seller_id" form:"seller_id"`
	Condition        string    `json:"condition" form:"condition"`
	SoldQuantity     int       `json:"sold_quantity" form:"sold_quantity"`
	CreateAt         time.Time `json:"create_at" form:"create_at"`
	UpdateAt         time.Time `json:"update_at" form:"update_at"`
}

type Price struct {
	Amount   float64 `json:"amount" form:"amount"`
	Currency string  `json:"currency" form:"currency"`
}

type Image struct {
	URL  string `json:"url" form:"url"`
	Size string `json:"size" form:"size"`
}

type ListProductQueryParameters struct {
	Page  int `json:"page" form:"page,default=1" binding:"required,gt=0"`
	Limit int `json:"limit" form:"limit,default=10" binding:"required,gt=0,lte=100"`
}

type ListProductsResponse struct {
	Products        []*Product `json:"products" form:"products"`
	QueryParameters ListProductQueryParameters
}
