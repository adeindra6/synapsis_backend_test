package controllers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/adeindra6/synapsis_backend_test/pkg/models"
	"github.com/adeindra6/synapsis_backend_test/pkg/utils"
	"github.com/dgrijalva/jwt-go"
	"github.com/gorilla/mux"
)

var NewCart models.Cart
var jwtKey = "Synapsis_Test"

type NonAuthorizedMsg struct {
	Status  string `json:"status"`
	Message string `json:"message"`
	Code    int64  `json:"code"`
}

type CustomerCart struct {
	Customer models.Customer  `json:"customer"`
	Product  []models.Product `json:"product"`
	Amount   uint64           `json:"amount"`
	Total    uint64           `json:"total"`
}

func AddToCart(w http.ResponseWriter, r *http.Request) {
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
		AddToCart := &models.Cart{}
		utils.ParseBody(r, AddToCart)
		AddToCart.CustomerId = customer_id
		AddToCart.IsPaid = false

		product, _ := models.GetProductById(AddToCart.ProductId)
		if product.ID > 0 {
			amount := product.Price * AddToCart.Quantity
			AddToCart.Amount = amount

			c := AddToCart.AddToCart()
			res, err := json.Marshal(c)
			if err != nil {
				fmt.Println("Error while adding to cart")
				err_msg := ErrMessage{
					Status:  "ERROR",
					Message: "Error while adding to cart",
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
			msg.Message = "Product Not Found"
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

func GetUnPaidCartByCustomerId(w http.ResponseWriter, r *http.Request) {
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
		customerId := vars["customerId"]
		id, err := strconv.ParseInt(customerId, 0, 0)
		if err != nil {
			fmt.Println("Error while parsing")
		}

		var customerCart CustomerCart
		customer, _ := models.GetCustomerById(id)
		customerCart.Customer.ID = customer.ID
		customerCart.Customer.Fullname = customer.Fullname
		customerCart.Customer.Email = customer.Email
		customerCart.Customer.Address = customer.Address
		customerCart.Customer.Phone = customer.Phone
		customerCart.Customer.CreatedAt = customer.CreatedAt
		customerCart.Customer.UpdatedAt = customer.UpdatedAt
		customerCart.Customer.DeletedAt = customer.DeletedAt

		carts := models.GetCartByCustomerId(id, false)
		customerCart.Total = 0
		customerCart.Product = make([]models.Product, 0)
		for i, v := range carts {
			product, _ := models.GetProductById(v.ProductId)
			cp := models.Product{
				ProductName: product.ProductName,
				Price:       product.Price,
				Category:    product.Category,
			}
			customerCart.Product = append(customerCart.Product, cp)
			customerCart.Product[i].ID = v.ID
			customerCart.Product[i].CreatedAt = v.CreatedAt
			customerCart.Product[i].UpdatedAt = v.UpdatedAt
			customerCart.Product[i].DeletedAt = v.DeletedAt

			customerCart.Amount = v.Amount
			customerCart.Total = customerCart.Total + v.Amount
		}

		res, err := json.Marshal(customerCart)
		if err != nil {
			fmt.Println("Error when fetching customer's cart")
			err_msg := ErrMessage{
				Status:  "ERROR",
				Message: "Error when fetching customer's cart",
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
		w.WriteHeader(http.StatusUnauthorized)
		resBodyBytes := new(bytes.Buffer)
		json.NewEncoder(resBodyBytes).Encode(msg)
		w.Write(resBodyBytes.Bytes())
	}
}

func GetPaidCartByCustomerId(w http.ResponseWriter, r *http.Request) {
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
		customerId := vars["customerId"]
		id, err := strconv.ParseInt(customerId, 0, 0)
		if err != nil {
			fmt.Println("Error while parsing")
		}

		var customerCart CustomerCart
		customer, _ := models.GetCustomerById(id)
		customerCart.Customer.ID = customer.ID
		customerCart.Customer.Fullname = customer.Fullname
		customerCart.Customer.Email = customer.Email
		customerCart.Customer.Address = customer.Address
		customerCart.Customer.Phone = customer.Phone
		customerCart.Customer.CreatedAt = customer.CreatedAt
		customerCart.Customer.UpdatedAt = customer.UpdatedAt
		customerCart.Customer.DeletedAt = customer.DeletedAt

		carts := models.GetCartByCustomerId(id, true)
		customerCart.Total = 0
		customerCart.Product = make([]models.Product, 0)
		for i, v := range carts {
			product, _ := models.GetProductById(v.ProductId)
			cp := models.Product{
				ProductName: product.ProductName,
				Price:       product.Price,
				Category:    product.Category,
			}
			customerCart.Product = append(customerCart.Product, cp)
			customerCart.Product[i].ID = v.ID
			customerCart.Product[i].CreatedAt = v.CreatedAt
			customerCart.Product[i].UpdatedAt = v.UpdatedAt
			customerCart.Product[i].DeletedAt = v.DeletedAt

			customerCart.Amount = v.Amount
			customerCart.Total = customerCart.Total + v.Amount
		}

		res, err := json.Marshal(customerCart)
		if err != nil {
			fmt.Println("Error when fetching customer's cart")
			err_msg := ErrMessage{
				Status:  "ERROR",
				Message: "Error when fetching customer's cart",
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
		w.WriteHeader(http.StatusUnauthorized)
		resBodyBytes := new(bytes.Buffer)
		json.NewEncoder(resBodyBytes).Encode(msg)
		w.Write(resBodyBytes.Bytes())
	}
}

func UpdateItemInCart(w http.ResponseWriter, r *http.Request) {
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
		var updateCart = &models.Cart{}
		utils.ParseBody(r, updateCart)
		vars := mux.Vars(r)
		cartId := vars["cartId"]

		id, err := strconv.ParseInt(cartId, 0, 0)
		if err != nil {
			fmt.Println("Error when updating cart")
			err_msg := ErrMessage{
				Status:  "ERROR",
				Message: "Error when updating cart",
				Code:    http.StatusInternalServerError,
			}
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(err_msg)
		}

		cartDetails, db := models.GetCartById(id)
		product, _ := models.GetProductById(cartDetails.ProductId)
		if updateCart.Quantity > 0 {
			cartDetails.Quantity = updateCart.Quantity
			cartDetails.Amount = product.Price * updateCart.Quantity
		}

		db.Save(&cartDetails)

		var resMsg NonAuthorizedMsg
		resMsg.Code = http.StatusOK
		resMsg.Message = "Success Updating shopping cart"
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
		w.WriteHeader(http.StatusUnauthorized)
		resBodyBytes := new(bytes.Buffer)
		json.NewEncoder(resBodyBytes).Encode(msg)
		w.Write(resBodyBytes.Bytes())
	}
}

func DeleteItemsFromCart(w http.ResponseWriter, r *http.Request) {
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
		cartId := vars["cartId"]
		id, err := strconv.ParseInt(cartId, 0, 0)
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

		_ = models.DeleteCartById(id)
		var resMsg NonAuthorizedMsg
		resMsg.Code = http.StatusOK
		resMsg.Message = "Success Delete Item from shopping cart"
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
		w.WriteHeader(http.StatusUnauthorized)
		resBodyBytes := new(bytes.Buffer)
		json.NewEncoder(resBodyBytes).Encode(msg)
		w.Write(resBodyBytes.Bytes())
	}
}
