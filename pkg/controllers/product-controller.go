package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/adeindra6/synapsis_backend_test/pkg/models"
	"github.com/adeindra6/synapsis_backend_test/pkg/utils"
	"github.com/gorilla/mux"
)

var NewProduct models.Product

func CreateProduct(w http.ResponseWriter, r *http.Request) {
	CreateProduct := &models.Product{}
	utils.ParseBody(r, CreateProduct)
	p := CreateProduct.CreateProduct()
	res, err := json.Marshal(p)
	if err != nil {
		err_msg := ErrMessage{
			Status:  "ERROR",
			Message: "Error while creating new product",
			Code:    http.StatusInternalServerError,
		}
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(err_msg)
	}

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func GetProducts(w http.ResponseWriter, r *http.Request) {
	newProduct := models.GetAllProducts()
	res, err := json.Marshal(newProduct)
	if err != nil {
		fmt.Print("Error when fetching all products")
		err_msg := ErrMessage{
			Status:  "ERROR",
			Message: "Error when fetching all products",
			Code:    http.StatusInternalServerError,
		}
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(err_msg)
	}

	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func GetProductListByCategory(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	newProduct := models.GetProductListByCategory(query["category"][0])

	res, err := json.Marshal(newProduct)
	if err != nil {
		fmt.Println("Error when fetching list products by category")
		err_msg := ErrMessage{
			Status:  "ERROR",
			Message: "Error when fetching list products by category",
			Code:    http.StatusInternalServerError,
		}
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(err_msg)
	}

	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func GetProductById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	productId := vars["productId"]
	id, err := strconv.ParseInt(productId, 0, 0)
	if err != nil {
		fmt.Println("Error while parsing")
	}

	ProductDetails, _ := models.GetProductById(id)
	res, err := json.Marshal(ProductDetails)
	if err != nil {
		fmt.Println("Error when fetching product")
		err_msg := ErrMessage{
			Status:  "ERROR",
			Message: "Error when fetching product",
			Code:    http.StatusInternalServerError,
		}
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(err_msg)
	}

	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func UpdateProduct(w http.ResponseWriter, r *http.Request) {
	var updateProduct = &models.Product{}
	utils.ParseBody(r, updateProduct)
	vars := mux.Vars(r)
	productId := vars["productId"]
	id, err := strconv.ParseInt(productId, 0, 0)

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

	productDetails, db := models.GetProductById(id)
	if updateProduct.ProductName != "" {
		productDetails.ProductName = updateProduct.ProductName
	}
	if updateProduct.Price >= 0 {
		productDetails.Price = updateProduct.Price
	}
	if updateProduct.Stock >= 0 {
		productDetails.Stock = updateProduct.Stock
	}
	if updateProduct.Category != "" {
		productDetails.Category = updateProduct.Category
	}

	db.Save(&productDetails)
	res, err := json.Marshal(productDetails)
	if err != nil {
		fmt.Println("Error while parsing!!!")
	}

	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func DeleteProduct(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	productId := vars["productId"]
	id, err := strconv.ParseInt(productId, 0, 0)

	if err != nil {
		fmt.Println("Error when deleting product")
		err_msg := ErrMessage{
			Status:  "ERROR",
			Message: "Error when deleting product",
			Code:    http.StatusInternalServerError,
		}
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(err_msg)
	}

	product := models.DeleteProductById(id)
	res, err := json.Marshal(product)
	if err != nil {
		fmt.Println("Error while parsing!!!")
	}

	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
