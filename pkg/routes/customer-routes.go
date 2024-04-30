package routes

import (
	"github.com/adeindra6/synapsis_backend_test/pkg/controllers"
	"github.com/gorilla/mux"
)

var RegisterCustomerRoutes = func(router *mux.Router) {
	router.HandleFunc("/customer", controllers.CreateCustomer).Methods("POST")
	router.HandleFunc("/customer", controllers.GetCustomers).Methods("GET")
	router.HandleFunc("/customer/{customerId}", controllers.GetCustomerById).Methods("GET")
	router.HandleFunc("/customer/{customerId}", controllers.UpdateCustomer).Methods("PUT")
	router.HandleFunc("/customer/{customerId}", controllers.DeleteCustomer).Methods("DELETE")
	router.HandleFunc("/customer/login", controllers.Login).Methods("POST")
}
