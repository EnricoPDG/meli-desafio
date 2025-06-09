package model

import (
	"github.com/google/uuid"
	"time"
)

type Seller struct {
	ID         uuid.UUID  `uri:"id" json:"id" form:"id" example:"123e4567-e89b-12d3-a456-426655440000"`
	Nickname   string     `json:"nickname" form:"nickname"`
	Name       string     `json:"name" form:"title"`
	Email      string     `json:"email" form:"email"`
	Phone      string     `json:"phone" form:"phone"`
	Reputation Reputation `json:"reputation" form:"reputation"`
	Address    Address    `json:"address" form:"address"`
	CreateAt   time.Time  `json:"create_at" form:"create_at"`
	UpdateAt   time.Time  `json:"update_at" form:"update_at"`
}

type Address struct {
	City    string `json:"city" form:"city"`
	State   string `json:"state" form:"state"`
	Country string `json:"country" form:"country"`
}

type Reputation struct {
	Level       string      `json:"level" form:"level"`
	Transaction Transaction `json:"transaction" form:"transaction"`
}

type Transaction struct {
	Completed int     `json:"completed" form:"completed"`
	Canceled  int     `json:"canceled" form:"canceled"`
	Rating    float64 `json:"rating" form:"rating"`
}
