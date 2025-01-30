package main

type service struct {
	store OrderStore
}

func NewOrderService(store OrderStore) *service {
	return &service{store: store}
}

func (s *service) ProcessOrder(orderId string) error {
	return s.store.SaveOrder(orderId)
}
