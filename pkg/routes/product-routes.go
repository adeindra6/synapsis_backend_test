package routes

import (
	"github.com/adeindra6/synapsis_backend_test/pkg/controllers"
	"github.com/gorilla/mux"
)

var RegisterProductRoutes = func(router *mux.Router) {
	router.HandleFunc("/product", controllers.CreateProduct).Methods("POST")
	router.HandleFunc("/product", controllers.GetProducts).Methods("GET")
	router.HandleFunc("/product/{productId}", controllers.GetProductById).Methods("GET")
	router.HandleFunc("/product/{productId}", controllers.UpdateProduct).Methods("PUT")
	router.HandleFunc("/product/{productId}", controllers.DeleteProduct).Methods("DELETE")
}
