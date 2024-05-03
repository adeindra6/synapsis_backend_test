package routes

import (
	"github.com/adeindra6/synapsis_backend_test/pkg/controllers"
	"github.com/gorilla/mux"
)

var RegisterCartRoutes = func(router *mux.Router) {
	router.HandleFunc("/cart", controllers.AddToCart).Methods("POST")
	router.HandleFunc("/cart/{customerId}", controllers.GetUnPaidCartByCustomerId).Methods("GET")
	router.HandleFunc("/cart/paid/{customerId}", controllers.GetPaidCartByCustomerId).Methods("GET")
	router.HandleFunc("/cart/{cartId}", controllers.UpdateItemInCart).Methods("PUT")
	router.HandleFunc("/cart/{cartId}", controllers.DeleteItemsFromCart).Methods("DELETE")
}
