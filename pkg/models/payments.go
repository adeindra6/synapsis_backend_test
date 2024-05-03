package models

import (
	"time"

	"github.com/adeindra6/synapsis_backend_test/pkg/config"
	"gorm.io/gorm"
)

type Payment struct {
	gorm.Model
	CartId        int64     `gorm:"type:int"json:"cart_id"`
	TotalPaid     uint64    `gorm:"type:uint"json:"total_paid"`
	PaymentMethod string    `gorm:"type:varchar(255)"json:"payment_method"`
	PaymentTime   time.Time `gorm:"type:TIMESTAMP;null;default:null"json:"payment_time"`
}

func init() {
	config.Connect()
	db = config.GetDB()
	db.AutoMigrate(&Payment{})
}

func (p *Payment) CreatePayment() *Payment {
	db.Create(&p)
	return p
}

func GetAllPayments() []Payment {
	var Payments []Payment
	db.Find(&Payments)
	return Payments
}

func GetPaymentsById(id int64) (*Payment, *gorm.DB) {
	var getPayment Payment
	db := db.Where("id = ?", id).Find(&getPayment)
	return &getPayment, db
}

func CheckIfCheckedOut(cartId int64) (count int64) {
	var c int64
	db.Where("cart_id = ?", cartId).Count(&c)
	return c
}

func DeletePaymentById(id int64) Payment {
	var payment Payment
	db.Where("id = ?", id).Delete(&payment)
	return payment
}
