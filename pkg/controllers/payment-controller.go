package controllers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/adeindra6/synapsis_backend_test/pkg/models"
	"github.com/adeindra6/synapsis_backend_test/pkg/utils"
	"github.com/dgrijalva/jwt-go"
	"github.com/gorilla/mux"
)

var NewPayment models.Payment

type CartPayment struct {
	ID   int64       `json:"id"`
	Cart models.Cart `json:"cart"`

func CreatePayment(w http.ResponseWriter, r *http.Request) {
	var msg NonAuthorizedMsg
	var authorized bool
	var customer_id int64

	reqToken := r.Header.Get("Authorization")
	authorized = false
	customer_id = 0
	if strings.Contains(reqToken, "Bearer ") {
		splitToken := strings.Split(reqToken, "Bearer ")
		reqToken = splitToken[1]

		claims := jwt.MapClaims{}
		_, err := jwt.ParseWithClaims(reqToken, claims, func(token *jwt.Token) (interface{}, error) {
			return []byte(jwtKey), nil
		})

		if err != nil {
			panic(err)
		}

		for key, val := range claims {
			if key == "authorized" && val == true {
				authorized = true
			}
			if key == "customer_id" {
				customer_id = int64(val.(float64))
			}
		}
	} else {
		msg.Status = "UNAUTHORIZED"
		msg.Message = "Required Token!"
		msg.Code = http.StatusUnauthorized
	}

	if authorized && customer_id > 0 {
		CreatePayment := &models.Payment{}
		utils.ParseBody(r, CreatePayment)

		cart, _ := models.GetCartById(CreatePayment.CartId)
		count := models.CheckIfCheckedOut(CreatePayment.CartId)
		if cart.ID > 0 && count == 0 {
			p := CreatePayment.CreatePayment()
			res, err := json.Marshal(p)
			if err != nil {
				fmt.Println("Error while creating payment")
				err_msg := ErrMessage{
					Status:  "ERROR",
					Message: "Error while creating payment",
					Code:    http.StatusInternalServerError,
				}
				w.WriteHeader(http.StatusInternalServerError)
				json.NewEncoder(w).Encode(err_msg)
			}

			w.WriteHeader(http.StatusOK)
			w.Write(res)
		} else {
			w.Header().Set("Content-Type", "pkglication/json")
			w.WriteHeader(http.StatusNotFound)
			resBodyBytes := new(bytes.Buffer)
			msg.Code = 400
			msg.Message = "Cart Not Found or Already Checked Out"
			msg.Status = "BAD REQUEST"
			json.NewEncoder(resBodyBytes).Encode(msg)
			w.Write(resBodyBytes.Bytes())
		}
	} else {
		w.Header().Set("Content-Type", "pkglication/json")
		w.WriteHeader(http.StatusBadRequest)
		resBodyBytes := new(bytes.Buffer)
		json.NewEncoder(resBodyBytes).Encode(msg)
		w.Write(resBodyBytes.Bytes())
	}
}

func GetPaymentById(w http.ResponseWriter, r *http.Request) {
	var msg NonAuthorizedMsg
	var authorized bool
	var customer_id int64

	reqToken := r.Header.Get("Authorization")
	authorized = false
	customer_id = 0
	if strings.Contains(reqToken, "Bearer ") {
		splitToken := strings.Split(reqToken, "Bearer ")
		reqToken = splitToken[1]

		claims := jwt.MapClaims{}
		_, err := jwt.ParseWithClaims(reqToken, claims, func(token *jwt.Token) (interface{}, error) {
			return []byte(jwtKey), nil
		})

		if err != nil {
			panic(err)
		}

		for key, val := range claims {
			if key == "authorized" && val == true {
				authorized = true
			}
			if key == "customer_id" {
				customer_id = int64(val.(float64))
			}
		}
	} else {
		msg.Status = "UNAUTHORIZED"
		msg.Message = "Required Token!"
		msg.Code = http.StatusUnauthorized
	}

	if authorized && customer_id > 0 {
		vars := mux.Vars(r)
		paymentId := vars["paymentId"]
		id, err := strconv.ParseInt(paymentId, 0, 0)
		if err != nil {
			fmt.Println("Error while parsing")
		}

		var cartPayment CartPayment
		PaymentDetails, _ := models.GetPaymentsById(id)
		cartDetails, _ := models.GetCartById(PaymentDetails.CartId)
		customerDetails, _ := models.GetCustomerById(cartDetails.CustomerId)
		productDetails, _ := models.GetProductById(cartDetails.ProductId)

		cartPayment.ID = id
		cartPayment.Cart.ID = cartDetails.ID
		cartPayment.Cart.Quantity = cartDetails.Quantity
		cartPayment.Cart.Amount = cartDetails.Amount
		cartPayment.Cart.Customer.ID = customerDetails.ID
		cartPayment.Cart.Customer.Fullname = customerDetails.Fullname
		cartPayment.Cart.Customer.Email = customerDetails.Email
		cartPayment.Cart.Customer.Address = customerDetails.Address
		cartPayment.Cart.Customer.Phone = customerDetails.Phone
		cartPayment.Cart.Product.ID = productDetails.ID
		cartPayment.Cart.Product.ProductName = productDetails.ProductName
		cartPayment.Cart.Product.Price = productDetails.Price
		cartPayment.Cart.Product.Category = productDetails.Category

		res, err := json.Marshal(cartPayment)
		if err != nil {
			fmt.Println("Error when fetching payment")
			err_msg := ErrMessage{
				Status:  "ERROR",
				Message: "Error when fetching payment",
				Code:    http.StatusInternalServerError,
			}
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(err_msg)
		}

		w.Header().Set("Content-Type", "pkglication/json")
		w.WriteHeader(http.StatusOK)
		w.Write(res)
	} else {
		w.Header().Set("Content-Type", "pkglication/json")
		w.WriteHeader(http.StatusBadRequest)
		resBodyBytes := new(bytes.Buffer)
		json.NewEncoder(resBodyBytes).Encode(msg)
		w.Write(resBodyBytes.Bytes())
	}
}

func CheckoutPayment(w http.ResponseWriter, r *http.Request) {
	var msg NonAuthorizedMsg
	var authorized bool
	var customer_id int64

	reqToken := r.Header.Get("Authorization")
	authorized = false
	customer_id = 0
	if strings.Contains(reqToken, "Bearer ") {
		splitToken := strings.Split(reqToken, "Bearer ")
		reqToken = splitToken[1]

		claims := jwt.MapClaims{}
		_, err := jwt.ParseWithClaims(reqToken, claims, func(token *jwt.Token) (interface{}, error) {
			return []byte(jwtKey), nil
		})

		if err != nil {
			panic(err)
		}

		for key, val := range claims {
			if key == "authorized" && val == true {
				authorized = true
			}
			if key == "customer_id" {
				customer_id = int64(val.(float64))
			}
		}
	} else {
		msg.Status = "UNAUTHORIZED"
		msg.Message = "Required Token!"
		msg.Code = http.StatusUnauthorized
	}

	if authorized && customer_id > 0 {
		var checkoutPayment = &models.Payment{}
		utils.ParseBody(r, checkoutPayment)
		vars := mux.Vars(r)
		paymentId := vars["paymentId"]
		id, err := strconv.ParseInt(paymentId, 0, 0)
		if err != nil {
			fmt.Println("Error when updating product")
			err_msg := ErrMessage{
				Status:  "ERROR",
				Message: "Error when updating product",
				Code:    http.StatusInternalServerError,
			}
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(err_msg)
		}

		var paid bool
		paid = false
		paymentDetails, db := models.GetPaymentsById(id)
		if checkoutPayment.TotalPaid > 0 {
			paymentDetails.TotalPaid = checkoutPayment.TotalPaid
		}
		if checkoutPayment.PaymentMethod != "" {
			paymentDetails.PaymentMethod = checkoutPayment.PaymentMethod
		}
		if checkoutPayment.TotalPaid > 0 && checkoutPayment.PaymentMethod != "" {
			paid = true
		}

		var resMsg NonAuthorizedMsg

		if paid {
			paymentDetails.PaymentTime = time.Now()
			db.Save(&paymentDetails)

			paidCart, db2 := models.GetCartById(paymentDetails.CartId)
			paidCart.IsPaid = true
			db2.Save(&paidCart)

			resMsg.Code = http.StatusOK
			resMsg.Message = "Payment Successful"
			resMsg.Status = "SUCCESS"
		} else {
			resMsg.Code = http.StatusBadRequest
			resMsg.Message = "Payment Failed"
			resMsg.Status = "FAILED"
		}

		res, err := json.Marshal(resMsg)
		if err != nil {
			fmt.Println("Error while parsing!!!")
		}

		w.Header().Set("Content-Type", "pkglication/json")
		w.WriteHeader(http.StatusOK)
		w.Write(res)
	} else {
		w.Header().Set("Content-Type", "pkglication/json")
		w.WriteHeader(http.StatusBadRequest)
		resBodyBytes := new(bytes.Buffer)
		json.NewEncoder(resBodyBytes).Encode(msg)
		w.Write(resBodyBytes.Bytes())
	}
}

func DeletePayment(w http.ResponseWriter, r *http.Request) {
	var msg NonAuthorizedMsg
	var authorized bool
	var customer_id int64

	reqToken := r.Header.Get("Authorization")
	authorized = false
	customer_id = 0
	if strings.Contains(reqToken, "Bearer ") {
		splitToken := strings.Split(reqToken, "Bearer ")
		reqToken = splitToken[1]

		claims := jwt.MapClaims{}
		_, err := jwt.ParseWithClaims(reqToken, claims, func(token *jwt.Token) (interface{}, error) {
			return []byte(jwtKey), nil
		})

		if err != nil {
			panic(err)
		}

		for key, val := range claims {
			if key == "authorized" && val == true {
				authorized = true
			}
			if key == "customer_id" {
				customer_id = int64(val.(float64))
			}
		}
	} else {
		msg.Status = "UNAUTHORIZED"
		msg.Message = "Required Token!"
		msg.Code = http.StatusUnauthorized
	}

	if authorized && customer_id > 0 {
		vars := mux.Vars(r)
		paymentId := vars["paymentId"]
		id, err := strconv.ParseInt(paymentId, 0, 0)
		if err != nil {
			fmt.Println("Error when deleting customer's cart")
			err_msg := ErrMessage{
				Status:  "ERROR",
				Message: "Error when deleting customer's cart",
				Code:    http.StatusInternalServerError,
			}
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(err_msg)
		}

		_ = models.DeletePaymentById(id)
		var resMsg NonAuthorizedMsg
		resMsg.Code = http.StatusOK
		resMsg.Message = "Success Delete Payment"
		resMsg.Status = "SUCCESS"

		res, err := json.Marshal(resMsg)
		if err != nil {
			fmt.Println("Error while parsing!!!")
		}

		w.Header().Set("Content-Type", "pkglication/json")
		w.WriteHeader(http.StatusOK)
		w.Write(res)
	} else {
		w.Header().Set("Content-Type", "pkglication/json")
		w.WriteHeader(http.StatusBadRequest)
		resBodyBytes := new(bytes.Buffer)
		json.NewEncoder(resBodyBytes).Encode(msg)
		w.Write(resBodyBytes.Bytes())
	}
}
