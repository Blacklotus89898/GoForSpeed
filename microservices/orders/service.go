package main

import "context"

type service struct {
	store OrdersStore
}



func NewServcie(store OrdersStore) *service {
	return &service{store}
}


func (s *service) CreateOrder(ctx context.Context) error {
	return nil
}

