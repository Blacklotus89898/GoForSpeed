package main

import (
	"log"
	"net/http"
)

const (
	httpAddress = ":8080"

)

func main() {

	mux := http.NewServeMux()
	handler := NewHandler()
	handler.registerRoutes(mux)
	
	log.Printf("Server is running on %s", httpAddress)
	
	if err := http.ListenAndServe(httpAddress, mux); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}


}