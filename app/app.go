package app

import (
	"fmt"
	"log"
	"net/http"
	"github.com/gorilla/mux"
)

func Start() {
	//mux := http.NewServeMux()
	router := mux.NewRouter()
	
	fmt.Println("Server is listening on port 8082...")

	router.HandleFunc("/greet", greet)
	router .HandleFunc("/customers", getAllCustomers)

	// Start the HTTP server on port 8081
	log.Fatal(http.ListenAndServe("localhost:8082", router))
}
