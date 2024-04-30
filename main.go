package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	_ "gorm.io/driver/mysql"
)

func main() {
	r := mux.NewRouter()
	http.Handle("/", r)

	localServer := "http://localhost:8080"
	fmt.Println(fmt.Sprintf("Server running on %s", localServer))
	log.Fatal(http.ListenAndServe("localhost:8080", r))
}
