package model

type OrderItem interface {
	isOrderItem()
}

// TODO: create unmarshaller
type OrderItems []OrderItem
