package main

import "net/http"

type handler struct {
	// gateway

}

func NewHandler() *handler {
	return &handler{}
}

func (h *handler) registerRoutes(mux *http.ServeMux) {
	mux.HandleFunc("POST /api/customers/{customerID}/orders", h.HandleCreateOrder)
	
}

func (h *handler) HandleCreateOrder(w http.ResponseWriter, r *http.Request) {
	// 1. Parse the request body to get the order details
	// 2. Validate the order details
	// 3. Call the orders service to create the order
	// 4. Return the response to the client

}