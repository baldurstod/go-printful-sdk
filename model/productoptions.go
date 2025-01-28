package model

type ProductOption interface {
	IsProductOption()
}

// TODO: create unmarshaller
type ProductOptions []ProductOption
