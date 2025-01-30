package main

type store struct {
	// add mongo
}

func NewStore() *store {
	return &store{}
}

func (s *store) SaveOrder(orderId string) error {
	return nil
}
