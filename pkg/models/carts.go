package models

import (
	"github.com/adeindra6/synapsis_backend_test/pkg/config"
	"gorm.io/gorm"
)

type Cart struct {
	gorm.Model
	CustomerId int64    `gorm:"type:int"json:"customer_id,omitempty"`
	ProductId  int64    `gorm:"type:int"json:"product_id,omitempty"`
	Quantity   uint64   `gorm:"type:uint"json:"quantity,omitempty"`
	Amount     uint64   `gorm:"type:uint"json:"amount,omitempty"`
	IsPaid     bool     `gorm:"type:bool"json:"is_paid,omitempty"`
	Customer   Customer `gorm:"foreignkey:customer_id;references:id"`
	Product    Product  `gorm:"foreignkey:product_id;references:id"`
}

func init() {
	config.Connect()
	db = config.GetDB()
	db.AutoMigrate(&Cart{})
}

func (c *Cart) AddToCart() *Cart {
	db.Create(c)
	return c
}

func GetAllCarts() []Cart {
	var Carts []Cart
	db.Find(&Carts)
	return Carts
}

func GetCartById(id int64) (*Cart, *gorm.DB) {
	var getCart Cart
	db := db.Where("id = ?", id).Find(&getCart)
	return &getCart, db
}

func GetCartByCustomerId(customerId int64, isPaid bool) []Cart {
	var Carts []Cart
	db.Where("customer_id = ? AND is_paid = ?", customerId, isPaid).Find(&Carts)
	return Carts
}

func DeleteCartById(id int64) Cart {
	var cart Cart
	db.Where("id = ?", id).Delete(&cart)
	return cart
}
