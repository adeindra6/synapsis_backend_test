package models

import (
	"net/http"
	"time"

	"github.com/adeindra6/synapsis_backend_test/pkg/config"
	jwt "github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

var db *gorm.DB

type Customer struct {
	gorm.Model
	Fullname string `gorm:"type:text"json:"fullname"`
	Email    string `gorm:"type:varchar(255)"json:"email"`
	Address  string `gorm:"type:text"json:"address"`
	Phone    string `gorm:"type:varchar(255)"json:"phone"`
	Password string `gorm:"type:text"json:"password"`
}

type LoginRes struct {
	Token  string    `gorm:""json:"token"`
	Status string    `json:"status"`
	Code   int64     `json:"code"`
	Exp    time.Time `json:"exp"`
}

func init() {
	config.Connect()
	db = config.GetDB()
	db.AutoMigrate(&Customer{})
}

func (c *Customer) CreateCustomer() *Customer {
	password := []byte(c.Password)
	hashPassword, err := bcrypt.GenerateFromPassword(password, bcrypt.DefaultCost)
	if err != nil {
		panic(err)
	}
	c.Password = string(hashPassword)

	db.Create(&c)
	return c
}

func GetAllCustomers() []Customer {
	var Customers []Customer
	db.Find(&Customers)
	return Customers
}

func GetCustomerById(id int64) (*Customer, *gorm.DB) {
	var getCustomer Customer
	db := db.Where("id = ?", id).Find(&getCustomer)
	return &getCustomer, db
}

func DeleteCustomerById(id int64) Customer {
	var customer Customer
	db.Where("id = ?", id).Delete(&customer)
	return customer
}

func (c *Customer) Login() *LoginRes {
	var LoginCustomer Customer
	var PasswordCorrect bool
	var Token LoginRes
	db.Where("email = ?", c.Email).Find(&LoginCustomer)

	password := []byte(c.Password)
	storedPassword := []byte(LoginCustomer.Password)
	PasswordCorrect = false

	err := bcrypt.CompareHashAndPassword(storedPassword, password)
	if err == nil {
		PasswordCorrect = true
	}

	if PasswordCorrect {
		expTime := time.Now().Add(time.Hour * 24)
		jwtToken, err := CreateJWTToken(LoginCustomer.ID, LoginCustomer.Email, expTime)
		if err != nil {
			panic(err)
		}

		Token.Token = jwtToken
		Token.Status = "SUCCESS"
		Token.Code = http.StatusOK
		Token.Exp = expTime
	}

	return &Token
}

func CreateJWTToken(id uint, email string, expTime time.Time) (string, error) {
	var err error

	atClaims := jwt.MapClaims{}
	atClaims["authorized"] = true
	atClaims["customer_id"] = id
	atClaims["username"] = email
	atClaims["exp"] = expTime.Unix()
	at := jwt.NewWithClaims(jwt.SigningMethodHS512, atClaims)

	token, err := at.SignedString([]byte("Synapsis_Test"))
	if err != nil {
		return "", err
	}

	return token, nil
}
