package main

import "context"

func main() {

	store := NewStore()
	service := NewServcie(store)


	service.CreateOrder(context.Background())

}