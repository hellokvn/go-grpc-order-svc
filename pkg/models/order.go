package models

import "gorm.io/gorm"

type Order struct {
	gorm.Model
	Id        int32 `json:"id" gorm:"primaryKey"`
	Price     int32 `json:"price"`
	ProductId int32 `json:"product_id"`
	UserId    int32 `json:"user_id"`
}
