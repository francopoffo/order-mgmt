package main

type OrderService interface {
	ProcessOrder(orderId string) error
}

type OrderStore interface {
	SaveOrder(orderId string) error
}
