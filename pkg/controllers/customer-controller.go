package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/adeindra6/synapsis_backend_test/pkg/models"
	"github.com/adeindra6/synapsis_backend_test/pkg/utils"
	"github.com/gorilla/mux"
	"golang.org/x/crypto/bcrypt"
)

var NewCustomer models.Customer

type ErrMessage struct {
	Status  string `json:"status"`
	Message string `json:"message"`
	Code    int64  `json:"code"`
}

func CreateCustomer(w http.ResponseWriter, r *http.Request) {
	CreateCustomer := &models.Customer{}
	utils.ParseBody(r, CreateCustomer)
	c := CreateCustomer.CreateCustomer()
	res, err := json.Marshal(c)
	if err != nil {
		err_msg := ErrMessage{
			Status:  "ERROR",
			Message: "Error while creating new customer",
			Code:    http.StatusInternalServerError,
		}
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(err_msg)
	}

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func GetCustomers(w http.ResponseWriter, r *http.Request) {
	newCustomer := models.GetAllCustomers()
	res, err := json.Marshal(newCustomer)
	if err != nil {
		fmt.Print("Error when fetching all customers")
		err_msg := ErrMessage{
			Status:  "ERROR",
			Message: "Error when fetching all customers",
			Code:    http.StatusInternalServerError,
		}
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(err_msg)
	}

	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func GetCustomerById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	customerId := vars["customerId"]
	id, err := strconv.ParseInt(customerId, 0, 0)
	if err != nil {
		fmt.Println("Error while parsing")
	}

	CustomerDetails, _ := models.GetCustomerById(id)
	res, err := json.Marshal(CustomerDetails)
	if err != nil {
		fmt.Println("Error when fetching customer")
		err_msg := ErrMessage{
			Status:  "ERROR",
			Message: "Error when fetching customer",
			Code:    http.StatusInternalServerError,
		}
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(err_msg)
	}

	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func UpdateCustomer(w http.ResponseWriter, r *http.Request) {
	var updateCustomer = &models.Customer{}
	utils.ParseBody(r, updateCustomer)
	vars := mux.Vars(r)
	customerId := vars["customerId"]
	id, err := strconv.ParseInt(customerId, 0, 0)

	if err != nil {
		fmt.Println("Error when updating customer")
		err_msg := ErrMessage{
			Status:  "ERROR",
			Message: "Error when updating customer",
			Code:    http.StatusInternalServerError,
		}
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(err_msg)
	}

	customerDetails, db := models.GetCustomerById(id)
	if updateCustomer.Fullname != "" {
		customerDetails.Fullname = updateCustomer.Fullname
	}
	if updateCustomer.Email != "" {
		customerDetails.Email = updateCustomer.Email
	}
	if updateCustomer.Address != "" {
		customerDetails.Address = updateCustomer.Address
	}
	if updateCustomer.Phone != "" {
		customerDetails.Phone = updateCustomer.Phone
	}
	if updateCustomer.Password != "" {
		newPassword, err := bcrypt.GenerateFromPassword([]byte(updateCustomer.Password), bcrypt.DefaultCost)
		if err != nil {
			panic(err)
		}

		customerDetails.Password = string(newPassword)
	}

	db.Save(&customerDetails)
	res, err := json.Marshal(customerDetails)
	if err != nil {
		fmt.Println("Error while parsing!!!")
	}

	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func DeleteCustomer(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	customerId := vars["customerId"]
	id, err := strconv.ParseInt(customerId, 0, 0)

	if err != nil {
		fmt.Println("Error when deleting customer")
		err_msg := ErrMessage{
			Status:  "ERROR",
			Message: "Error when deleting customer",
			Code:    http.StatusInternalServerError,
		}
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(err_msg)
	}

	customer := models.DeleteCustomerById(id)
	res, err := json.Marshal(customer)
	if err != nil {
		fmt.Println("Error while parsing!!!")
	}

	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func Login(w http.ResponseWriter, r *http.Request) {
	Login := &models.Customer{}
	utils.ParseBody(r, Login)
	c := Login.Login()
	res, err := json.Marshal(c)
	if err != nil {
		err_msg := ErrMessage{
			Status:  "ERROR",
			Message: "Wrong Email or Password!",
			Code:    http.StatusUnauthorized,
		}
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(err_msg)
	}

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
