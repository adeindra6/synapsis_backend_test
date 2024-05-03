package routes

import (
	"github.com/adeindra6/synapsis_backend_test/pkg/controllers"
	"github.com/gorilla/mux"
)

var RegisterPaymentRoutes = func(router *mux.Router) {
	router.HandleFunc("/payment", controllers.CreatePayment).Methods("POST")
	router.HandleFunc("/payment/{paymentId}", controllers.GetPaymentById).Methods("GET")
	router.HandleFunc("/payment/{paymentId}", controllers.CheckoutPayment).Methods("PUT")
	router.HandleFunc("/payment/{paymentId}", controllers.DeletePayment).Methods("DELETE")
}
