package models

import (
	"github.com/adeindra6/synapsis_backend_test/pkg/config"
	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	ProductName string `gorm:"type:text"json:"product_name,omitempty"`
	Price       uint64 `gorm:"type:uint"json:"price,omitempty"`
	Stock       uint64 `gorm:"type:uint"json:"stock,omitempty"`
	Category    string `gorm:"type:varchar(255)"json:"category,omitempty"`
}

func init() {
	config.Connect()
	db = config.GetDB()
	db.AutoMigrate(&Product{})
}

func (p *Product) CreateProduct() *Product {
	db.Create(&p)
	return p
}

func GetAllProducts() []Product {
	var Products []Product
	db.Find(&Products)
	return Products
}

func GetProductListByCategory(category string) []Product {
	var Products []Product
	db.Where("category = ?", category).Find(&Products)
	return Products
}

func GetProductById(id int64) (*Product, *gorm.DB) {
	var getProduct Product
	db := db.Where("id = ?", id).Find(&getProduct)
	return &getProduct, db
}

func DeleteProductById(id int64) Product {
	var product Product
	db.Where("id = ?", id).Delete(&product)
	return product
}
