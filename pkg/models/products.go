package models

import (
	"github.com/adeindra6/synapsis_backend_test/pkg/config"
	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	ProductName string `gorm:"type:text"json:"product_name"`
	Price       uint64 `gorm:"type:uint"json:"price"`
	Stock       uint64 `gorm:"type:uint"json:"stock"`
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
