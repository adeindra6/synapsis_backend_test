package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/adeindra6/synapsis_backend_test/pkg/routes"
	"github.com/gorilla/mux"
	_ "gorm.io/driver/mysql"
)

func main() {
	r := mux.NewRouter()
	routes.RegisterCustomerRoutes(r)
	routes.RegisterProductRoutes(r)
	routes.RegisterCartRoutes(r)
	routes.RegisterPaymentRoutes(r)
	http.Handle("/", r)

	localServer := "http://localhost:8080"
	fmt.Println(fmt.Sprintf("Server running on %s", localServer))
	log.Fatal(http.ListenAndServe("localhost:8080", r))
}
