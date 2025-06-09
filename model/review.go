package model

import (
	"github.com/google/uuid"
	"time"
)

type Review struct {
	ID        uuid.UUID `uri:"id" json:"id" form:"id" example:"123e4567-e89b-12d3-a456-426655440000"`
	ProductID uuid.UUID `json:"product_id" form:"product_id"`
	Rating    float64   `json:"rating" form:"rating"`
	Content   string    `json:"content" form:"content"`
	Author    string    `json:"author" form:"author"`
	CreateAt  time.Time `json:"create_at" form:"create_at"`
	UpdateAt  time.Time `json:"update_at" form:"update_at"`
}
